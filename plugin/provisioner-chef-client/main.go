package main

import (
	"github.com/outscale/packer/packer/plugin"
	"github.com/outscale/packer/provisioner/chef-client"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(chefclient.Provisioner))
	server.Serve()
}
