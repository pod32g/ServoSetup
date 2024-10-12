
# ServoSetup

ServoSetup is a CLI tool designed to quickly set up development environments by selecting various technologies or stacks like Go, Python, Node.js, MongoDB, PostgreSQL, SQLite, and Docker. It automatically generates configuration files, installs dependencies, and sets up Docker if needed, helping you jumpstart your development with minimal effort.

## Features
- **Choose your language**: Go, Python, or Node.js
- **Choose your database**: PostgreSQL, MongoDB, or SQLite
- **Generates environment variables**: Automatically creates a `.env` file based on the selected stack
- **Automatic dependency installation**: Installs dependencies based on your stack choices
- **Supports Docker Compose**: Optionally generates Docker Compose files for database setup
- **Verbose mode**: Control command outputs using the `--verbose` flag
- **Progress bars**: Visual feedback for long-running tasks like installing dependencies

## Usage

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/servosetup.git
    cd servosetup
    ```

2. Build the CLI:
    ```bash
    go build -o servosetup
    ```

3. Run the CLI interactively:
    ```bash
    ./servosetup
    ```

4. Use the `--verbose` flag to see detailed output:
    ```bash
    ./servosetup --verbose
    ```

### Example Workflow
1. Select your language (Go, Python, Node.js)
2. Select your database (PostgreSQL, MongoDB, SQLite)
3. The tool generates:
    - Configuration files (`go.mod`, `requirements.txt`, `package.json`, `docker-compose.yml`)
    - `.env` file with database credentials
4. Automatically installs dependencies

### Dependencies

- **Go**
    - PostgreSQL: `github.com/lib/pq`
    - MongoDB: `go.mongodb.org/mongo-driver/mongo`
    - SQLite: `github.com/mattn/go-sqlite3`
    
- **Python**
    - PostgreSQL: `psycopg2-binary==2.9.1`
    - MongoDB: `pymongo==3.12.0`
    - SQLite: `sqlite3==3.31.1`
    
- **Node.js**
    - PostgreSQL: `pg`
    - MongoDB: `mongoose`
    - SQLite: `sqlite3`

## Name Origin

The name ServoSetup is inspired by the servitors of the Warhammer 40,000 universe—cybernetic beings who tirelessly perform essential tasks with precision and efficiency. Similarly, ServoSetup automates the setup of development environments, executing complex tasks in a streamlined and reliable manner, allowing developers to focus on building their projects quickly and effortlessly.

## Contributing

Feel free to contribute by opening issues or pull requests to improve ServoSetup.

## License

This project is licensed under the MIT License.