package systeminfo

import (
	"github.com/shirou/gopsutil/v4/host"
)

type HostInfo struct {
	Hostname      string
	HostID        string
	Processes     uint64
	Uptime        uint64
	BootTime      uint64
	KernelArch    string
	KernelVersion string
	OS            string
	Family        string
	Version       string
}

func GetHostInfo() (*HostInfo, error) {
	info, err := host.Info()
	if err != nil {
		return nil, err
	}

	return &HostInfo{
		Hostname:      info.Hostname,
		HostID:        info.HostID,
		Processes:     info.Procs,
		Uptime:        info.Uptime,
		BootTime:      info.BootTime,
		KernelArch:    info.KernelArch,
		KernelVersion: info.KernelVersion,
		OS:            info.Platform,
		Family:        info.PlatformFamily,
		Version:       info.PlatformVersion,
	}, nil
}