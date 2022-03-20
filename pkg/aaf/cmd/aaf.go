package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type AafOptions struct {
}

func NewDefaultAafCommand() *cobra.Command {
	return NewDefaultAafCommandWithArgs(AafOptions{})
}
func NewDefaultAafCommandWithArgs(o AafOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aaf",
		Short: "anki-autofill",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Run aaf root command")
		},
	}

	return cmd
}
