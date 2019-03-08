package data

import (
	ui "github.com/sqshq/termui"
	"os/exec"
	"strings"
)

type Item struct {
	Label  string
	Script string
	Color  ui.Color
}

func (i *Item) nextValue() (value string, err error) {

	output, err := exec.Command("sh", "-c", i.Script).Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
