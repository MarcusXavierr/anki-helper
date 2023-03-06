<section align="center">

# Anki Helper
[![Go Report Card](https://goreportcard.com/badge/github.com/MarcusXavierr/anki-helper)](https://goreportcard.com/report/github.com/MarcusXavierr/anki-helper)
![example workflow](https://github.com/MarcusXavierr/anki-helper/actions/workflows/go.yml/badge.svg)

<br/>
Track new words you learn in a foreign language and add them to Anki for further practice and memorization. 

[Usage Examples](#usage-examples) •
[Installation](#installation) •
[Getting Startd](#getting-started) •
[Usage](#usage)

</section>

## Usage Examples

### Insert flashcards on anki

https://user-images.githubusercontent.com/59923581/222997612-ba367d28-f2aa-47dd-9b8e-51c83dac1b56.mp4

### Get sentence definitions

![definition](https://user-images.githubusercontent.com/59923581/220511647-44ac85d3-a1cc-4eef-ae78-114f41dc45b8.gif)
<hr>

### Add new sentences to further practice

![add](https://user-images.githubusercontent.com/59923581/220512928-f4a311f8-256c-4af0-8e98-279b7775fb88.gif)

## Installation

### go install
you can install this project with go
```bash
go install github.com/MarcusXavierr/anki-helper@latest
```
### Binaries
You can also install a compiled binary to your machine and then put it in your shell path.

Go to the [releases page](https://github.com/MarcusXavierr/anki-helper/releases) and choose the option that best fits your environment.

## Getting Startd
To get started, visit the [project's wiki](https://github.com/MarcusXavierr/anki-helper/wiki). There you will find instructions on how to set up your Anki and how to configure Anki-helper, in a simple and fast way.

## Usage

### Insert flashcards on anki
You can easilly insert flashcards on anki using anki-helper.
```bash
anki-helper ankiInsert

# Or if you want to specify how many flashcards you wanna insert on anki, use the -i flag

anki-helper ankiInsert -i 10
```

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
