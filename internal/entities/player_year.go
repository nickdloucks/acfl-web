package entities

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
