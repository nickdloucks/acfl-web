package entities

import (
	"errors"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

type UuidV7Str string // A uuid v7 string.
type UuidV7ProviderInterface interface {
	NewUuidV7() UuidV7Str
}

type GameStatus string // Enumerated categorization of a game's current state.

const (
	TBA        GameStatus = "TBA" // not yet scheduled
	Upcoming   GameStatus = "UPCOMING"
	InProgress GameStatus = "IN-PROGRESS"
	Final      GameStatus = "FINAL"
)

var IntToGameStatus = map[int]GameStatus{
	1: TBA,
	2: Upcoming,
	3: InProgress,
	4: Final,
}
var StrToGameStatus = map[string]GameStatus{
	"TBA":         TBA,
	"UPCOMING":    Upcoming,
	"IN-PROGRESS": InProgress,
	"FINAL":       Final,
}

// --- Color Hex String ---
type ColorStr string // hex color code with a leading "#" character i.e. "#0f0f0f"

const (
	BLACK     ColorStr = "#000000" // rgb(0 0 0)
	WHITE     ColorStr = "#ffffff" // 255 255 255
	DARKGREY  ColorStr = "#1f1f1f" // 31 31 31
	MEDGREY   ColorStr = "#363636" // 54 54 54
	GREY      ColorStr = "#818181" // 129 129 129
	LIGHTGREY ColorStr = "#999999" // 153 153 153
	// EXTRALIGHTGREY ColorStr = "#cccccc" //
	BROWN      ColorStr = "#8b3613" // 139 54 19
	RED        ColorStr = "#ff0000" // 255 0 0
	ORANGE     ColorStr = "#ff6100" // 255 97 0
	YELLOW     ColorStr = "#ffff00" // 255 255 0
	GREEN      ColorStr = "#008800" // 0 136 0
	OLIVE      ColorStr = "#556b2f" // 85 107 47
	LIGHTGREEN ColorStr = "#32cd32" // 50 205 50
	LIME       ColorStr = "#00ff00" // 0 255 0
	BLUE       ColorStr = "#0000ff" // 0 0 255
	LIGHTBLUE  ColorStr = "#00bfff" // 0 191 255
	// INDIGO         ColorStr = "#4b0082" //
	NAVY        ColorStr = "#00008b" // 0 0 139
	VIOLET      ColorStr = "#7f00ff" // 127 0 255
	PURPLE      ColorStr = "#4b0082" // 75 0 130
	GOLD        ColorStr = "#d4af37" // 212 175 55
	SILVER1     ColorStr = "#bcc6cc" // 188 198 204
	SILVER2     ColorStr = "#999b9b" // 153 155 155
	GRAPHITE    ColorStr = "#a8afb1" // 168 175 177
	JET         ColorStr = "#373635" // 55 54 53
	SILVER3     ColorStr = "#eeeeee" // 238 238 238
	SILVER4     ColorStr = "#cccccc" // 204 204 204
	SILVER5     ColorStr = "#bbbbbb" // 187 187 187
	SILVER6     ColorStr = "#aaaaaa" // 170 170 170
	SILVER7     ColorStr = "#777777" // 119 119 119
	SILVER8     ColorStr = "#c0c0c0" // 192 192 192
	COPPER      ColorStr = "#b87333" // 184 115 51
	LIGHTpURPLE ColorStr = "#8a2be2" // 138 43 226
)

var ColorPresets = map[string]ColorStr{
	"#000000": BLACK,
	"#ffffff": WHITE,
	"#1f1f1f": DARKGREY,
	"#363636": MEDGREY,
	"#818181": GREY,
	"#999999": LIGHTGREY,
	// "#cccccc": EXTRALIGHTGREY,
	"#8b3613": BROWN,
	"#ff0000": RED,
	"#ff6100": ORANGE,
	"#ffff00": YELLOW,
	"#008800": GREEN,
	"#556b2f": OLIVE,
	"#00ff00": LIME,
	"#32cd32": LIGHTGREEN,
	"#0000ff": BLUE,
	"#00bfff": LIGHTBLUE,
	// "#4b0082": INDIGO,
	"#00008b": NAVY,
	"#7f00ff": VIOLET,
	"#4b0082": PURPLE,
	"#d4af37": GOLD,
	"#bcc6cc": SILVER1,
	"#999b9b": SILVER2,
	"#a8afb1": GRAPHITE,
	"#373635": JET,
	"#eeeeee": SILVER3,
	"#cccccc": SILVER4,
	"#bbbbbb": SILVER5,
	"#aaaaaa": SILVER6,
	"#777777": SILVER7,
	"#c0c0c0": SILVER8,
	"#b87333": COPPER,
}

var hexRunes = map[rune]bool{
	'a': true,
	'b': true,
	'c': true,
	'd': true,
	'e': true,
	'f': true,
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
}

// Validates a color string.
// Expects a 6-character hex string prefixed with "#".
// Returns a default (#363636) grey color hex string if input is invald.
func validateColor(color string) ColorStr {
	defaultColor := "#363636"
	color = strings.ToLower(color)
	if validated, ok := ColorPresets[color]; ok {
		log.Println("color: " + ColorPresets[color])
		return validated
	} else {
		color = strings.Trim(color, "#")
		if len(color) != 6 {
			return ColorStr(defaultColor)
		}
		for _, char := range color {
			if _, ok := hexRunes[char]; !ok {
				return ColorStr(defaultColor)
			}
		}
		return ColorStr("#" + color)
	}
}

func validateGameStatus(status string) error {
	status = strings.ToUpper(status)
	if _, ok := StrToGameStatus[status]; ok {
		return nil
	}
	return errors.New("invalid game status string")
}

func capitalizeWords(input string) (string, error) {
	if input == "" {
		return "", errors.New("invalid input: empty")
	}
	parts := strings.Split(input, " ")
	capitalizedParts := []string{}
	for _, item := range parts {
		r, size := utf8.DecodeRuneInString(item)
		if r == utf8.RuneError {
			return "", errors.New("invalid input")
		}
		item = string(unicode.ToUpper(r)) + item[size:]
		capitalizedParts = append(capitalizedParts, item)
	}
	input = strings.Join(capitalizedParts, " ")
	return input, nil
}