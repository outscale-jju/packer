package null

import (
	"github.com/outscale/packer/packer"
	"testing"
)

func TestBuilder_implBuilder(t *testing.T) {
	var _ packer.Builder = new(Builder)
}
