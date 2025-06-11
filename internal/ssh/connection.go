package ssh

import "golang.org/x/crypto/ssh"

func NewSession(config *ssh.ClientConfig, host string) (*ssh.Session, error) {
	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}
