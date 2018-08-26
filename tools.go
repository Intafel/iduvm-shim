// +build tools
package main

// See https://github.com/golang/go/wiki/Modules#faq
// This recommends tracking tools in a file such as this
import (
	// For explanations of goa dependencies, see Solitare 2
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goagen"
	"github.com/goadesign/goa/goagen/gen_app"
	"github.com/goadesign/goa/goagen/gen_client"
	"github.com/goadesign/goa/goagen/gen_main"
	"github.com/Intafel/iduvm-api/design"
)
