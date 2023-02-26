package IO

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/user"
	"strings"
)

const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[31m"
	ColorGreen = "\033[32m"
	ColorCyan  = "\033[36m"
	ColorPink  = "\033[38;5;205m"
	ColorGold  = "\033[38;2;243;134;48m"
)

type File struct {
	FilePath string
}

type IFile interface {
	ReadFile() string
}

func (f File) ReadFile() string {
	buffer, err := os.ReadFile(f.FilePath)
	check(err)
	return string(buffer)
}

func WriteFile(sentence, filepath string) error {
	file, e := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(644))
	check(e)
	_, err := file.WriteString(sentence + "\n")
	check(err)
	return file.Close()
}

func GetHomeDir() string {
	usr, err := user.Current()
	check(err)
	return usr.HomeDir
}

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

func GetWords(fsys fs.FS, filename string, numberOfLines int) ([]string, error) {
	file, err := fs.ReadFile(fsys, filename)

	if err != nil {
		return []string{}, err
	}

	lines := strings.Split(string(file), "\n")

	lines = removeBlankLineFromEnd(lines)

	if numberOfLines > len(lines) {
		numberOfLines = len(lines)
	}

	return lines[:numberOfLines], nil

}

func removeBlankLineFromEnd(lines []string) []string {
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}
