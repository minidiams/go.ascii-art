package module

import (
	"io/ioutil"
	"os"
)

type Alphabet struct {
	LetterAscii map[string][8]string
}

func getLetter(text []byte, num int) (result [8]string) {
	line := -1
	i := -1
	for j := 0; j < len(text); j++ {
		if text[j] == '\n' {
			line++
		}
		if line/9 == num && line != -1 {
			if string(text[j]) != "\n" && i <= 7 {
				result[i] += string(text[j])
			}
			if text[j] == '\n' {
				i++
			}
		}
	}
	return
}

func GetAlphabet(file []byte) *Alphabet {
	var n = make(map[string][8]string)
	result := &Alphabet{LetterAscii: n}
	for i := ' ' - 32; i < '~'-32; i++ {
		result.LetterAscii[string(i+32)] = getLetter(file, int(i))
	}
	return result
}

func GetAlphabetFile() (result []byte) {
	file := "standard.txt"
	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "standard":
			file = "standard.txt"
		case "shadow":
			file = "shadow.txt"
		case "thinkertoy":
			file = "thinkertoy.txt"
		}
	}
	result, _ = ioutil.ReadFile("file/" + file)
	return
}