package serr

import (
	"fmt"
	"strings"
)

// elementType type
type elementType int

const (
	// elementTypeForeground constant for foreground color type
	elementTypeForeground elementType = iota

	// elementTypeBackground constant for background color type
	elementTypeBackground
)

// consoleColor type
type consoleColor string

const (
	// consoleColorDefault constant for default color
	consoleColorDefault consoleColor = "39:49"

	// consoleColorBlack constant for black color
	consoleColorBlack consoleColor = "30:40"

	// consoleColorRed constant for red color
	consoleColorRed consoleColor = "31:41"

	// consoleColorGreen constant for green color
	consoleColorGreen consoleColor = "32:42"

	// consoleColorYellow constant for yellow color
	consoleColorYellow consoleColor = "33:43"

	// consoleColorBlue constant for blue color
	consoleColorBlue consoleColor = "34:44"

	// consoleColorMagenta constant for magenta color
	consoleColorMagenta consoleColor = "35:45"

	// consoleColorCyan constant for cyan color
	consoleColorCyan consoleColor = "36:46"

	// consoleColorLightGray constant for light gray color
	consoleColorLightGray consoleColor = "37:47"

	// consoleColorDarkGray constant for dark gray color
	consoleColorDarkGray consoleColor = "90:100"

	// consoleColorLightRed constant for light red color
	consoleColorLightRed consoleColor = "91:101"

	// consoleColorLightGreen constant for light green color
	consoleColorLightGreen consoleColor = "92:102"

	// consoleColorLightYellow constant for light yellow color
	consoleColorLightYellow consoleColor = "93:103"

	// consoleColorLightBlue constant for light blue color
	consoleColorLightBlue consoleColor = "94:104"

	// consoleColorLightMagenta constant for light magenta color
	consoleColorLightMagenta consoleColor = "95:105"

	// consoleColorLightCyan constant for light cyan color
	consoleColorLightCyan consoleColor = "96:106"

	// consoleColorWhite constant for white color
	consoleColorWhite consoleColor = "97:107"
)

const (
	// escChar constant
	escChar = "\x1B"

	// resetChar constant
	resetChar = escChar + "[0m"
)

// extractConsoleColorCode function
func extractConsoleColorCode(color consoleColor, typ elementType) (string, bool) {
	splittedColor := strings.Split(string(color), ":")
	if len(splittedColor) == 0 {
		splittedColor = append(splittedColor, splittedColor[0])
	}

	return splittedColor[int(typ)], true
}

// applyConsoleForeColor function
func applyConsoleForeColor(msg string, color consoleColor) string {
	if color, ok := extractConsoleColorCode(color, elementTypeForeground); ok {
		return fmt.Sprintf("%s[%sm%s%s", escChar, color, msg, resetChar)
	}
	return msg
}

// applyConsoleBackgroundColor function
func applyConsoleBackgroundColor(msg string, color consoleColor) string {
	if color, ok := extractConsoleColorCode(color, elementTypeBackground); ok {
		return fmt.Sprintf("%s[%sm%s%s", escChar, color, msg, resetChar)
	}
	return msg
}
