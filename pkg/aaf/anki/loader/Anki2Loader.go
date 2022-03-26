package loader

import (
	"bytes"
	"encoding/json"
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

	cards, err := i.getCard()
	if err != nil {
		return nil, err
	}
	fmt.Println(cards)

	models, err := i.getCol()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Models : %v\n", models)

	return &anki.Deck{}, nil
}

func (i *Anki2Importer) getCard() ([]Card, error) {
	models, err := i.getModels()
	if err != nil {
		return nil, err
	}

	notedCards := make([]notedCard, 0)

	queryResult := i.db.Table("cards").
		Select("cards.Id, nid, did, ord, cards.mod, cards.usn, type, queue, due, ivl, factor, reps, lapses, left, odue, odid, cards.flags, cards.data, notes.id, guid, mid, notes.mod, notes.usn, tags, flds, sfld, csum, notes.flags, notes.data").
		Joins("LEFT JOIN notes ON cards.nid = notes.id").
		Find(&notedCards)
	if queryResult.Error != nil {
		return nil, queryResult.Error
	}
	fmt.Println(notedCards[0].NFlds)

	cards := make([]Card, len(notedCards))
	for i := range notedCards {
		cards[i].Card = &card{
			notedCards[i].Id,
			notedCards[i].Nid,
			notedCards[i].Did,
			notedCards[i].Ord,
			notedCards[i].Mod,
			notedCards[i].Usn,
			notedCards[i].Type,
			notedCards[i].Queue,
			notedCards[i].Due,
			notedCards[i].Ivl,
			notedCards[i].Factor,
			notedCards[i].Reps,
			notedCards[i].Lapses,
			notedCards[i].Left,
			notedCards[i].Odue,
			notedCards[i].Odid,
			notedCards[i].Flags,
			notedCards[i].Data,
		}
		cards[i].Note = &note{
			notedCards[i].NId,
			notedCards[i].NGuid,
			notedCards[i].NMid,
			notedCards[i].NMod,
			notedCards[i].NUsn,
			notedCards[i].NTags,
			notedCards[i].NFlds,
			notedCards[i].NSfld,
			notedCards[i].NCsum,
			notedCards[i].NFlags,
			notedCards[i].NData,
		}
		model := models[notedCards[i].NMid]
		cards[i].Model = &model
	}

	return nil, fmt.Errorf("not implemented")
}

func (i *Anki2Importer) getCol() ([]col, error) {
	cols := make([]col, 0)
	queryResult := i.db.Table("col").
		Select("*").
		Find(&cols)

	if queryResult.Error != nil {
		return nil, queryResult.Error
	}
	return cols, nil
}

func (i *Anki2Importer) getModels() (map[string]model, error) {
	modelJsons := make([]string, 0)
	queryResult := i.db.Table("col").
		Select("models").
		Find(&modelJsons)
	if queryResult.Error != nil {
		return nil, queryResult.Error
	}

	m := make(map[string]model, 0)
	for i := range modelJsons {
		parsedModels := make(map[string]model, 0)
		err := json.Unmarshal([]byte(modelJsons[i]), &parsedModels)
		if err != nil {
			return nil, err
		}

		for k := range parsedModels {
			m[k] = parsedModels[k]
		}
	}

	return m, nil
}

func splitFields(fields string) []string {
	return strings.Split(fields, "\u001f")
}

type Card struct {
	Card  *card
	Note  *note
	Model *model
}

func (c Card) String() string {
	if c.Card == nil {
		return "empty card"
	}

	if c.Model != nil && c.Note != nil {
		var b bytes.Buffer
		fields := splitFields(c.Note.Flds)
		for i := range fields {
			b.WriteString(fields[i] + "(" + c.Model.Fields[i].Name + "),")
		}
		return b.String()
	}

	return "unreadable card"
}

type card struct {
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
	Mid   string // related col
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
	Usn  int
	Oid  int
	Type int
}

type revlog struct {
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

type deck struct {
}

type template struct {
	Name  string  `json:"name"`
	Ord   int     `json:"ord"`
	QFmt  string  `json:"qfmt"`
	AFmt  string  `json:"afmt"`
	BQFmt string  `json:"bqfmt"`
	BAFmt string  `json:"bafmt"`
	Did   *string `json:"did"`
	BFont string  `json:"bfont"`
	BSize int     `json:"bsize"`
}

type field struct {
	Name   string `json:"name"`
	Ord    int    `json:"ord"`
	Sticky bool   `json:"sticky"`
	RTL    bool   `json:"RTL"`
	Font   string `json:"font"`
	Size   int    `json:"size"`
}
type model struct {
	Id        int           `json:"id"`
	Name      string        `json:"name"`
	Type      int           `json:"type"`
	Mod       int           `json:"mod"`
	Usn       int           `json:"usn"`
	Sortf     int           `json:"sortf"`
	Did       *int          `json:"did"`
	Templates []template    `json:"tmpls"`
	Fields    []field       `json:"flds"`
	CSS       string        `json:"css"`
	LatexPre  string        `json:"latexPre"`
	LatexPost string        `json:"latexPost"`
	LatexSvg  bool          `json:"latexsvg"`
	Req       []interface{} `json:"req"`
}

type notedCard struct {
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
	NId    int
	NGuid  string
	NMid   string // related col
	NMod   int
	NUsn   int
	NTags  string
	NFlds  *string
	NSfld  string
	NCsum  int
	NFlags int
	NData  string
}
