package command

import (
	"strings"

	"github.com/xiaozefeng/gbdc/command"
)

func Paste() (string, error) {
	content, err := command.Run("pbpaste")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(content), nil
}
