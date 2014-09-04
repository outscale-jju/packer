package main

import (
	"github.com/outscale/packer/command/build"
	"github.com/outscale/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterCommand(new(build.Command))
	server.Serve()
}
