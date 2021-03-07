.PHONY: build-react build-go run

go: build
	go build -o 'build/ttt3' 'tic-tac-toe-3/backend/cmd/ttt3'

build:
	mkdir build

all: go react

react: build
	( cd frontend && npm run build )
	rm -rf build/react/
	mv frontend/build/ build/react/

run:
	build/ttt3

