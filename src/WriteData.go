package hangman

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//Permet de rajouter un mots a la base de donné de la dificulté
func AddWord(file string, word string) (string, bool) {
	word, valid := Isletter(word) //Permet de vérifier si le mot a des caractere spéciaux
	for !valid {
		return "Entrer un seul mots", false //si il possede des caracteres speciaux
	}
	word = "\n" + word

	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600) //Ouvre le fichier de la difficulté choisi
	if err != nil {
		panic(err)
	}

	if _, err = f.WriteString(word); err != nil { //ecrit le mot a la suite dans le fichier
		panic(err)
	}
	f.Close()
	return "Mots ajouter", true

}

func Save(guess string, word string, stock string, attemptsRemaining int, found []bool) {

	//Creation du fichier et ajout du mot
	wordByte := []byte(word + "\n")
	err := ioutil.WriteFile("Save.txt", wordByte, 0777)
	if err != nil {
		fmt.Print(err)
	}
	//ouverture du fichier pour pouvoir ecrire dedans
	f, err := os.OpenFile("Save.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	//Ecrie dans le fichier
	//Stock
	if _, err = f.WriteString(stock + "\n"); err != nil {
		panic(err)
	}

	//attemptsRemaining
	attemptsRemainingString := strconv.Itoa(attemptsRemaining)
	if _, err = f.WriteString(attemptsRemainingString + "\n"); err != nil {
		panic(err)
	}

	//found
	for _, element := range found {
		boolString := strconv.FormatBool(element)
		if _, err = f.WriteString(boolString + " "); err != nil {
			panic(err)
		}

		f.Close()
	}
}
