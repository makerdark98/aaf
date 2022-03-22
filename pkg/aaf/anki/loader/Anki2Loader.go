package loader

import (
	"fmt"
	"strings"

	"github.com/makerdark98/aaf/pkg/aaf/anki"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

const (
	GUID = 1
	MID  = 2
	MOD  = 3
)

type Anki2Importer struct {
	needMapper  bool
	deckPrefix  string
	allowUpdate bool
	filepath    string
	db          *gorm.DB
}

func NewAnki2Importer(filepath string) (*Anki2Importer, error) {
	return &Anki2Importer{
		filepath: filepath,
		db:       nil,
	}, nil
}

type cardKey struct {
	guid, ord int
}

func (i *Anki2Importer) Load() (*anki.Deck, error) {
	//cards := make(map[cardKey]int, 0)

	db, err := gorm.Open(sqlite.Open(i.filepath), &gorm.Config{})
	i.db = db

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	_, err = i.getCard()

	if err != nil {
		return nil, err
	}

	return &anki.Deck{}, nil
}

func (i *Anki2Importer) getCard() ([]card, error) {
	notedCards := make([]struct {
		Guid string
		Mid  string
		card
	}, 0)

	// SELECT notes.guid, notes.mid, cards.* FROM cards LEFT JOIN notes ON cards.nid = notes.id
	queryResult := i.db.Table("cards").
		Select("notes.guid, notes.mid, cards.*").
		Joins("LEFT JOIN notes ON cards.nid = notes.id").
		Find(&notedCards)
	if queryResult.Error != nil {
		return nil, queryResult.Error
	}

	notes := make([]note, 0)

	queryResult = i.db.Table("notes").
		Select("*").
		Find(&notes)
	if queryResult.Error != nil {
		return nil, queryResult.Error
	}

	return nil, fmt.Errorf("not implemented")
}

func splitFields(fields string) []string {
	return strings.Split(fields, "\u001f")
}

type card struct {
	gorm.Model
	Id     int
	Nid    int
	Did    int
	Ord    int
	Mod    int
	Usn    int
	Type   int
	Queue  int
	Due    int
	Ivl    int
	Factor int
	Reps   int
	Lapses int
	Left   int
	Odue   int
	Odid   int
	Flags  int
	Data   string
}

type note struct {
	Id    int
	Guid  string
	Mid   int
	Mod   int
	Usn   int
	Tags  string
	Flds  string
	Sfld  string
	Csum  int
	Flags int
	Data  string
}

type grave struct {
	gorm.Model
	Usn  int
	Oid  int
	Type int
}

type revlog struct {
	gorm.Model
	Id      int
	Cid     int
	Usn     int
	Ease    int
	Ivl     int
	LastIvl int
	Factor  int
	Time    int
	Type    int
}

type col struct {
	gorm.Model
	Id     int
	Crt    int
	Mod    int
	Scm    int
	Ver    int
	Dty    int
	Usn    int
	Ls     int
	Conf   string
	Models string
	Decks  string
	Dconf  string
	Tags   string
}

type Media struct {
}
