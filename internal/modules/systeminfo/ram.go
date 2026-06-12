package systeminfo

import (
	"github.com/shirou/gopsutil/v4/mem"
)

type RAMStatus struct {
	Total     uint64
	Used      uint64
	Available uint64
}

func GetRAMInfo() (*RAMStatus, error) {

	vm, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	return &RAMStatus{
		Total:     vm.Total,
		Used:      vm.Used,
		Available: vm.Available,
	}, nil

}