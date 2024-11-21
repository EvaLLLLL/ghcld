BINARY_NAME=ghcld
GOOS=linux
GOARCH=amd64

all: build

build:
	GOOS=linux GOARCH=$(GOARCH) go build -o $(BINARY_NAME)_linux_$(GOARCH) .
	GOOS=darwin GOARCH=$(GOARCH) go build -o $(BINARY_NAME)_macos_$(GOARCH) .
	GOOS=windows GOARCH=$(GOARCH) go build -o $(BINARY_NAME)_windows_$(GOARCH).exe .

clean:
	rm -f $(BINARY_NAME)_linux_$(GOARCH) $(BINARY_NAME)_macos_$(GOARCH) $(BINARY_NAME)_windows_$(GOARCH).exe

.PHONY: all build clean