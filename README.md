<section align="center">

# Anki Helper
[![Go Report Card](https://goreportcard.com/badge/github.com/MarcusXavierr/anki-helper)](https://goreportcard.com/report/github.com/MarcusXavierr/anki-helper)
![example workflow](https://github.com/MarcusXavierr/anki-helper/actions/workflows/go.yml/badge.svg)

<br/>
Track new words you learn in a foreign language and add them to Anki for further practice and memorization. 

[Getting started](#getting-started) •
[Installation](#installation) •
[Configuration](#configuration) •
[Usage](#usage)

</section>

## Getting Started

### Insert flashcards on anki

https://user-images.githubusercontent.com/59923581/222997612-ba367d28-f2aa-47dd-9b8e-51c83dac1b56.mp4

### Get sentence definitions

![definition](https://user-images.githubusercontent.com/59923581/220511647-44ac85d3-a1cc-4eef-ae78-114f41dc45b8.gif)
<hr>

### Add new sentences to further practice

![add](https://user-images.githubusercontent.com/59923581/220512928-f4a311f8-256c-4af0-8e98-279b7775fb88.gif)

## Installation
you can install this project with go
```bash
go install github.com/MarcusXavierr/anki-helper@latest
```
## Configuration
To use the CLI, you must provide two files: one to save the sentences you have yet to learn and add to Anki, and another to store the sentences you have already added to Anki.

There are two ways to specify the file paths in the CLI. You can pass them as flags with every command you run, or you can configure a file called .anki-config in your home folder.

The .anki-config file allows you to set default options for the CLI, including the file paths for the sentences you have yet to learn and the sentences you have already added to Anki. By using this file, you can avoid passing the file paths as flags with every command, which can be useful if you frequently use the same file paths.

Overall, the CLI offers flexibility in how you choose to provide the file paths, making it easy for you to use it in the way that best fits your needs.

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

### Add word to further practice
You can run this command to `add` a word to your track file

```bash
anki-helper add 'sentence'
```
### Get definition
you can also use the command `dictionary` to get the definitions of a sentence (works for words and phrasal verbs)

```bash
anki-helper dictionary 'sentence'
```
