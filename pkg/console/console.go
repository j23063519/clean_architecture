package console

import (
	"os"

	"github.com/fatih/color"
)

func Success(msg string) {
	newColor(color.FgGreen, color.Bold, color.BgWhite).
		Printf("\nSuccess: %+v\n\n", msg)
}

func Error(msg string) {
	newColor(color.FgRed, color.Bold, color.BgWhite).
		Printf("\nError: %+v\n\n", msg)
}

func Warn(msg string) {
	newColor(color.FgMagenta, color.Bold, color.BgWhite).
		Printf("\nWarn: %+v\n\n", msg)
}

func newColor(colors ...color.Attribute) *color.Color {
	col := color.New(colors[0])
	var colorSlice []*color.Color
	colorSlice = append(colorSlice, col)
	for i, c := range colors {
		if i != 0 {
			colorSlice = append(colorSlice, colorSlice[i-1].Add(c))
		}
	}

	return colorSlice[len(colorSlice)-1]
}

func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}
