package logger

import "github.com/common-nighthawk/go-figure"

// PrintFigure print the string in elegant way
func PrintFigure(appName string) {
	var fontName = "stop"
	if len(appName) >= 16 {
		fontName = "thin"
	}
	if len(appName) >= 20 {
		fontName = "straight"
	}
	if len(appName) >= 28 {
		fontName = "pepper"
	}
	if len(appName) >= 36 {
		fontName = "short"
	}
	if len(appName) >= 48 {
		fontName = "eftipiti"
	}
	if len(appName) >= 72 {
		fontName = "term"
	}
	figure.NewColorFigure(appName, fontName, "blue", true).Print()
}

// PrintErrorMsgFigure print error message elegantly
func PrintErrorMsgFigure(msg string) {
	figure.NewColorFigure(msg, "term", "red", true).Blink(60000, 1000, -1)
}
