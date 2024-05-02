game: 
	@go run ./cmd/game

build:
	@go build -o ./bin/minesweeper ./cmd/minesweeper

install:
	@go install ./cmd/minesweeper
