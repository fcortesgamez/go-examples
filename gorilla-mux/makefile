NO_VENDOR := `go list ./... | grep -v 'vendor/'`

build: test
	go build main.go app.go model.go

build_skip_tests:
	go build main.go app.go model.go

test:
	go test github.com/fcortesgamez/go-examples/gorilla-mux/...

sanitize:
	go fmt $(NO_VENDOR)
	go vet $(NO_VENDOR)

clean:
	rm -f main
