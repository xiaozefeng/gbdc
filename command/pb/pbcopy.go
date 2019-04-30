package command

import (
	"fmt"

	"github.com/xiaozefeng/gbdc/command"
)

func Copy(content string) error {
	cmd := fmt.Sprintf("echo %s | pbcopy", content)
	_, err := command.Run(cmd)
	return err
}
