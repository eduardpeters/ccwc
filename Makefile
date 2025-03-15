main_package_path = ./cmd
binary_name = ccwc

.PHONY: test
test:
	go test ./internal

.PHONY: build
build:
	go build -o ${binary_name} ${main_package_path}

.PHONY: run
run: build
	./${binary_name}

.PHONY: clean
clean:
	rm ${binary_name}