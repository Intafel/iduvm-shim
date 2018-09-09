package main

import (
	"C" // Required for export to be recognized

	"context"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"os/exec"
	
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/Intafel/iduvm-shim/client"
	"github.com/Intafel/iduvm-shim/plugin"
)

// handshakeConfig is copied from the basic example from plugins 
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

/**
 * Sets up the default plugin logger
 */
func setupLogger() hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		Name: "plugin",
		Output: os.Stdout,
		Level: hclog.Debug,
	})
}

/**
 * First tries to read the path stored at environment variable
 * $IDUVM_PLUGIN_CONFIG, otherwise reads ~/.iduvm/plugin.conf
 */
func getPluginList() iplugin.PluginList {
	var log = setupLogger()
	// Get the path
	path, present := os.LookupEnv("IDUVM_PLUGIN_CONFIG")
	if !present {
		path = "~/.iduvm/plugin.conf"
	}
	// Test if the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Error("Could not find plugin configuration file. Not running plugins.", path)
		return iplugin.PluginList{}
	}
	file, err := ioutil.ReadFile(path)
	log.Debug("Read plugin file at path: ", path)
	if err != nil {
		log.Error("Failed to read file", err)
		return iplugin.PluginList{}
	}
	// Test if it's good JSON
	if !json.Valid(file) {
		log.Error("Plugin configuration file is malformed. Not running plugins.")
		return iplugin.PluginList{}
	}
	// Fill the PluginList
	pluginList := iplugin.PluginList{}
	err = json.Unmarshal(file, &pluginList)
	if err != nil {
		log.Error("Failed to unmarshal plugins file", err)
		return iplugin.PluginList{}
	}
	log.Info("Running the follow plugins:")
	for _, p := range pluginList.Plugins {
		log.Info(p)
	}
	return pluginList
}

/**
 * Run a client for each plugin in the plugin list and call the plugin
 */
func callPlugins(resource string) {
	var log = setupLogger()
	// There should only be a decider plugin provided by the plugin
	var pluginMap = map[string]plugin.Plugin{
		"decider": &iplugin.DeciderPlugin{},
	}
	pluginList := getPluginList()
	for _, p := range pluginList.Plugins {
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: handshakeConfig,
			Plugins:         pluginMap,
			Cmd:             exec.Command(p),
			Logger:          log,
		})
		defer client.Kill()
		// Connect via RPC
		rpcClient, err := client.Client()
		if err != nil {
			log.Error("Failed to instantiate rpc client", err.Error())
			return
		}
		// Request the plugin
		raw, err := rpcClient.Dispense("decider")
		if err != nil {
			log.Error("Failed to get decider plugin", err.Error())
			return
		}
		decider := raw.(iplugin.Decider)
		decider.Decide(resource)
	}
}

//export ISTOpenBrowser
func ISTOpenBrowser(guestID string) error {
	c := client.New(nil)
	path := client.OpenBrowserGuestPath(guestID)
	resp, err := c.OpenBrowserGuest(context.TODO(), path)
	fmt.Println(resp)
	fmt.Println(err)
	return err
}

//export ISTOpenFile
func ISTOpenFile(guestID string, filePath string) error {
	c := client.New(nil)
	path := client.OpenFileGuestPath(guestID, filePath)
	resp, err := c.OpenFileGuest(context.TODO(), path)
	fmt.Println(resp)
	fmt.Println(err)
	return err
}

//export ISTOpenURL
func ISTOpenURL(guestID string, url_ string) error {
	c := client.New(nil)
	path := client.OpenURLGuestPath(guestID, url_)
	resp, err := c.OpenURLGuest(context.TODO(), path)
	fmt.Println(resp)
	fmt.Println(err)
	return err
}

//export ISTStartGuest
func ISTStartGuest(guestID string) error {
	c := client.New(nil)
	path := client.StartGuestPath(guestID)
	resp, err := c.StartGuest(context.TODO(), path)
	fmt.Println(resp)
	fmt.Println(err)
	return err
}

//export ISTStopGuest
func ISTStopGuest(guestID string) error {
	c := client.New(nil)
	path := client.StopGuestPath(guestID)
	resp, err := c.StopGuest(context.TODO(), path)
	fmt.Println(resp)
	fmt.Println(err)
	return err
}

func main() {
	callPlugins("blah.com")
}
