package qemu

import (
	gossh "code.google.com/p/go.crypto/ssh"
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/outscale/packer/communicator/ssh"
	"io/ioutil"
	"os"
)

func sshAddress(state multistep.StateBag) (string, error) {
	sshHostPort := state.Get("sshHostPort").(uint)
	return fmt.Sprintf("127.0.0.1:%d", sshHostPort), nil
}

func sshConfig(state multistep.StateBag) (*gossh.ClientConfig, error) {
	config := state.Get("config").(*config)

	auth := []gossh.AuthMethod{
		gossh.Password(config.SSHPassword),
		gossh.KeyboardInteractive(
			ssh.PasswordKeyboardInteractive(config.SSHPassword)),
	}

	if config.SSHKeyPath != "" {
		signer, err := sshKeyToSigner(config.SSHKeyPath)
		if err != nil {
			return nil, err
		}

		auth = append(auth, gossh.PublicKeys(signer))
	}

	return &gossh.ClientConfig{
		User: config.SSHUser,
		Auth: auth,
	}, nil
}

func sshKeyToSigner(path string) (gossh.Signer, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	keyBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	signer, err := gossh.ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("Error setting up SSH config: %s", err)
	}

	return signer, nil
}
