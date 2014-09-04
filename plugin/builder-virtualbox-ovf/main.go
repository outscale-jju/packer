package main

import (
	"github.com/outscale/packer/builder/virtualbox/ovf"
	"github.com/outscale/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(ovf.Builder))
	server.Serve()
}
