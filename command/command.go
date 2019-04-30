package command

import (
	"errors"
	"os/exec"
)

func Run(command string) (string, error) {
	if command == "" {
		return "", errors.New("command is empty")
	}
	cmd := exec.Command("sh", "-c", command)
	bytes, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
