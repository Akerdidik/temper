package exam

import (
	"fmt"
)

func Brainfuck(s string) {
	progpoint := []byte(s)
	var arby [2048]byte
	pos := 0
	openBr := 0         // opened brackets
	i := 0              // iterates through the source code passed in the argument
	N := len(progpoint) // length of the source code
	for i >= 0 && i < N {
		switch progpoint[i] {
		case '>':
			pos++
		case '<':
			pos--
		case '+':
			arby[pos]++
		case '-':
			arby[pos]--
		case '.':
			fmt.Printf("%c", rune(arby[pos]))
		case '[':
			openBr = 0
			if arby[pos] == 0 {
				for i < N && (progpoint[i] != byte(']') || openBr > 1) {
					if progpoint[i] == byte('[') {
						openBr++
					} else if progpoint[i] == byte(']') {
						openBr--
					}
					i++
				}
			}
		case ']':
			openBr = 0
			if arby[pos] != 0 {
				for i >= 0 && (progpoint[i] != byte('[') || openBr > 1) {
					if progpoint[i] == byte(']') {
						openBr++
					} else if progpoint[i] == byte('[') {
						openBr--
					}
					i--
				}
			}
		}
		i++
	}
}
