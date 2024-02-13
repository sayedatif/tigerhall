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
    DB_HOST=localhost
    DB_NAME=
    SECRET_KEY=
5. Install dependencies
   `go mod download`

## Create Database in MySQL

1. Login as root `mysql -u root -p`
2. Enter password
4. Create database using `Create Database tigerhall;`
5. Use database using `use tigerhall;`

## Usage on localhost

Instructions for using the project.

1. Make sure you have add env file and MySQL server is running.
2. Run `go run main.go`
3. Run test `go test ./test/... -v` I have test cases in test folder

## Usage via Docker

Instructions for using the project via docker.

1. Add env file
   ```bash
    PORT=8080
    DB_USER=
    DB_PASSWORD=
    DB_HOST=172.18.0.2
    DB_NAME=
    SECRET_KEY=
3. Build docker image: `docker build -t tigerhall-backend .`
4. Run docker compose
   ```bash
   export MYSQL_ROOT_PASSWORD=root@123
   docker-compose up -d

## Using Makefile

1. Build the project: `make build`
2. Start the server: `make run`
3. Run test cases: `make test`

## Access

1. Open a web browser and go to http://localhost:8080 to access the application.

## Swager docs

1. Build `swag init`
2. Access `http://localhost:8080/swagger/index.html`
