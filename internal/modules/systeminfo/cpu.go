package systeminfo

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUInfo struct {
	ModelName string
	Cores int32
	Threads int
	PercentUsage float64
}

func GetCPUInfo() (*CPUInfo, error) {
	info, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	usage, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	logical, _ := cpu.Counts(true)

	if len(info) == 0 {
		return nil, fmt.Errorf("no cpu information found")
	}

	return &CPUInfo{
		ModelName: info[0].ModelName,
		Cores:     info[0].Cores,
		Threads: logical,
		PercentUsage: usage[0],
	}, nil
}