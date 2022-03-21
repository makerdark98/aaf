package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/dyatlov/gostardict/stardict"
	"github.com/spf13/cobra"
)

type StarDictOptions struct {
	filePath string
}

func NewDefaultStarDictCommand(aafOptions AafOptions) *cobra.Command {
	o := StarDictOptions{}

	cmd := &cobra.Command{
		Use:   "stardict",
		Short: "Fill data using stardict file",
		Long: `
Generate data using stardict file.
When the filepath option is "data/KoreanDic", you should set like
.
+-- data
    |-- KoreanDic.dict.dz (or KoreanDic.dict)
    |-- KoreanDic.idx
    +-- KoreanDic.ifo

Some dictionaries may not have a .dict or .idx file, but there should be an ifo file.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			err := o.Complete()
			if err != nil {
			}

			err = o.Run()
			if err != nil {

			}
		},
	}

	cmd.Flags().StringVar(&o.filePath, "filepath", "", "stardict filename to translate data")
	return cmd
}

func (o *StarDictOptions) Complete() error {
	if len(o.filePath) == 0 {
		return fmt.Errorf("stardict command error: empty filepath") // TODO: make new error type
	}
	return nil
}

func (o *StarDictOptions) Run() error {
	path, name := filepath.Split(o.filePath)
	dict, err := stardict.NewDictionary(path, name)
	if err != nil {
		return err
	}

	/*
		TODO:
		1. Read exported file from Anki
		2. Loop items and translate items
		3. Save file
	*/
	_ = dict.Translate("closely") // Delete me: it's just dummy call.

	return nil
}
