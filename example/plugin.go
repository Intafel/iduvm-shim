package main

import (
	"os"
	
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/Intafel/iduvm-shim/plugin"
)

type HelloDecider struct {
	logger hclog.Logger
}

func (d *HelloDecider) Decide(resource string) string {
	d.logger.Error("HELLO WORLDXXXXX")
	return "Hello"
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level: hclog.Trace,
		Output: os.Stderr,
		JSONFormat: true,
	})

	decider := &HelloDecider {
		logger: logger,
	}

	var pluginMap = map[string]plugin.Plugin{
		"decider": &iplugin.DeciderPlugin{Impl: decider},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins: pluginMap,
	})
}
