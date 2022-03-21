package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type AafOptions struct {
	InputFilePath  string
	OutputFilePath string
	Verbose        bool
}

func NewDefaultAafCommand() *cobra.Command {
	return NewDefaultAafCommandWithArgs(AafOptions{})
}
func NewDefaultAafCommandWithArgs(o AafOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aaf",
		Short: "anki-autofill",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please use subCmd")
		},
	}

	cmd.PersistentFlags().StringVar(&o.InputFilePath, "input-filepath", "input.txt", "exported text file from Anki")
	cmd.PersistentFlags().StringVar(&o.OutputFilePath, "output-filepath", "output.txt", "output file name to import into Anki")
	cmd.PersistentFlags().BoolVarP(&o.Verbose, "verbose", "v", false, "Print verbose messages")

	cmd.AddCommand(NewDefaultStarDictCommand(o))

	return cmd
}
