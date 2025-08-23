GOCMD = go
SWAG = swag

build:
	$(GOCMD) build -o bin/main cmd/server/main.go

run:
	$(GOCMD) run cmd/server/main.go

swagger:
	$(SWAG) init --dir=cmd/server,internal/interfaces/http

# Run linters and checks
vet:
	$(GOCMD) vet ./...


# Format code
fmt:
	$(GOCMD) fmt ./...


