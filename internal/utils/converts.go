package utils

import (
	"fmt"
	"time"
)

func BytesToGB(value uint64 ) float64 {
	return float64(value) / 1024 / 1024 / 1024
}

func FormatUptime(seconds uint64) string {
	d := time.Duration(seconds) * time.Second

	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60

	return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
}

func FormatBootTime(unix uint64) string {
	return time.Unix(int64(unix), 0).
		Local().
		Format("2006-01-02 15:04:05")
}