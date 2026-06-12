package apt

import "nandcli/internal/shell"

func InstallPackage(packageName string) ([]byte, error) {
	return shell.Run("sudo", "apt", "install", packageName)
}

func SearchPackages(search string) ([]byte, error) {
	return shell.Run("apt", "search", search)
}

func ViewPackage(packageName string) ([]byte, error) {
	return shell.Run("apt", "show", packageName)
}

func UpdatePackages() ([]byte, error) {
	return shell.Run("sudo", "apt-get", "update", "-y")
}

func UpgradePackages() ([]byte, error) {
	return shell.Run("sudo", "apt-get", "upgrade", "-y")
}



