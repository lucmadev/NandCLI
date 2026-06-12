package systeminfo

import "github.com/shirou/gopsutil/v4/net"

type NetworkInterface struct {
	Name string
	IPs  []string
}

func GetNetworkInterfaces() ([]NetworkInterface, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var result []NetworkInterface

	for _, iface := range ifaces {
		var ips []string

		for _, addr := range iface.Addrs {
			ips = append(ips, addr.Addr)
		}

		result = append(result, NetworkInterface{
			Name: iface.Name,
			IPs:  ips,
		})
	}

	return result, nil
}