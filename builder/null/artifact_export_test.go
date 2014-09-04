package null

import (
	"github.com/outscale/packer/packer"
	"testing"
)

func TestNullArtifact(t *testing.T) {
	var _ packer.Artifact = new(NullArtifact)
}
