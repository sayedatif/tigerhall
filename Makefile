build:
	go build main.go

run: build
	./main

clean:
	rm -f main