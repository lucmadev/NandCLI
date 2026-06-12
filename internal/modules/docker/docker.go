package docker

import "nandcli/internal/shell"

func ListContainers() (Container, error) {
	output, err := shell.Run("docker", "ps", "--format json")
	return ConvertJsonToContainer(output), err
}

func ListAllContainers() ([]byte, error) {
	return shell.Run("docker", "ps", "-a", "--format json")
}

func StartContainer(container string) ([]byte, error) {
	return shell.Run("docker", "start", container)
}

func RestartContainer(container string) ([]byte, error) {
	return shell.Run("docker", "restart", container)
}

func StopContainer(container string) ([]byte, error) {
	return shell.Run("docker", "stop", container)
}

func ObtainLogsContainer(container string) ([]byte, error) {
	return shell.Run("docker", "logs", container)
}

func RemoveContainer(container string) ([]byte, error) {
	return shell.Run("docker", "rm", container)
}

func InspectContainer(container string) (Container, error) {
	output, err := shell.Run("docker", "inspect", container)
	return ConvertJsonToContainer(output), err
}

func ListImages() ([]byte, error) {
	return shell.Run("docker", "images")
}
