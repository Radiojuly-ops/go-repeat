package text

import "strings"

func ReverseWords(content string) string {
	words := strings.Fields(content)

	var revwords []string

	for i := 0; i < len(words); i++ {
		letters := strings.Split(words[i], "")

		var revbukva []string

		for r := len(letters) - 1; r >= 0; r-- {
			revbukva = append(revbukva, letters[r])
		}
		revwords = append(revwords, strings.Join(revbukva, ""))

	}
	return strings.Join(revwords, " ")
}
