package main

import (
	"C" // Required for export to be recognized

	"context"
	"fmt"

	"github.com/Intafel/iduvm-shim/client"
)

//export OpenBrowser
func OpenBrowser(guestID string) error {
	c := client.New(nil)
	path := client.OpenBrowserGuestPath(guestID)
	resp, err := c.OpenBrowserGuest(context.TODO(), path)
	fmt.Println(resp)
	fmt.Println(err)
	return err
}

func main() {}
