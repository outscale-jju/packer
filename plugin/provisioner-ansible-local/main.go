package main

import (
	"github.com/outscale/packer/packer/plugin"
	"github.com/outscale/packer/provisioner/ansible-local"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(ansiblelocal.Provisioner))
	server.Serve()
}
