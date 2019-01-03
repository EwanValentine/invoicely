GOBUILD=env GOOS=linux go build -ldflags="-s -w" -o

build:
	go get -u
	go mod vendor
	$(GOBUILD) bin/clients functions/clients/main.go
	$(GOBUILD) bin/items functions/items/main.go
	$(GOBUILD) bin/sprints functions/sprints/main.go

test:
	go test ./...

deploy:
	sls deploy