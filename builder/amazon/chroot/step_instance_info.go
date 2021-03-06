package chroot

import (
	"fmt"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
	"github.com/mitchellh/multistep"
	"github.com/outscale/packer/packer"
	"log"
)

// StepInstanceInfo verifies that this builder is running on an EC2 instance.
type StepInstanceInfo struct{}

func (s *StepInstanceInfo) Run(state multistep.StateBag) multistep.StepAction {
	ec2conn := state.Get("ec2").(*ec2.EC2)
	ui := state.Get("ui").(packer.Ui)

	// Get our own instance ID
	ui.Say("Gathering information about this EC2 instance...")
	instanceIdBytes, err := aws.GetMetaData("instance-id")
	if err != nil {
		log.Printf("Error: %s", err)
		err := fmt.Errorf(
			"Error retrieving the ID of the instance Packer is running on.\n" +
				"Please verify Packer is running on a proper AWS EC2 instance.")
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	instanceId := string(instanceIdBytes)
	log.Printf("Instance ID: %s", instanceId)

	// Query the entire instance metadata
	instancesResp, err := ec2conn.Instances([]string{instanceId}, ec2.NewFilter())
	if err != nil {
		err := fmt.Errorf("Error getting instance data: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if len(instancesResp.Reservations) == 0 {
		err := fmt.Errorf("Error getting instance data: no instance found.")
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	instance := &instancesResp.Reservations[0].Instances[0]
	state.Put("instance", instance)

	return multistep.ActionContinue
}

func (s *StepInstanceInfo) Cleanup(multistep.StateBag) {}
