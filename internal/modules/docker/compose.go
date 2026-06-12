package docker

import "nandcli/internal/shell"

func ComposeUpDaemon() ([]byte, error) {
	return shell.Run("docker", "compose", "up", "-d")
}

func ComposeDown() ([]byte, error) {
	return shell.Run("docker", "compose", "down")
}

func ComposeRestart() ([]byte, error) {
	return shell.Run("docker", "compose", "restart")
}

func ComposeLogs() ([]byte, error) {
	return shell.Run("docker", "compose", "logs")
}

func ComposeList(container string) ([]byte, error) {
	return shell.Run("docker", "compose", "ps")
}