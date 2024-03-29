package IO

import (
	"io/fs"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
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

type UserFilePath struct {
	WriteFile        string
	ManualInsertFile string
	TrashFile        string
}

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

func MoveSentenceToFile(trashPath, wordsPath, sentence string) error {
	err := DeleteSentenceFromFile(wordsPath, sentence)

	if err != nil {
		return err
	}

	return WriteFile(sentence, trashPath)
}

func DeleteSentenceFromFile(filePath, searchString string) error {
	dir, path := filepath.Split(filePath)

	newContent, err := FilterSentenceFromFile(os.DirFS(dir), path, searchString)

	if err != nil {
		return err
	}

	return OverrideFile(newContent, filePath)
}

func FilterSentenceFromFile(fsys fs.FS, filePath string, searchString string) ([]byte, error) {
	content, err := fs.ReadFile(fsys, filePath)
	if err != nil {
		return []byte{}, err
	}

	lines := strings.Split(string(content), "\n")
	var found bool
	var newContent []string
	for _, line := range lines {
		if strings.Contains(line, searchString) && !found {
			found = true
			continue
		}
		newContent = append(newContent, line)
	}

	newContentStr := strings.Join(newContent, "\n")
	return []byte(newContentStr + "\n"), nil
}

func OverrideFile(data []byte, filepath string) error {
	err := ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func removeBlankLineFromEnd(lines []string) []string {
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}
