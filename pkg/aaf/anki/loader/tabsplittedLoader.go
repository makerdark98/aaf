package loader

import (
	"fmt"

	"github.com/makerdark98/aaf/pkg/aaf/anki"
)

type TabSplittedLoader struct {
	filepath string
}

func NewTabSplittedLoader(filepath string) (*TabSplittedLoader, error) {
	return &TabSplittedLoader{
		filepath: filepath,
	}, nil
}

func (l *TabSplittedLoader) Load() (*anki.Deck, error) {
	return nil, fmt.Errorf("Not implemented")
}
