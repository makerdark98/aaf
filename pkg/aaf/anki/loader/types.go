package loader

import "github.com/makerdark98/aaf/pkg/aaf/anki"

type Interface interface {
	Load() (*anki.Deck, error)
}
