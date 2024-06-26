package linux

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jochenvg/go-udev"
	"github.com/neuroplastio/neuroplastio/internal/configsvc"
	"github.com/neuroplastio/neuroplastio/internal/hidsvc"
	"github.com/psanford/uhid"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/sstallion/go-hid"
	"go.uber.org/zap"
)

var defaultBackendOptions = backendOptions{
	pollInterval: 1 * time.Second,
}

type backendOptions struct {
	pollInterval time.Duration
}

func WithPollInterval(d time.Duration) Option {
	return func(o *backendOptions) {
		o.pollInterval = d
	}
}

type Option func(*backendOptions)

// Backend implements the hidsvc.Backend interface for Linux Kernel.
// It uses hidapi, udev and uhid kernel modules to communicate with HID devices.
type Backend struct {
	log     *zap.Logger
	options backendOptions
	udev    udev.Udev

	config   *configsvc.Service
	uhidPath string

	hidDevices  *xsync.MapOf[HidAddress, hid.DeviceInfo]
	uhidDevices *xsync.MapOf[string, UhidDeviceConfig]

	ready chan struct{}

	publisher hidsvc.BackendPublisher
}

type HidAddress struct {
	VendorID  uint16
	ProductID uint16
	Interface int
}

func (a HidAddress) String() string {
	return fmt.Sprintf("%04x:%04x:%d", a.VendorID, a.ProductID, a.Interface)
}

func ParseHidAddress(s string) (HidAddress, error) {
	var addr HidAddress
	_, err := fmt.Sscanf(s, "%04x:%04x:%d", &addr.VendorID, &addr.ProductID, &addr.Interface)
	if err != nil {
		return HidAddress{}, err
	}
	return addr, nil
}

func NewBackend(log *zap.Logger, configSvc *configsvc.Service, uhidPath string, opts ...Option) *Backend {
	options := defaultBackendOptions
	for _, opt := range opts {
		opt(&options)
	}

	return &Backend{
		options:     options,
		log:         log,
		config:      configSvc,
		uhidPath:    uhidPath,
		udev:        udev.Udev{},
		ready:       make(chan struct{}),
		hidDevices:  xsync.NewMapOf[HidAddress, hid.DeviceInfo](),
		uhidDevices: xsync.NewMapOf[string, UhidDeviceConfig](),
	}
}

type UhidConfig struct {
	Uhid []UhidDeviceConfig `json:"uhid"`
}

type UhidDeviceConfig struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	VendorID  uint32 `json:"vendorId"`
	ProductID uint32 `json:"productId"`
}

func (b *Backend) Ready() <-chan struct{} {
	return b.ready
}

func (b *Backend) Start(ctx context.Context, publisher hidsvc.BackendPublisher) error {
	hid.Init()
	defer hid.Exit()

	b.publisher = publisher

	b.log.Info("Starting Linux HID backend")
	select {
	case <-ctx.Done():
		return nil
	case <-b.config.Ready():
	}

	uhidConfig, err := configsvc.Register(b.config, b.uhidPath, UhidConfig{}, func(cfg UhidConfig, err error) {
		b.onUhidConfigChange(ctx, cfg, err)
	})
	if err != nil {
		return fmt.Errorf("failed to register UHID config: %w", err)
	}

	err = b.refreshUhidDevices(ctx, uhidConfig)
	if err != nil {
		return fmt.Errorf("failed to refresh UHID devices: %w", err)
	}

	err = b.refreshHidDevices(ctx)
	if err != nil {
		return fmt.Errorf("failed to refresh HID devices: %w", err)
	}

	close(b.ready)
	b.log.Info("Linux HID backend started")

	pollTicker := time.NewTicker(b.options.pollInterval)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-pollTicker.C:
			err := b.refreshHidDevices(ctx)
			if err != nil {
				b.log.Error("failed to refresh HID devices", zap.Error(err))
				continue
			}
		}
	}
}

func (b *Backend) onUhidConfigChange(ctx context.Context, cfg UhidConfig, err error) {
	if err != nil {
		b.log.Error("failed to parse UHID config", zap.Error(err))
		return
	}
	b.refreshUhidDevices(ctx, cfg)
}

