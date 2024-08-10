package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"unicode"
)

func main() {
	var stringCode string
	var delta int
	var code bool
	var decode bool
	var help bool
	var ret string
	accentedChars := "éèêëàâîïôùûç"

	flag.BoolVar(&code, "code", false, "Afin de coder un msg")
	flag.BoolVar(&decode, "decode", false, "afin de decoder un msg")
	flag.BoolVar(&help, "help", false, "Voir les options")
	flag.IntVar(&delta, "delta", 14, "(optionel) choisir le delta pour le cryptage")
	flag.Parse()
	if code {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("--- Info: Pas d'accent et remplace les espaces par [-] ---\nEcrit le message a coder:")
		scanner.Scan()
		stringCode = strings.Trim(scanner.Text(), " ")
		switch strings.Index(stringCode, " ") {
		case -1:
		default:
			log.Fatal("il ne faut pas mettre d'espace dans le message")
		}
		for _, char := range stringCode {
			if strings.ContainsRune(accentedChars, char) {
				log.Fatal("Erreur--le message comporte un accent--")
			}
		}
		ret += cipher(string(stringCode), delta)
		fmt.Printf("Voici le message coder:\n%s\n", ret)
	}
	if decode {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Rentre le message a decoder:")
		scanner.Scan()
		stringCode = strings.Trim(scanner.Text(), " ")
		ret += DecodeCipher(string(stringCode), delta)
		fmt.Printf("Voici le message décoder:\n%s\n", ret)
	} 
	if help {
		flag.Usage()
	}else if !code && !decode &&!help{
		fmt.Println("tapez -help pour voir les cmd")
	}

}

func CamelCaseCount(str string) int {
	count := 1
	for i := range str {
		if string(str[i]) == strings.ToUpper(string(str[i])) {
			count++
		}
	}

	return count
}

func cipher(str string, rotate int) string {
	OriginalAlphabet := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}

	maxLength := len(OriginalAlphabet)
	cypher := ""

	for _, char := range str {
		switch unicode.IsLetter(rune(char)) {
		case true:
			index := slices.Index(OriginalAlphabet, strings.ToLower(string(char)))
			if index != -1 {
				indexBis := (index + rotate) % maxLength
				// 		A retenir !!!!! Exemple Illustratif modulo dans debordement circulaire
				// Index de "z": 25
				// Rotation: 3
				// Calculons le nouvel index :
				// Calcul de index + rotate:
				// 25 + 3 = 28
				// Application du Modulo:
				// 28 % 26 = 2
				// Cela signifie que "z" (index 25) devient "c" (index 2) après une rotation de 3 positions.
				// Conclusion
				// L'utilisation de l'opération modulo dans l'expression indexBis := (index + rotate) % maxLength
				// est cruciale pour gérer les débordements d'index dans un alphabet circulaire.
				// Cela permet de s'assurer que, peu importe la valeur de rotate,
				// le nouvel index reste valide et correspond à un caractère de l'alphabet.
				// Cette technique est couramment utilisée dans les algorithmes de chiffrement
				// et d'autres applications nécessitant des calculs circulaires.
				if unicode.IsUpper(rune(char)) {
					cypher += strings.ToUpper(OriginalAlphabet[indexBis])
				} else {
					cypher += OriginalAlphabet[indexBis]
				}
			}
		case false:
			cypher += string(char)
		}
	}

	return cypher
}

func DecodeCipher(str string, rotate int) string {
	OriginalAlphabet := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}

	maxLength := len(OriginalAlphabet)
	cypher := ""

	for _, char := range str {
		switch unicode.IsLetter(rune(char)) {
		case true:
			index := slices.Index(OriginalAlphabet, strings.ToLower(string(char)))
			indexBis := (index - rotate + maxLength) % maxLength
			// 		A retenir !!!!! Exemple Illustratif modulo dans debordement circulaire
			// Index de "z": 25
			// Rotation: 3
			// Calculons le nouvel index :
			// Calcul de index + rotate:
			// 25 + 3 = 28
			// Application du Modulo:
			// 28 % 26 = 2
			// Cela signifie que "z" (index 25) devient "c" (index 2) après une rotation de 3 positions.
			// Conclusion
			// L'utilisation de l'opération modulo dans l'expression indexBis := (index + rotate) % maxLength
			// est cruciale pour gérer les débordements d'index dans un alphabet circulaire.
			// Cela permet de s'assurer que, peu importe la valeur de rotate,
			// le nouvel index reste valide et correspond à un caractère de l'alphabet.
			// Cette technique est couramment utilisée dans les algorithmes de chiffrement
			// et d'autres applications nécessitant des calculs circulaires.
			if index != -1 {
				if unicode.IsUpper(rune(char)) {
					cypher += strings.ToUpper(OriginalAlphabet[indexBis])
				} else {
					cypher += OriginalAlphabet[indexBis]
				}
			}
		case false:
			cypher += string(char)
		}
	}

	return cypher
}
