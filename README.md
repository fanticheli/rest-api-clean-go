# rest-api-clean-go
Training skills to create a restApi in golang and PostgresSQL, using hexagonal architecture.

# Run application
    docker-compose up
    docker exec -it rest-api-clean-go /bin/bash
    go mod tidy
    go run adapter/http/main.go

# Test application
    Run test: go test ./...
    Run test verbose: go test ./... -v
    Run test cov out: go test -coverprofile cover.out ./...
    Run test html output: go tool cover -html=cover.out -o cover.html

