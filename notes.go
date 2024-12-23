package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	x := 1
	counter := 0
	round := 1

	incorrectList := []string{}
	wordList := []string{}
	var correctCodes = [][]string{{}, {}, {}, {}}
	intList := []int{}

	for x > 0 {
		fmt.Println("Round:", round)

		fmt.Println("-------------------------------------")
		if counter < 3 {
			for i := 0; i < 3; i++ {
				getNewWord(&wordList)
				counter++
			}
		}

		if counter == 3 {

			fmt.Println("-------------------------------------")
			printWordList(wordList)

			fmt.Print("Enemy's Numbers: ")
			number, _ := reader.ReadString('\n')

			number = number[:len(number)-1]
			// windows append "\r\n" to the end of input
			if strings.Contains(number, "\r") {

				number = number[:len(number)-1]

			}

			fmt.Println(number)

			for i := 0; i < 3; i++ {
				// Convert the string to an integer

				num, err := strconv.Atoi(string(number[i]))

				fmt.Println(err)

				intList = append(intList, num)
			}

			fmt.Println("-------------------------------------")

			fmt.Println("\nEnemy Guess:")
			for i := 0; i < 3; i++ {
				fmt.Println(wordList[i], " -> ", intList[i])
			}
			fmt.Println()
			counter++

		}

		if counter == 4 {

			//print correct list
			printCorrect(correctCodes)

			//print incorrect List
			printIncorrect(incorrectList)

			fmt.Println("are they correct?")
			fmt.Println("1: Yes")
			fmt.Println("2: No")

			input, _ := reader.ReadString('\n')
			input = input[:len(input)-1]
			userIn, _ := strconv.Atoi(input)

			fmt.Println()

			if userIn == 1 || input == "1\r" {
				// add to enemy list
				for i := 0; i < 3; i++ {
					correctCodes[intList[i]-1] = append(correctCodes[intList[i]-1], wordList[i])
				}

				counter = 0
				wordList = []string{}
				intList = []int{}

				printCorrect(correctCodes)
				printIncorrect(incorrectList)
				fmt.Println("-------------------------------------")
				round++
			} else if userIn == 2 || input == "2\r" {
				// add to wrong guesses
				for i := 0; i < 3; i++ {
					incorrectList = append(incorrectList, wordList[i])
				}

				counter = 0
				wordList = []string{}
				intList = []int{}

				printCorrect(correctCodes)
				printIncorrect(incorrectList)
				fmt.Println("-------------------------------------")
				round++
			}

		}

	}

}

func getNewWord(wordList *[]string) {

	var word string

	fmt.Print("Enter Opponent's Word: ")
	word, _ = reader.ReadString('\n')
	word = word[:len(word)-1]

	*wordList = append(*wordList, word)

}

func printWordList(wordList []string) {

	for i := 0; i < 3; i++ {
		fmt.Println(i+1, ": ", wordList[i])
	}

}

func printCorrect(correctCodes [][]string) {
	fmt.Println("Enemy Correct Guesses")
	for i := 0; i < 4; i++ {

		out, _ := fmt.Print(i+1, ": ")
		fmt.Println(out)

		for j := 0; j < len(correctCodes[i]); j++ {
			fmt.Printf("%s, ", correctCodes[i][j])
		}
		fmt.Printf("\n\n")
	}
}

func printIncorrect(incorrectList []string) {
	fmt.Println("Enemy Incorrect Guesses")
	for i := 0; i < len(incorrectList); i++ {

		fmt.Printf("%s, ", incorrectList[i])

	}
	fmt.Printf("\n\n")
}
