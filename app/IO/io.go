package IO

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/MarcusXavierr/anki-helper/app/check"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

type File struct {
	FilePath string
}

type IFile interface {
	ReadFile() string
}

func (f File) ReadFile() string {
	buffer, err := os.ReadFile(f.FilePath)
	check.Check(err)
	return string(buffer)
}

func WriteFile(sentence, filepath string) error {
	file, e := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(644))
	check.Check(e)
	_, err := file.WriteString(sentence + "\n")
	check.Check(err)
	return file.Close()
}

func GetHomeDir() string {
	usr, err := user.Current()
	check.Check(err)
	return usr.HomeDir
}

func PrintRed(out io.Writer, message string) {
	fmt.Fprintln(out, string(colorRed), message, string(colorReset))
}

func PrintGreen(out io.Writer, message string) {
	fmt.Fprintln(out, string(colorGreen), message, string(colorReset))
}
