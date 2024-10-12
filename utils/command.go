package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
)

var Verbose bool

// Function to run shell commands for dependency installation
func RunCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
	}
}

// Function to run shell commands with error handling, optional verbosity, and message filtering
func RunCommandWithProgress(command string, args ...string) error {
	bar := progressbar.Default(100)

	// Simulate a delay for the progress bar (for demo purposes)
	go func() {
		for i := 0; i < 100; i++ {
			bar.Add(1)
			time.Sleep(25 * time.Millisecond)
		}
	}()

	cmd := exec.Command(command, args...)

	// Capture output in case we're not in verbose mode
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	err := cmd.Run()
	bar.Finish()

	// If the command fails, return the error
	if err != nil {
		if !Verbose {
			fmt.Println("Error:", errBuffer.String()) // Show captured error if not in verbose mode
		}
		return fmt.Errorf("error running command %s: %w", command, err)
	}

	// Filter the "requirement already satisfied" message for pip
	if !Verbose {
		output := outBuffer.String()
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			if !strings.Contains(line, "Requirement already satisfied") {
				fmt.Println(line)
			}
		}
	} else {
		// In verbose mode, just print everything
		fmt.Println(outBuffer.String())
	}

	return nil
}
