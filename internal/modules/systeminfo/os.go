package systeminfo

import (
	"github.com/shirou/gopsutil/v4/host"
	"nandcli/internal/nand/version"
)

type SystemInfo struct {
    OSName        string
    OSFamily      string
    OSVersion     string
    KernelVersion string
    NandVersion   string
}

func GetOSInfo() (*SystemInfo, error) {
	h, err := host.Info()
	if err != nil {
		return nil, err
	}

	return &SystemInfo{
		OSName:        h.Platform,
		OSFamily:      h.PlatformFamily,
		OSVersion:     h.PlatformVersion,
		KernelVersion: h.KernelVersion,
		NandVersion:   version.NandVersion,
	}, nil
}