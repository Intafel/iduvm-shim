# Get goagen src. Probably shouldn't be run by hand
# As this depends on go mod, if this repository is  on the GOPATH
# you need to set GO111MODULE=on
./vendor/github.com/goadesign/goa/goagen/main.go:
	go mod vendor

# Build the goagen executable.
bin/goagen: ./vendor/github.com/goadesign/goa/goagen/main.go
	cd vendor/github.com/goadesign/goa/goagen; go build
	/usr/bin/cp vendor/github.com/goadesign/goa/goagen/goagen bin/

# Generate goagen client code
client: bin/goagen
	bin/goagen -d github.com/Intafel/iduvm-api/design client

# Build a shared library for Darwin.
bin/client.dylib: main.go client
	go build -v -o bin/client.dylib -buildmode=c-shared main.go
