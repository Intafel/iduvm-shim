bin/goagen: ./vendor/github.com/goadesign/goa/goagen/main.go
	cd vendor/github.com/goadesign/goa/goagen; go build
	cp vendor/github.com/goadesign/goa/goagen/goagen bin/