func (b *Backend) refreshUhidDevices(ctx context.Context, cfg UhidConfig) error {
	newDevices := make(map[string]UhidDeviceConfig)
	for _, dev := range cfg.Uhid {
		newDevices[dev.ID] = dev
	}
	var disconnected []string
	var connected []hidsvc.BackendOutputDevice
	b.uhidDevices.Range(func(id string, dev UhidDeviceConfig) bool {
		if _, ok := newDevices[id]; !ok {
			disconnected = append(disconnected, fmt.Sprintf("uhid:%s", id))
			b.uhidDevices.Delete(id)
			return true
		}
		delete(newDevices, id)
		return true
	})
	for id, dev := range newDevices {
		b.uhidDevices.Store(id, dev)
		connected = append(connected, hidsvc.BackendOutputDevice{
			Address: fmt.Sprintf("uhid:%s", id),
			Name:    dev.Name,
		})
	}
	if len(connected) > 0 || len(disconnected) > 0 {
		b.publisher(ctx, hidsvc.BackendEvent{
			OutputsChanged: &hidsvc.BackendEventOutputsChanged{
				Connected:    connected,
				Disconnected: disconnected,
			},
		})
	}
	return nil
}

func (b *Backend) refreshHidDevices(ctx context.Context) error {
	newDevices, err := b.enumerateHidDevices()
	// TODO: exclude known uhid output devices
	if err != nil {
		return err
	}
	var disconnected []string
	var connected []hidsvc.BackendInputDevice
	b.hidDevices.Range(func(addr HidAddress, dev hid.DeviceInfo) bool {
		if _, ok := newDevices[addr]; !ok {
			disconnected = append(disconnected, addr.String())
			b.hidDevices.Delete(addr)
			return true
		}
		delete(newDevices, addr)
		return true
	})

	for addr, device := range newDevices {
		desc, err := b.loadDescriptor(device)
		if err != nil {
			b.log.Error("failed to load HID descriptor", zap.Error(err), zap.String("device", device.Path))
			continue
		}
		b.hidDevices.Store(addr, device)
		connected = append(connected, hidsvc.BackendInputDevice{
			Address:          addr.String(),
			Name:             generateName(device),
			ReportDescriptor: desc,
		})
	}

	if len(connected) > 0 || len(disconnected) > 0 {
		b.publisher(ctx, hidsvc.BackendEvent{
			InputsChanged: &hidsvc.BackendEventInputsChanged{
				Connected:    connected,
				Disconnected: disconnected,
			},
		})
	}

	return nil
}

func (b *Backend) loadDescriptor(device hid.DeviceInfo) ([]byte, error) {
	dev, err := hid.OpenPath(device.Path)
	if err != nil {
		return nil, err
	}
	defer dev.Close()

	reportDesc := make([]byte, 4096)
	count, err := dev.GetReportDescriptor(reportDesc)
	if err != nil {
		return nil, err
	}
	return reportDesc[:count], nil
}

func generateName(device hid.DeviceInfo) string {
	var parts []string
	if device.MfrStr != "" {
		parts = append(parts, device.MfrStr)
	}
	if device.ProductStr != "" {
		parts = append(parts, device.ProductStr)
	}
	if len(parts) == 0 {
		return fmt.Sprintf("%04x:%04x", device.VendorID, device.ProductID)
	}
	return strings.Join(parts, " ")
}

