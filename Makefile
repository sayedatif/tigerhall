.PHONY: build
build:
	go build main.go

run: build
	./main

clean:
	rm -f main

.PHONY: test
test:
	go test -v ./test/...