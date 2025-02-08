VERSION=0.1.0
LDFLAGS=-ldflags "-w -s -X main.version=${VERSION}"

all: check-dns-multi

.PHONY: check-dns-multi

check-dns-multi: cmd/check-dns-multi/main.go
	go build $(LDFLAGS) -o check-dns-multi cmd/check-dns-multi/main.go

linux: cmd/check-dns-multi/main.go
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o check-dns-multi cmd/check-dns-multi/main.go

check:
	go test ./...


