build:
	@go build -o bin/authentication-go

run: build
	@./bin/authentication-go

dev:
	@gin run *go