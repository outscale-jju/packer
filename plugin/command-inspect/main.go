package main

import (
	"github.com/outscale/packer/command/inspect"
	"github.com/outscale/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterCommand(new(inspect.Command))
	server.Serve()
}
