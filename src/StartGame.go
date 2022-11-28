package src

func StarGame(file string) (string, []bool) {
	wordsArray := GetWords(file)
	words := Random(wordsArray)
	mem := make([]bool, len(words))
	//
	for n := 0; n < len(words)/2-1; n++ {
		mem[RandomLetter(words)] = true
	}

	return words, mem
}
