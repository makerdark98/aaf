package saver

import (
	"fmt"

	"github.com/makerdark98/aaf/pkg/aaf/anki"
)

type TabSplittedSaver struct {
	filepath string
}

func NewTabSplittedSaver(filepath string) (*TabSplittedSaver, error) {
	return &TabSplittedSaver{
		filepath: filepath,
	}, nil
}

func (l *TabSplittedSaver) Save(deck *anki.Deck) error {
	return fmt.Errorf("Not implemented")
}
