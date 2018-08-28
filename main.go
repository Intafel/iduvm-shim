package main

import (
	"C" // Required for export to be recognized

	"context"
	"fmt"

	"github.com/Intafel/iduvm-shim/client"
)

//export OpenBrowser
func ISTOpenBrowser(guestID string) error {
	c := client.New(nil)
	path := client.OpenBrowserGuestPath(guestID)
	resp, err := c.OpenBrowserGuest(context.TODO(), path)
	fmt.Println(resp)
	fmt.Println(err)
	return err
}

//export OpenFile
func ISTOpenFile(guestID string, filePath string) error {
	c := client.New(nil)
	path := client.OpenFileGuestPath(guestID, filePath)
	resp, err := c.OpenFileGuest(context.TODO(), path)
	fmt.Println(resp)
	fmt.Println(err)
	return err
}

//export OpenURL
func ISTOpenURL(guestID string, url_ string) error {
	c := client.New(nil)
	path := client.OpenURLGuestPath(guestID, url_)
	resp, err := c.OpenURLGuest(context.TODO(), path)
	fmt.Println(resp)
	fmt.Println(err)
	return err
}

//export StartGuest
func ISTStartGuest(guestID string) error {
	c := client.New(nil)
	path := client.StartGuestPath(guestID)
	resp, err := c.StartGuest(context.TODO(), path)
	fmt.Println(resp)
	fmt.Println(err)
	return err
}

//export StopGuest
func ISTStopGuest(guestID string) error {
	c := client.New(nil)
	path := client.StopGuestPath(guestID)
	resp, err := c.StopGuest(context.TODO(), path)
	fmt.Println(resp)
	fmt.Println(err)
	return err
}

func main() {}
