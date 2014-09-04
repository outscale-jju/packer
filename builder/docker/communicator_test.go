package docker

import (
	"github.com/outscale/packer/packer"
	"testing"
)

func TestCommunicator_impl(t *testing.T) {
	var _ packer.Communicator = new(Communicator)
}
