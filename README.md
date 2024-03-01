# lets-go-snippetbox
This is Demo project build from lets-go book


## How to run the project (dev-mode)

### Pre-requisite
- Go 
- Docker
- Docker-compose

### Steps

1. Download the dependencies
```
go mod download
```

2. Generate TLS certificates
```
make generate_tls
```

3. Start the database
```
make start_db
```

4. Start the server
```
make start_server
```

5. Open the browser and navigate to `http://localhost:4000`

> Above commands will use default command line arguments.


## Using command line arguments

You can use the following command line arguments to start the server

```
go mod run ./cmd/web -addr=":4000" -dsn="host=localhost port=5432 user=snippetbox dbname=snippetbox sslmode=disable"
```

You can check command line arguments by running the following command
```
go mod run ./cmd/web -help
```