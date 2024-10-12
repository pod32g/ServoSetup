package utils

import (
	"fmt"
	"os"
)

// Node.js package.json templates
var ReactPackageJSON = `{
	"name": "react-app",
	"version": "1.0.0",
	"scripts": {
	  "start": "react-scripts start"
	},
	"dependencies": {
	  "react": "^17.0.2",
	  "react-dom": "^17.0.2"
	}
  }`

var ReactNativePackageJSON = `{
	"name": "react-native-app",
	"version": "1.0.0",
	"dependencies": {
	  "react": "^17.0.2",
	  "react-native": "^0.64.0"
	}
  }`

var NodeExpressPackageJSON = `{
	"name": "express-app",
	"version": "1.0.0",
	"main": "index.js",
	"scripts": {
	  "start": "node index.js"
	},
	"dependencies": {
	  "express": "^4.17.1"
	}
  }`

// Docker Compose templates for databases
var PostgresDockerCompose = `version: '3'
services:
  db:
    image: postgres
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: mydb
    ports:
      - "5432:5432"
`

var MongoDockerCompose = `version: '3'
services:
  db:
    image: mongo
    ports:
      - "27017:27017"
`

var SqliteDockerCompose = `version: '3'
services:
  db:
    image: nouchka/sqlite3
    volumes:
      - ./data:/data
`

var GoModContent = `module my-go-project

go 1.19
`

var PythonRequirements = `flask==2.0.1
`

// Helper function to create a file with error handling
func CreateFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fileName, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", fileName, err)
	}

	fmt.Printf("Successfully created %s\n", fileName)
	return nil
}

// Handle Node.js setup with enhanced error handling
func HandleNodeSelection(choice string) error {
	var err error
	switch choice {
	case "1":
		fmt.Println("Generating React setup...")
		err = CreateFile("package.json", ReactPackageJSON)
	case "2":
		fmt.Println("Generating React Native setup...")
		err = CreateFile("package.json", ReactNativePackageJSON)
	case "3":
		fmt.Println("Generating Node + Express setup...")
		err = CreateFile("package.json", NodeExpressPackageJSON)

		// Add database-specific dependencies for Node.js
		if err == nil {
			// Assuming PostgreSQL as the default for backend
			err = RunCommandWithProgress("npm", "install", "pg")
		}
	default:
		return fmt.Errorf("invalid Node.js selection")
	}

	if err != nil {
		return err
	}

	fmt.Println("Installing dependencies...")
	err = RunCommandWithProgress("npm", "install")
	if err != nil {
		return err
	}

	// Generate the .env file for Node.js projects as well
	err = CreateEnvFile("1") // Example using PostgreSQL for backend projects
	return err
}

// Handle Go or Python setup with enhanced error handling
func HandleBasicStack(lang, db string) error {
	var err error
	switch lang {
	case "1": // Go
		fmt.Println("Setting up Go project...")
		err = CreateFile("go.mod", GoModContent)
		if err == nil {
			fmt.Println("Installing Go dependencies...")

			// Add database-specific dependencies to Go project
			switch db {
			case "1": // PostgreSQL
				err = RunCommandWithProgress("go", "get", "github.com/lib/pq")
			case "2": // MongoDB
				err = RunCommandWithProgress("go", "get", "go.mongodb.org/mongo-driver/mongo")
			case "3": // SQLite
				err = RunCommandWithProgress("go", "get", "github.com/mattn/go-sqlite3")
			}
			if err == nil {
				err = RunCommandWithProgress("go", "mod", "tidy")
			}
		}
	case "2": // Python
		fmt.Println("Setting up Python project...")
		err = CreateFile("requirements.txt", PythonRequirements)

		if err == nil {
			// Add database-specific dependencies to Python
			switch db {
			case "1": // PostgreSQL
				err = AppendToFile("requirements.txt", "psycopg2-binary==2.9.1\n")
			case "2": // MongoDB
				err = AppendToFile("requirements.txt", "pymongo==3.12.0\n")
			case "3": // SQLite
				err = AppendToFile("requirements.txt", "sqlite3==3.31.1\n")
			}
			if err == nil {
				fmt.Println("Installing Python dependencies...")
				err = RunCommandWithProgress("pip", "install", "-r", "requirements.txt")
			}
		}
	default:
		return fmt.Errorf("invalid language selection")
	}

	if err != nil {
		return err
	}

	// Generate Docker Compose file and .env file for the selected database
	err = GenerateDockerCompose(db)
	return err
}

// Generate Docker Compose file with error handling
func GenerateDockerCompose(db string) error {
	var err error

	// Generate Docker Compose file based on the selected database
	switch db {
	case "1":
		fmt.Println("Setting up PostgreSQL...")
		err = CreateFile("docker-compose.yml", PostgresDockerCompose)
	case "2":
		fmt.Println("Setting up MongoDB...")
		err = CreateFile("docker-compose.yml", MongoDockerCompose)
	case "3":
		fmt.Println("Setting up SQLite...")
		err = CreateFile("docker-compose.yml", SqliteDockerCompose)
	default:
		return fmt.Errorf("invalid database selection")
	}
	if err != nil {
		return err
	}

	// Generate the .env file
	err = CreateEnvFile(db)
	return err
}

// Function to create an .env file
func CreateEnvFile(db string) error {
	envContent := ""

	// Add environment variables based on the selected database
	switch db {
	case "1": // PostgreSQL
		envContent = `DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=mydb
`
	case "2": // MongoDB
		envContent = `DB_HOST=localhost
DB_PORT=27017
DB_USER=user
DB_PASSWORD=password
`
	case "3": // SQLite
		envContent = `DB_PATH=./data/database.db
`
	default:
		return fmt.Errorf("invalid database selection")
	}

	// Create the .env file with the environment variables
	err := CreateFile(".env", envContent)
	if err != nil {
		return fmt.Errorf("failed to create .env file: %w", err)
	}

	fmt.Println("Successfully created .env file")
	return nil
}

func AppendToFile(fileName, content string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", fileName, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to append to file %s: %w", fileName, err)
	}

	return nil
}
