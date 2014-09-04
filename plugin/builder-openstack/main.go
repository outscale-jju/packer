package main

import (
	"github.com/outscale/packer/builder/openstack"
	"github.com/outscale/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(openstack.Builder))
	server.Serve()
}
