package main

import (
	"github.com/outscale/packer/packer/plugin"
	"github.com/outscale/packer/post-processor/vsphere"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(vsphere.PostProcessor))
	server.Serve()
}
