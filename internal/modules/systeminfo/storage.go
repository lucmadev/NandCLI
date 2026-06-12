package systeminfo

import "github.com/shirou/gopsutil/v4/disk"

type DiskInfo struct {
	Device      string
	MountPoint  string
	FileSystem  string
	Total       uint64
	Used        uint64
	Free        uint64
	UsedPercent float64
}

func GetDisks() ([]DiskInfo, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	var disks []DiskInfo

	for _, p := range partitions {
		u, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}

		disks = append(disks, DiskInfo{
			Device:      p.Device,
			MountPoint:  p.Mountpoint,
			FileSystem:  p.Fstype,
			Total:       u.Total,
			Used:        u.Used,
			Free:        u.Free,
			UsedPercent: u.UsedPercent,
		})
	}

	return disks, nil
}