package ssh

import (
	"backupper/internal/errors"
	"os"
	"path"

	"golang.org/x/crypto/ssh"
)

func NewSSHConfig(user, ip string) (*ssh.ClientConfig, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	keyFile, err := os.ReadFile(path.Join(homeDir, ".ssh/id_ed25519"))
	if err != nil {
		return nil, errors.ErrNoKeyFound
	}

	signer, err := ssh.ParsePrivateKey(keyFile)
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config, nil
}
