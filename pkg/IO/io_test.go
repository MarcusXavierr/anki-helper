package IO

import (
	"reflect"
	"testing"
	"testing/fstest"
)

type fakeFile struct {
	FilePath string
}

func (f fakeFile) ReadFile() string {
	return "this is\njust\ntesting"
}

func TestReadFile(t *testing.T) {
	f := fakeFile{FilePath: "testing"}
	got := f.ReadFile()
	want := "this is\njust\ntesting"

	validateString(want, got, t)
}

func TestGetWords(t *testing.T) {
	fs := fstest.MapFS{
		"english_words.txt": {Data: []byte("hi\nmy\nname\nis\nmarcus")},
		"trash.txt":         {Data: []byte("hello\nworld\n")},
	}

	t.Run("read 3 lines from a file with 5 lines", func(t *testing.T) {
		got, _ := GetWords(fs, "english_words.txt", 3)
		want := []string{"hi", "my", "name"}

		assertArray(got, want, t)
	})

	t.Run("read 2 lines from a file with 3 lines", func(t *testing.T) {
		got, _ := GetWords(fs, "trash.txt", 5)
		want := []string{"hello", "world"}

		assertArray(got, want, t)
	})
}

func TestDeleteSentenceFromFile(t *testing.T) {
	fs := fstest.MapFS{
		"english_words.txt": {Data: []byte("hi\nmy\nname\nis\nmarcus")},
	}

	content, err := FilterSentenceFromFile(fs, "english_words.txt", "name")

	if err != nil {
		t.Fatal(err)
	}

	got := string(content)
	want := "hi\nmy\nis\nmarcus\n"

	if got != want {
		t.Errorf("expected %s but got %s", want, got)
	}
}

func assertArray(got []string, want []string, t *testing.T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v", want, got)
	}
}

func validateString(got string, want string, t testing.TB) {
	t.Helper()

	if got != want {
		t.Errorf("Error printing, got %q want %q", got, want)
	}
}
