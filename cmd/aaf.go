package main

import (
	"fmt"

	aaf "github.com/makerdark98/aaf/pkg/aaf/cmd"
)

func main() {
	cmd := aaf.NewDefaultAafCommand()
	err := cmd.Execute()
	if err != nil {
		fmt.Errorf("command failed %v\n", err)
	}
}
