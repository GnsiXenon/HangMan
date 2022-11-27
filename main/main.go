package main

import (
	"src"
)

func main() {
	file := src.DisplayHome()
	if file == "Save.txt" {
		src.Guess("Save", "src.txt", "Save.txt")

	} else {
		wordsArray := src.GetWords(file)
		words := src.Random(wordsArray)
		src.Guess(words, "src.txt", file)
	}
}
