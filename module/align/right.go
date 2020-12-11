package align

import (
	"fmt"
	"strings"
)

func right(color string, sentence string, sizeCmd int) {
	fmt.Print(color)
	for _, l := range strings.Split(sentence, "\n") {
		if l != "" {
			fmt.Println(getSpace(sizeCmd-len(l)-1), l)
		}
	}
	fmt.Print("\033[0m")
}
