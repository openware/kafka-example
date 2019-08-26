default: build

clean:
	@rm -rf bin/benchmark

build: clean
	@go build -o ./bin/benchmark cmd/benchmark/main.go

deps:
	@docker-compose -f config/docker-compose.yaml up -d