package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pod32g/servosetup/utils"
)

func main() {
	// Parse the verbose flag
	verboseFlag := flag.Bool("verbose", false, "Enable verbose output")
	flag.Parse()

	// Set the verbosity in the utils package
	utils.Verbose = *verboseFlag

	reader := bufio.NewReader(os.Stdin)

	// Select the language
	fmt.Println("Select the language:")
	fmt.Println("1) Go")
	fmt.Println("2) Python")
	fmt.Println("3) Node.js")

	langChoice, _ := reader.ReadString('\n')
	langChoice = strings.TrimSpace(langChoice)

	// Select the database
	fmt.Println("Select the database:")
	fmt.Println("1) PostgreSQL")
	fmt.Println("2) MongoDB")
	fmt.Println("3) SQLite")

	dbChoice, _ := reader.ReadString('\n')
	dbChoice = strings.TrimSpace(dbChoice)

	var err error
	// Handle Node.js specialization if Node.js is chosen
	if langChoice == "3" {
		fmt.Println("For Node.js, select the specialization:")
		fmt.Println("1) Frontend (React)")
		fmt.Println("2) Mobile (React Native)")
		fmt.Println("3) Backend (Node + Express)")

		nodeChoice, _ := reader.ReadString('\n')
		nodeChoice = strings.TrimSpace(nodeChoice)

		// Generate Node.js setup
		err = utils.HandleNodeSelection(nodeChoice)
	} else {
		// Handle Go or Python setup
		err = utils.HandleBasicStack(langChoice, dbChoice)
	}

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Setup completed successfully!")
}
