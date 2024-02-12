## Installation

Instructions for installing the project.

1. Clone the repository:
   `git clone git@github.com:sayedatif/tigerhall.git`
2. Navigate to the project directory:
  `cd tigerhall`
3. Add env file
   ```bash
    PORT=8080
    DB_USER=
    DB_PASSWORD=
    DB_HOST=
    DB_NAME=
    SECRET_KEY=
5. Install dependencies
   `go mod download`

## Usage on localhost

Instructions for using the project.

1. Make sure you have add env file and MySQL server is running.
2. Run `go run main.go`
3. Open a web browser and go to http://localhost:8080 to access the application.

## Usage via Docker

Instructions for using the project via docker.

1. Build docker image: `docker build -t tigerhall-backend .`
2. Run docker compose
   ```bash
   export MYSQL_ROOT_PASSWORD=root@123
   docker-compose up -d
3. Open a web browser and go to http://localhost:8080 to access the application.

## Using Makefile

1. Build the project: `make build`
2. Start the server: `make run`
3. Open a web browser and go to http://localhost:8080 to access the application.
