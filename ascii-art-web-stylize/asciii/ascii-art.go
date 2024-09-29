package assci

import (
	"strings"
)

func AscciArt(input string, banner string) string {
	contant := strings.ReplaceAll(banner, "\r", "")
	lines := strings.Split(contant, "\n")
	var table [][]string
	cou := 0
	for i := 1; i < len(lines); i++ {
		if cou < len(lines) {
			table = append(table, []string{})
		}
		for l := 0; l < 8; l++ {
			table[cou] = append(table[cou], lines[i])
			i++
		}
		cou++
	}
	var result string
	var runes []rune
	var NewInput []string

	inputsplit := strings.Split(input, "\r\n")
	for _, str := range inputsplit {
		for i := 0; i <= len(str)-1; i++ {
			if str[i] > 126 || str[i] < 32 {
				continue
			} else {
				runes = append(runes, rune(str[i]))
			}
		}
		NewInput = append(NewInput, string(runes))
		runes = []rune{}
	}

	for _, str := range NewInput {
		if len(str) == 0 {
			result += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range str {
				result += (table[int(char)-32][i])
			}
			result += "\n"
		}
	}
	return result
}
