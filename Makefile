build:
	go build -o ./bin/app

# ex. ARGS="--port 8080" make run
run: build
	./bin/app $(ARGS)
