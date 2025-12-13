package ssh

import (
	"fmt"
	"os"
	"path"
	"time"

	_ssh "golang.org/x/crypto/ssh"
)

type Session = _ssh.Session

var (
	max = 100
)

func Run(ip string) (*_ssh.Session, error) {
	attempts := 0
	sshConfig, err := newClientConfig()
	if err != nil {
		return nil, err
	}
	for {
		conn, err := _ssh.Dial("tcp", fmt.Sprintf("%s:%s", ip, "22"), sshConfig)
		if err != nil {
			attempts++
			if attempts >= max {
				return nil, err
			}
			time.Sleep(1 * time.Second)
			fmt.Printf("%v (%d/%d)\n", err, attempts, max)
		} else {
			return conn.NewSession()
		}
	}
}

func newClientConfig() (*_ssh.ClientConfig, error) {
	authMethod, err := authMethod()
	if err != nil {
		return nil, err
	}
	return &_ssh.ClientConfig{
		User: "runner",
		Auth: []_ssh.AuthMethod{
			authMethod,
		},
		HostKeyCallback: _ssh.InsecureIgnoreHostKey(),
	}, nil
}

func authMethod() (_ssh.AuthMethod, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	keyPath := path.Join(home, ".ssh", "id_25519")
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	signer, err := _ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}
	return _ssh.PublicKeys(signer), nil
}
