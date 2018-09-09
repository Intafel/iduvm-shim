package iplugin

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type PluginList struct {
	Plugins []string
}

/**
 * A plugin that makes a decision as to how to open a resource
 */
type Decider interface {
	// Make a decision as to how to open a resource
	// resource is either a path to a file or a URL
	Decide(resource string) string
}

type DeciderRPCClient struct {
	client *rpc.Client
}

func (d *DeciderRPCClient) Decide(resource string) string {
	var resp string
	err := d.client.Call("Plugin.Decide", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

type DeciderRPCServer struct {
	Impl Decider
}

func (d *DeciderRPCServer) Decide(args interface{}, resp *string) error {
	*resp = d.Impl.Decide("blah.com")
	return nil
}

type DeciderPlugin struct {
	Impl Decider
}

func (p *DeciderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &DeciderRPCServer{Impl: p.Impl}, nil
}


func (p DeciderPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &DeciderRPCClient{client: c}, nil
}
