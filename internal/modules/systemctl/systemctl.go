package systemctl

import (
	"strings"
	"nandcli/internal/shell"
)

func GetService(name string) (*Service, error) {

	output, err := shell.Run("systemctl", "show", name)

	if err != nil {
		return nil, err
	}

	service := &Service{}

	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		switch key {
		case "Id":
			service.ID = value

		case "ActiveState":
			service.ActiveState = value

		case "SubState":
			service.SubState = value

		case "UnitFileState":
			service.UnitFileState = value

		case "Description":
			service.Description = value
		}
	}

		return service, nil
}
