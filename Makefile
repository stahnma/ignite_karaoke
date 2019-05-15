

all: fmt
	go build .

fmt:
	go fmt *.go

clean:
	rm pkg/templates.go
