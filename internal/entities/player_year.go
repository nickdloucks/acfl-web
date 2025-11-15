package entities

import (
	"errors"
	"strings"
)

type PlayerYear struct {
	Id       UuidV7Str      `json:"id"`        // primary key
	PlayerId UuidV7Str      `json:"player_id"` // foreign key to Player
	Year     uint16         `json:"year"`
	TeamId   UuidV7Str      `json:"team"` // foreign key to TeamYear
	Position PlayerPosition `json:"position"`
}

type PlayerPosition string

const (
	// --- OFFENSIVE POSITIONS:
	QB PlayerPosition = "QB"
	RB PlayerPosition = "RB"
	FB PlayerPosition = "FB"
	WR PlayerPosition = "WR"
	TE PlayerPosition = "TE"
	LT PlayerPosition = "LT"
	LG PlayerPosition = "LG"
	C  PlayerPosition = "C"
	RG PlayerPosition = "RG"
	RT PlayerPosition = "RT"
	// --- DEFENSIVE POSITIONS:
	DE  PlayerPosition = "DE"
	DT  PlayerPosition = "DT"
	OLB PlayerPosition = "OLB"
	ILB PlayerPosition = "ILB"
	RCB PlayerPosition = "RCB"
	LCB PlayerPosition = "LCB"
	FS  PlayerPosition = "FS"
	SS  PlayerPosition = "SS"
	NCB PlayerPosition = "NCB"
	DCB PlayerPosition = "DCB"
	// --- SPECIAL TEAMS POSITIONS:
	K PlayerPosition = "K"
	P PlayerPosition = "P"
)

var StringToPosition = map[string]PlayerPosition{
	// --- OFFENSIVE POSITIONS:
	"QB": QB,
	"RB": RB,
	"FB": FB,
	"WR": WR,
	"TE": TE,
	"LT": LT,
	"LG": LG,
	"C":  C,
	"RG": RG,
	"RT": RT,
	// --- DEFENSIVE POSITIONS:
	"DE":  DE,
	"DT":  DT,
	"OLB": OLB,
	"ILB": ILB,
	"RCB": RCB,
	"LCB": LCB,
	"FS":  FS,
	"SS":  SS,
	"NCB": NCB,
	"DCB": DCB,
	// --- SPECIAL TEAMS POSITIONS:
	"K": K,
	"P": P,
}

func validatePlayerPosition(position string) error {
	position = strings.ToUpper(position)
	if _, ok := StringToPosition[position]; ok {
		return nil
	}
	return errors.New("invalid player position")
}

type PlayerYearProfile struct { // this represents a single line record on a Roster list
	FIRST      string `json:"FIRST"`
	LAST       string `json:"LAST"`
	NUMBER     uint8  `json:"NUMBER"`
	Heigh      uint8  `json:"Heigh"`
	Weigh      uint16 `json:"Weigh"`
	POS        uint8  `json:"POS"`
	AWARE      uint8  `json:"AWARE"`
	SPEED      uint8  `json:"SPEED"`
	TLK_BRK    uint8  `json:"TLK BRK"`
	FUMBL      uint8  `json:"FUMBL"`
	CATCH      uint8  `json:"CATCH"`
	RTRN       uint8  `json:"RTRN"`
	PAS_BLK    uint8  `json:"PAS BLK"`
	RUN_BLK    uint8  `json:"RUN BLK"`
	THR_PWR    uint8  `json:"THR PWR"`
	THR_ACC    uint8  `json:"THR ACC"`
	KCK_ACC    uint8  `json:"KCK ACC"`
	KCK_PWR    uint8  `json:"KCK PWR"`
	BLK_SHD    uint8  `json:"BLK SHD"`
	PAS_RSH    uint8  `json:"PAS RSH"`
	TACKL      uint8  `json:"TACKL"`
	HIT_PWR    uint8  `json:"HIT PWR"`
	M_COV      uint8  `json:"M COV"`
	Z_COV      uint8  `json:"Z COV"`
	FITNESS    uint8  `json:"FITNESS"`
	DISCIPLINE uint8  `json:"DISCIPLINE"`
	POTENTIAL  uint8  `json:"POTENTIAL"`
	AGE        uint8  `json:"AGE"`
}
