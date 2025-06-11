package ssh

import "golang.org/x/crypto/ssh"

func RunCommand(connection *ssh.ClientConfig, host, command string) (string, error) {
	session, err := NewSession(connection, host)
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
