# Anki Helper
[![Go Report Card](https://goreportcard.com/badge/github.com/MarcusXavierr/anki-helper)](https://goreportcard.com/report/github.com/MarcusXavierr/anki-helper)
![example workflow](https://github.com/MarcusXavierr/anki-helper/actions/workflows/go.yml/badge.svg)

<br/>
This application aims to allow the user to track the words he learns in a language, so that they can put these words in anki

<img width="1049" alt="image" src="https://user-images.githubusercontent.com/59923581/219258940-86583cb9-6cc4-4cf0-af9e-b8e98472a540.png">

## Installation
you can install this project with go
```bash
go install github.com/MarcusXavierr/anki-helper@latest
```
## Configuration
You need to provide two files to the cli, one to save the sentences you've yet to learn and put in anki, and another to store the sentences you've already put in anki.
There are two ways to pass the path of these files to the cli. You can pass these as flags to every command you make. Or you can configure a file called .anki-config that will be in your home folder.

### Configuration storage file paths
Here, I'll show the two ways to set your anki helper storage file paths. Remember to use your own paths.

#### Using flags
Setting your storage file paths with flags is pretty easy. Just use `-n` in the file where you'll store unknown sentences, and `-t` to store learned sentences.

```bash
anki-helper add "unknown sentence" -n "/tmp/new_sentences_file.txt" -t "/tmp/learned_sentences_file.txt"
```
#### Using a configuration file
Create a file called `.anki-config.toml` in your home folder. Then put this content there
```toml
new-words-file-path="/tmp/new_sentences_file.txt"
trash-file-path="/tmp/learned_sentences_file.txt"
```

## Usage
You can run this command to add a word to your track file

```bash
anki-helper add word
```
you can also use -d flag to get definition of this word

```bash
anki-helper add word -d

#or

anki-helper add -d word
```
