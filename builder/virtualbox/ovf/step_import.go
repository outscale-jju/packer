package ovf

import (
	"fmt"
	"github.com/mitchellh/multistep"
	vboxcommon "github.com/outscale/packer/builder/virtualbox/common"
	"github.com/outscale/packer/packer"
)

// This step imports an OVF VM into VirtualBox.
type StepImport struct {
	Name       string
	SourcePath string
	ImportOpts string

	vmName string
}

func (s *StepImport) Run(state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(vboxcommon.Driver)
	ui := state.Get("ui").(packer.Ui)

	ui.Say(fmt.Sprintf("Importing VM: %s", s.SourcePath))
	if err := driver.Import(s.Name, s.SourcePath, s.ImportOpts); err != nil {
		err := fmt.Errorf("Error importing VM: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	s.vmName = s.Name
	state.Put("vmName", s.Name)
	return multistep.ActionContinue
}

func (s *StepImport) Cleanup(state multistep.StateBag) {
	if s.vmName == "" {
		return
	}

	driver := state.Get("driver").(vboxcommon.Driver)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Unregistering and deleting imported VM...")
	if err := driver.Delete(s.vmName); err != nil {
		ui.Error(fmt.Sprintf("Error deleting VM: %s", err))
	}
}