func (b *Backend) enumerateHidDevices() (map[HidAddress]hid.DeviceInfo, error) {
	devices := make(map[HidAddress]hid.DeviceInfo)
	err := hid.Enumerate(hid.VendorIDAny, hid.ProductIDAny, func(device *hid.DeviceInfo) error {
		addr := HidAddress{
			VendorID:  device.VendorID,
			ProductID: device.ProductID,
			Interface: device.InterfaceNbr,
		}
		devices[addr] = *device
		return nil
	})
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (b *Backend) OpenInput(id string) (hidsvc.BackendInputDeviceHandle, error) {
	addr, err := ParseHidAddress(id)
	if err != nil {
		return nil, err
	}
	dev, ok := b.hidDevices.Load(addr)
	if !ok {
		return nil, fmt.Errorf("device not found: %s", id)
	}
	hidDev, err := hid.OpenPath(dev.Path)
	if err != nil {
		return nil, err
	}
	ud := udev.Udev{}
	enumerate := ud.NewEnumerate()
	enumerate.AddMatchSubsystem("input")
	enumerate.AddMatchProperty("ID_USB_VENDOR_ID", fmt.Sprintf("%04x", dev.VendorID))
	enumerate.AddMatchProperty("ID_USB_MODEL_ID", fmt.Sprintf("%04x", dev.ProductID))
	enumerate.AddMatchProperty("ID_USB_INTERFACE_NUM", fmt.Sprintf("%02d", dev.InterfaceNbr))
	devices, err := enumerate.Devices()
	if err != nil {
		return nil, fmt.Errorf("failed to enumerate devices: %w", err)
	}
	var udevDev *udev.Device
	for _, device := range devices {
		if !strings.HasPrefix(device.Sysname(), "event") {
			continue
		}
		udevDev = device
		break
	}
	if udevDev == nil {
		return nil, fmt.Errorf("device not found in udev")
	}
	err = os.WriteFile(udevDev.Syspath()+"/uevent", []byte("remove"), 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to remove device: %w", err)
	}

	return &hidDeviceHandle{
		log:  b.log,
		hid:  hidDev,
		udev: udevDev,
	}, nil
}

type hidDeviceHandle struct {
	log  *zap.Logger
	hid  *hid.Device
	udev *udev.Device
}

func (h *hidDeviceHandle) Read(buf []byte) (int, error) {
	n, err := h.hid.Read(buf)
	return n, err
}

func (h *hidDeviceHandle) GetInputReport() ([]byte, error) {
	buf := make([]byte, 4096)
	n, err := h.hid.GetInputReport(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func (h *hidDeviceHandle) Close() error {
	h.log.Debug("closing device", zap.String("path", h.udev.Syspath()))
	err := os.WriteFile(h.udev.Syspath()+"/uevent", []byte("add"), 0644)
	if err != nil {
		h.log.Error("failed to re-add device", zap.Error(err))
	}
	return h.hid.Close()
}

func (h *hidDeviceHandle) Write(buf []byte) (int, error) {
	return h.hid.Write(buf)
}

func (b *Backend) OpenOutput(id string, desc []byte) (hidsvc.BackendOutputDeviceHandle, error) {
	if !strings.HasPrefix(id, "uhid:") {
		return nil, fmt.Errorf("invalid output device address: %s", id)
	}
	id = strings.TrimPrefix(id, "uhid:")
	dev, ok := b.uhidDevices.Load(id)
	if !ok {
		return nil, fmt.Errorf("device not found: %s", id)
	}
	uhidDev, err := uhid.NewDevice(id, desc)
	if err != nil {
		return nil, fmt.Errorf("failed to create uhid device: %w", err)
	}

	uhidDev.Data.Bus = 0x03
	uhidDev.Data.VendorID = dev.VendorID
	uhidDev.Data.ProductID = dev.ProductID

	ctx, cancel := context.WithCancel(context.Background())
	events, err := uhidDev.Open(ctx)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to open uhid device: %w", err)
	}

	return &uhidDeviceHandle{
		log:    b.log,
		uhid:   uhidDev,
		events: events,
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

type uhidDeviceHandle struct {
	log    *zap.Logger
	uhid   *uhid.Device
	events chan uhid.Event
	ctx    context.Context
	cancel context.CancelFunc
}

func (h *uhidDeviceHandle) Close() error {
	h.cancel()
	return h.uhid.Close()
}

func (h *uhidDeviceHandle) Write(buf []byte) (int, error) {
	err := h.uhid.InjectEvent(buf)
	if err != nil {
		return 0, err
	}
	return len(buf), nil
}

func (h *uhidDeviceHandle) Read(buf []byte) (int, error) {
	for {
		select {
		case <-h.ctx.Done():
			return 0, h.ctx.Err()
		case event := <-h.events:
			if event.Type != uhid.Output {
				continue
			}
			n := copy(buf, event.Data)
			return n, nil
		}
	}
}
