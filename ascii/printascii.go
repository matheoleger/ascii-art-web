package ascii

import "fmt"

func PrintAscii(eachTextSplit string, asciiLines []string) {

	characterLine := 0

	for everyLines := 0; everyLines < 8; everyLines++ { // Boucle permettant de mettre chaque ligne

		str := ""

		for _, char := range eachTextSplit { //Boucle regardant chaque string du tableau de string

			characterLine = (int(char) - 32) * 9             //cherche la première ligne du caractère correspondant
			str = str + asciiLines[characterLine+everyLines] // concaténation de la ligne des différents caracères

		}
		fmt.Println(str)
	}
}
