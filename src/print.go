package hangman

func Prints(word string, found []bool) string {
	runeWord := []rune(word) // converting our word -string to runes
	wordString := ""
	for i := 0; i < len(runeWord); i++ { // looping through the word in rune form
		if found[i] { // case: letter is found/true at any given index
			wordString += string(runeWord[i])
		} else {
			wordString += "_"
		}
	}
	return wordString
}
