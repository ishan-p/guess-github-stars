.PHONY = all clean
all: build

build: main.go
	@echo "Building binary..."
	go build -o guess-github-stars.o

clean:
	@echo "Cleaning up..."
	rm guess-github-stars.o
	go clean
