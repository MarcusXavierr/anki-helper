package IO

import (
	"fmt"
	"io"
)

func PrintRed(out io.Writer, message string) {
	PrintWithColor(out, message, ColorRed)
}

func PrintGreen(out io.Writer, message string) {
	PrintWithColor(out, message, ColorGreen)
}

func PrintWithColor(out io.Writer, message, color string) {
	fmt.Fprint(out, color, message, string(ColorReset))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
