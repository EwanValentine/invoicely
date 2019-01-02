GOBUILD=env GOOS=linux go build -ldflags="-s -w" -o

build:
	go get -u ./...
	go mod vendor
	$(GOBUILD) bin/create-client functions/create-client/main.go
	$(GOBUILD) bin/fetch-clients functions/fetch-clients/main.go
	$(GOBUILD) bin/fetch-client functions/fetch-client/main.go

	$(GOBUILD) bin/create-sprint functions/create-sprint/main.go
	$(GOBUILD) bin/fetch-sprints function/fetch-sprints/main.go
	$(GOBUILD) bin/fetch-sprint functions/fetch-sprint/main.go

	$(GOBUILD) bin/create-item functions/create-item/main.go
	$(GOBUILD) bin/fetch-items functions/fetch-items/main.go
	$(GOBUILD) bun/fetch-item functions/fetch-item/main.go

test:
	go test ./...