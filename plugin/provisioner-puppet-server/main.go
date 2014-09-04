package main

import (
	"github.com/outscale/packer/packer/plugin"
	"github.com/outscale/packer/provisioner/puppet-server"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(puppetserver.Provisioner))
	server.Serve()
}
