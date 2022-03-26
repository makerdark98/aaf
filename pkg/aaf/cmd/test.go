package cmd

import (
	"fmt"

	"github.com/makerdark98/aaf/pkg/aaf/anki/loader"
	"github.com/spf13/cobra"
)

func NewDefaultTestCommand(aafOptions *AafOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "test playground",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run test command")
			loader, err := loader.NewAnki2Importer("data/collection.anki21")
			if err != nil {
				fmt.Errorf("loader error : %v", err)
				return
			}

			_, err = loader.Load()
			if err != nil {
				fmt.Errorf("loader error : %v", err)
				return
			}
		},
	}

	return cmd
}
