# Get goagen src. Probably shouldn't be run by hand
./vendor/github.com/goadesign/goa/goagen/main.go:
	go mod vendor

# Build the goagen executable.
bin/goagen: ./vendor/github.com/goadesign/goa/goagen/main.go
	cd vendor/github.com/goadesign/goa/goagen; go build
	cp vendor/github.com/goadesign/goa/goagen/goagen bin/

# Generate goagen client code
client: bin/goagen
	bin/goagen -d github.com/Intafel/iduvm-api/design client
