package services

import "github.com/arjunajithtp/excel-column-finder/internal/helpers"

func LetterIncrementation(startingColumn string, count int) []string {
	var words []string
	startingLen := len(startingColumn)
	if startingLen > 0 {
		for i := 0; i < count; i++ {
			wordRune := []rune(helpers.Reverse(startingColumn))
			var newRunes [][]rune
			var newLetter []rune
			increment := rune(i)
			for j := 0; j < len(wordRune); j++ {
				oldByte := wordRune[j]
				newByte, remainder := NewLetter(oldByte, increment)
				newLetter = append(newLetter, newByte)
				increment = remainder

				if j == len(wordRune) - 1 && increment > 0{
					newByte, remainder = NewLetter(64, increment)
					newLetter = append(newLetter, newByte)
				}
				newRunes = append(newRunes, newLetter)
			}
			for k := 0; k < len(newRunes); k++ {
				words = append(words, helpers.Reverse(string(newRunes[k])))
			}
		}
	}
	return words
}

func NewLetter(oldByte, increment rune) (rune, rune) {
	newByte := oldByte + increment
	if newByte > 90 {
		r := newByte % 90
		if r > 26{
			if r % 26 == 0 {
				return 90, (newByte / 90) + (r / 26) - 1
			}
			return 64 + (r % 26), (newByte / 90) + (r / 26)
		}
		return 64 + r, newByte / 90
	}
	return newByte, 0
}
