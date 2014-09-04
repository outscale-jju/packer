package main

import (
	"github.com/outscale/packer/packer/plugin"
	"github.com/outscale/packer/provisioner/salt-masterless"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(saltmasterless.Provisioner))
	server.Serve()
}
