.PHONY: build clean test lint

build:
	go build -o contractEngineEdge ./

clean:
	${RM} contractEngineEdge

test:
	go test ./...

lint:
	golangci-lint run ./...
	go vet ./...

coverage:
	go test -covermode=count -coverprofile cov --tags unit ./...
	go tool cover -html=cov -o coverage.html

race:
	go test -short -race ./...

docker:
	docker build -t contract_engine_edge -f Dockerfile .
