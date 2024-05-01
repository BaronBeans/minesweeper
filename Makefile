test:
	@go run ./cmd/test

game: 
	@go run ./cmd/game

build:
	@go build -o ./bin/game ./cmd/game

install:
	@go install ./cmd/game
