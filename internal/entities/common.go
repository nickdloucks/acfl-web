package entities

import "strings"

type UuidV7Str string // A uuid v7 string.


// --- Color Hex String ---
type ColorStr string // hex color code with a leading "#" character i.e. "#0f0f0f"

const (
	BLACK          ColorStr = "#000000"
	WHITE          ColorStr = "#ffffff"
	DARKGREY       ColorStr = "#1f1f1f"
	MEDGREY        ColorStr = "#363636"
	GREY           ColorStr = "#818181"
	LIGHTGREY      ColorStr = "#999999"
	EXTRALIGHTGREY ColorStr = "#cccccc"
	BROWN          ColorStr = "#8b3613"
	RED            ColorStr = "#ff0000"
	ORANGE         ColorStr = "#ff6100"
	YELLOW         ColorStr = "#ffff00"
	GREEN          ColorStr = "#008800"
	OLIVE          ColorStr = "#556b2f"
	BLUE           ColorStr = "#0000ff"
	LIGHTBLUE      ColorStr = "#00bfff"
	INDIGO         ColorStr = "#"
	NAVY           ColorStr = "#"
	VIOLET         ColorStr = "#"
	PURPLE         ColorStr = "#4b0082"
	// LIME      ColorStr = "#"
)

var ColorPresets = map[string]ColorStr{}

// Validates the hex string is a valid color and returns a default grey color hex string if it is not.
func validateColor(color string) ColorStr {
	color = strings.ToLower(color)
	if _, ok := ColorPresets[color]; ok {
		return ColorStr(color)
	}
	if !strings.HasPrefix(color, "#") {
		color = "#" + color
	}
	if len(color) <= 4 {
		// if color hex is only 3 characters, duplicate it so that it is 6
		color = strings.Trim(color, "#")
		color = "#" + color + color
	}
	return ColorStr(color)
}