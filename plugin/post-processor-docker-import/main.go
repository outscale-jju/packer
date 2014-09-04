package main

import (
	"github.com/outscale/packer/packer/plugin"
	"github.com/outscale/packer/post-processor/docker-import"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(dockerimport.PostProcessor))
	server.Serve()
}
