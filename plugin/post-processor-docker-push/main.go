package main

import (
	"github.com/outscale/packer/packer/plugin"
	"github.com/outscale/packer/post-processor/docker-push"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(dockerpush.PostProcessor))
	server.Serve()
}
