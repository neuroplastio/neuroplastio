package hidapi

import (
	"github.com/neuroplastio/neio-agent/hidapi/hiddesc"
	"github.com/neuroplastio/neio-agent/pkg/bits"
)

type UsageValues interface {
	Contains(usage Usage) bool
	Usages() []Usage
	LogicalMinimum() int32
	LogicalMaximum() int32
	GetValue(bits bits.Bits, usage Usage) int32
	SetValue(bits bits.Bits, usage Usage, value int32)
}

type usageValues struct {
	usages     []Usage
	usageIndex map[Usage]int
	size       uint32
	minimum    int32
	maximum    int32
}

func NewUsageValues(usages []Usage, size uint32, logicalMinimum, logicalMaximum int32) UsageValues {
	usageIndex := make(map[Usage]int, len(usages))
	for i, usage := range usages {
		usageIndex[usage] = i
	}
	return &usageValues{
		usages:     usages,
		usageIndex: usageIndex,
		size:       size,
		minimum:    logicalMinimum,
		maximum:    logicalMaximum,
	}
}

func (u usageValues) Contains(usage Usage) bool {
	_, ok := u.usageIndex[usage]
	return ok
}

func (u usageValues) Usages() []Usage {
	return u.usages
}

func (u usageValues) LogicalMinimum() int32 {
	return u.minimum
}

func (u usageValues) LogicalMaximum() int32 {
	return u.maximum
}

func (u usageValues) GetValue(bits bits.Bits, usage Usage) int32 {
	index, ok := u.usageIndex[usage]
	if !ok {
		return 0
	}

	switch {
	case u.size == 7:
		return int32(int8(bits.Uint7(index)))
	case u.size == 8:
		return int32(int8(bits.Uint8(index)))
	case u.size == 16:
		return int32(int16(bits.Uint16(index)))
	case u.size == 24:
		return int32(int16(bits.Uint24(index)))
	case u.size == 32:
		return int32(bits.Uint32(index))
	default:
		return 0
	}
}

func (u usageValues) SetValue(bits bits.Bits, usage Usage, value int32) {
	index, ok := u.usageIndex[usage]
	if !ok {
		return
	}

	switch {
	case u.size == 7:
		bits.SetUint7(index, uint8(int8(value)))
	case u.size == 8:
		bits.SetUint8(index, uint8(int8(value)))
	case u.size == 16:
		bits.SetUint16(index, uint16(int16(value)))
	case u.size == 24:
		bits.SetUint24(index, uint16(int16(value)))
	case u.size == 32:
		bits.SetUint32(index, uint32(value))
	}
}

func NewUsageValuesItems(dataItems []hiddesc.DataItem) map[int]UsageValues {
	values := make(map[int]UsageValues)
	for i, item := range dataItems {
		if len(item.UsageIDs) == 0 || !item.Flags.IsVariable() {
			// not a usage-value data item
			continue
		}
		// TODO: dynamic size
		switch item.ReportSize {
		case 7, 8, 16, 24, 32:
		default:
			// not a usage-value data item
			continue
		}
		usages := make([]Usage, len(item.UsageIDs))
		for j, id := range item.UsageIDs {
			usages[j] = NewUsage(item.UsagePage, id)
		}
		values[i] = NewUsageValues(usages, item.ReportSize, item.LogicalMinimum, item.LogicalMaximum)
	}
	return values
}
