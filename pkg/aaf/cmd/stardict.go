package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/dyatlov/gostardict/stardict"
	"github.com/spf13/cobra"
)

type StarDictOptions struct {
	filePath   string
	aafOptions *AafOptions
}

func NewDefaultStarDictCommand(aafOptions *AafOptions) *cobra.Command {
	o := StarDictOptions{
		aafOptions: aafOptions,
	}

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
				fmt.Printf("stardict command initilization failed : %v\n", err)
				return
			}

			err = o.Run()
			if err != nil {
				fmt.Printf("stardict command failed : %v\n", err)
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

	deck, err := o.aafOptions.Loader.Load()
	if err != nil {
		return err
	}

	for i := range deck.Cards {
		if len(deck.Cards[i].Items) == 0 {
			continue
		}
		fmt.Println(deck.Cards[i].Items[0])

		translatedItems := dict.Translate(deck.Cards[i].Items[0])
		if len(translatedItems) == 0 || len(translatedItems[0].Parts) == 0 {
			continue
		}

		deck.Cards[i].Items = append(deck.Cards[i].Items, string(translatedItems[0].Parts[0].Data))
	}

	err = o.aafOptions.Saver.Save(deck)
	if err != nil {
		return err
	}

	return nil
}
