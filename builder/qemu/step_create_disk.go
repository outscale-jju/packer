package qemu

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/outscale/packer/packer"
	"path/filepath"
	"strings"
)

// This step creates the virtual disk that will be used as the
// hard drive for the virtual machine.
type stepCreateDisk struct{}

func (s *stepCreateDisk) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*config)
	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packer.Ui)
	path := filepath.Join(config.OutputDir, fmt.Sprintf("%s.%s", config.VMName,
		strings.ToLower(config.Format)))

	command := []string{
		"create",
		"-f", config.Format,
		path,
		fmt.Sprintf("%vM", config.DiskSize),
	}

	ui.Say("Creating hard drive...")
	if err := driver.QemuImg(command...); err != nil {
		err := fmt.Errorf("Error creating hard drive: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *stepCreateDisk) Cleanup(state multistep.StateBag) {}
