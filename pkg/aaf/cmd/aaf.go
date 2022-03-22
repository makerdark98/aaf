package cmd

import (
	"fmt"

	ankiLoader "github.com/makerdark98/aaf/pkg/aaf/anki/loader"
	ankiSaver "github.com/makerdark98/aaf/pkg/aaf/anki/saver"
	"github.com/spf13/cobra"
)

type AafOptions struct {
	InputFilePath  string
	OutputFilePath string
	fileFormat     string
	Loader         ankiLoader.Interface
	Saver          ankiSaver.Interface
	Verbose        bool
}

func NewDefaultAafCommand() *cobra.Command {
	return NewDefaultAafCommandWithArgs(&AafOptions{})
}

func NewDefaultAafCommandWithArgs(o *AafOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aaf",
		Short: "anki-autofill",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := o.Complete()
			if err != nil {
				fmt.Printf("aaf PersistentPreRun failed : %v", err)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please use subCmd")
		},
	}

	cmd.PersistentFlags().StringVar(&o.InputFilePath, "input-filepath", "input.txt", "exported text file from Anki")
	cmd.PersistentFlags().StringVar(&o.OutputFilePath, "output-filepath", "output.txt", "output file name to import into Anki")
	cmd.PersistentFlags().StringVar(&o.fileFormat, "file-format", "tab-splitted", "anki format (supports: tab-splitted)")
	cmd.PersistentFlags().BoolVarP(&o.Verbose, "verbose", "v", false, "Print verbose messages")

	cmd.AddCommand(NewDefaultStarDictCommand(o))

	return cmd
}

func (o *AafOptions) Complete() error {
	var err error
	switch o.fileFormat {
	case "tab-splitted":
		o.Loader, err = ankiLoader.NewTabSplittedLoader(o.InputFilePath)
		if err != nil {
			return err
		}

		o.Saver, err = ankiSaver.NewTabSplittedSaver(o.OutputFilePath)
		if err != nil {
			return err
		}
		o.Saver, err = ankiSaver.NewTabSplittedSaver(o.OutputFilePath)
		if err != nil {
			return err
		}
	}
	return nil
}
