# Golang API Template

This repository serves as a template for building a RESTful API using Golang. It provides a structured layout with predefined folders and files to help you kickstart your API development.

## Project Structure

```plaintext
.
├── main.go
├── config
├── controller
├── database
│   ├── index.go
│   ├── migration
│   └── seeder
├── middleware
├── public
├── routes
└── utils
```

## Folder and File Descriptions

- **main.go**: The entry point of the application.
- **config**: Contains configuration files for the application.
- **controller**: Contains the logic for handling requests and sending responses.
- **database**: Contains database related files and folders.
  - **index.go**: Database connection and setup.
  - **migration**: Folder for database migration files.
  - **seeder**: Folder for database seed files.
- **middleware**: Contains middleware functions for the application.
- **public**: Publicly accessible files (e.g., images, CSS, JavaScript).
- **routes**: Contains route definitions and mappings.
- **utils**: Utility functions and helpers.

## Getting Started

### Prerequisites

- [golang](https://golang.org/dl/) installed (version 1.17 or higher recommended)
- [Docker](https://www.docker.com/) installed (optional, for containerized development)

### Installation

1. Clone the repository:

   ```bash
    git clone
    cd golang-template
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up the database:
   - Ensure your database is running
   - Copy <code>.env.example</code> and save as <code>.env</code> file
   - Congifure the database connection in the <code>.env</code> file such as **username**, **password** and **database name**.
   - Migrations and seeders will be automatically when the file <code>main.go</code> is run.
4. Run the application:
   ```bash
   go run main.go
   ```

### Docker Setup (Optional)

To run the application in a Docker container, follow these steps:

1. Build the Docker image:
   ```bash
   docker build -t golang-api-template .
   ```
2. Run the Docker container:
   ```bash
   docker run -p 8080:8000 golang-template
   ```
