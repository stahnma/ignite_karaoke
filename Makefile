

all: fmt
	go mod tidy
	go build .

fmt:
	go fmt *.go

clean:
	rm -f pkg/templates.go
	rm -f ignite_karaoke
