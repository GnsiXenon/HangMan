package hangman

func InWord(word string, guess string, InWord []bool) ([]bool, bool) {
	guessRune := []rune(guess)
	ValidGuess := false

	if len(guess) <= 1 { // case: guess is a single letter
		for i, element := range word { // loop through the word and compare the letter
			if guessRune[0] == element { // case: letter has been InWord at any given index in the word
				InWord[i] = true //case: boolean value is true
				ValidGuess = true

			}
		}
	}
	return InWord, ValidGuess
}
