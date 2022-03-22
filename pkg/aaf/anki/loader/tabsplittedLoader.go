package loader

import (
	"bufio"
	"os"
	"strings"

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
	f, err := os.Open(l.filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	deck := &anki.Deck{}

	// TODO: Check total line and pre-allocate space
	for scanner.Scan() {
		line := scanner.Text()
		card := anki.Card{
			Items: strings.Split(line, "\t"),
		}
		deck.Cards = append(deck.Cards, card)
	}

	return deck, nil
}
