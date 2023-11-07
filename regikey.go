package html

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows/registry"
)

const registryKey = `Software\Microsoft\Windows\CurrentVersion\Run`

func REGISTRY() {
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}

	err = addStartupEntry("MyGolangApp", executablePath)
	if err != nil {
		fmt.Println("Error adding startup entry:", err)
		return
	}

	fmt.Println("Your Golang application has been added to startup.")
}

func addStartupEntry(entryName, executablePath string) error {
	key, err := registry.OpenKey(registry.CURRENT_USER, registryKey, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	err = key.SetStringValue(entryName, executablePath)
	if err != nil {
		return err
	}

	return nil
}
