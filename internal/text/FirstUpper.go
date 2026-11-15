package text

import "strings"

func CapitalizeWords(content string) string {
	words := strings.Fields(content)

	var upperWords []string

	for i := 0; i < len(words); i++ {
		letters := strings.Split(words[i], "")
		letters[0] = ToUpperCase(letters[0])
		upperWords = append(upperWords, strings.Join(letters, ""))
	}
	return strings.Join(upperWords, " ")
}
