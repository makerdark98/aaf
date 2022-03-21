package saver

import "github.com/makerdark98/aaf/pkg/aaf/anki"

type Interface interface {
	Save(*anki.Deck) error
}
