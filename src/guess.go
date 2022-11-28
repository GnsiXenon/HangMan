package src

func Guess(word string, guess string, inWord []bool, stock string, attempts int) (string, int, []bool, string, string) { //advenced word attemtps
	// declaring variables
	validGuess := false
	stock = stockGuess(guess, stock)
	txt := ""
	// if our file is Save.txt, save the game.
	// otherwise, we pick a random letter to display and start the game
	// initializing a loop for each attempt
	// evlautating, printing, scanning, and storing based on the guess
	guess, _ = Isletter(guess)
	inWord, validGuess = InWord(word, guess, inWord) // reinitializing boolean variables based on Found()
	if attempts == 0 && !Success(inWord) {           // case: loss
		txt = "Perdu"
	} else if Success(inWord) { // case: win
		txt = "Félicitations ! Vous avez gagnez !"
	} else if validGuess { // case: valid single letter guess
		txt = " Tu l'as eu ! Continue !"
	} else { // case: invalid single letter guess & not a word
		txt = " Mauvaise réponse ! Essaye encore !"
		attempts--
	}

	return Prints(word, inWord), attempts, inWord, stock, txt

}

/*func Guess(word string, hangman string, file string) {
	// declaring variables
	var attempts = 10
	var inWord = []bool{}
	stock := ""
	// if our file is Save.txt, save the game.
	// otherwise, we pick a random letter to display and start the game
	if file == "Save.txt" {
		word, stock, attempts, inWord, file = GetSave(file)
	} else {
		inWord = make([]bool, len(word))
		for n := 0; n < len(word)/2-1; n++ {
			inWord[RandomLetter(word)] = true
		}
	}
	isLetter := false
	fmt.Println("\033[2J\033[1;1H")
	fmt.Println("You are now playing HANGMAN, you have 10 attempts total. Good luck!")

	// initializing a loop for each attempt
	// evlautating, printing, scanning, and storing based on the guess
	for attemptsRemaining := attempts; attemptsRemaining <= 11 && attemptsRemaining >= 1; attemptsRemaining-- {
		Prints(word, inWord) // displays platorm with letters or _
		fmt.Printf(" You have %v attempts. \n", attemptsRemaining)
		fmt.Print("Guess: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		guess := scanner.Text()           // taking user input
		guess, isLetter = Isletter(guess) // evaluating if user input is a letter
		if !isLetter {
			fmt.Println("Enter a valid letter: ")
		} else {
			Save(guess, word, stock, attemptsRemaining, inWord, file)           // save game
			inWord, validGuess, validAll, isWord := InWord(word, guess, inWord) // reinitializing boolean variables based on Found()
			fmt.Println("\033[2J\033[1;1H")

			stock = stockGuess(guess, stock) // store guess in array
			fmt.Println(stock)
		}
		GetHangman(hangman, attemptsRemaining) // display hangman based on attempts
		fmt.Println(stock)
	}
}
*/
