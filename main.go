package main

import (
	"fmt"
	"math/rand"
)

func babyMode() {
	number := rand.Intn(100)

	var guess int = 0
	for guess != number {
		fmt.Println("Enter guess: ")
		fmt.Scan(&guess)
		if guess < number {
			fmt.Println("Guess higher!")
		} else {
			if guess > number {
				fmt.Println("Guess lower !")
			} else {
				fmt.Println("You win!")
			}
		}
	}
}

func easyMode() {
	number := rand.Intn(100)
	lives := 10

	var guess int = 0
	for guess != number && lives > 0 {
		lives--
		fmt.Println("Enter guess: ")
		fmt.Scan(&guess)
		if guess < number {
			fmt.Println("Guess higher!")
		} else {
			if guess > number {
				fmt.Println("Guess lower !")
			} else {
				fmt.Println("You win!")
			}
		}
	}
	fmt.Println("You ran out of moves!")
}

func normalMode() {

	number := rand.Intn(100)
	lives := 7

	var guess int = 0
	for guess != number && lives > 0 {
		lives--
		fmt.Println("Enter guess: ")
		fmt.Scan(&guess)
		if guess < number {
			fmt.Println("Guess higher!")
		} else {
			if guess > number {
				fmt.Println("Guess lower !")
			} else {
				fmt.Println("You win!")
			}
		}
	}
	fmt.Println("You ran out of moves!")
}

func hardMode() {

	number := rand.Intn(100)
	lives := 5

	var guess int = 0
	for guess != number && lives > 0 {
		lives--
		fmt.Println("Enter guess: ")
		fmt.Scan(&guess)
		if guess < number {
			fmt.Println("Guess higher!")
		} else {
			if guess > number {
				fmt.Println("Guess lower !")
			} else {
				fmt.Println("You win!")
			}
		}
	}
	fmt.Println("You ran out of moves!")
}

func extremeMode() {

	number := rand.Intn(100)
	lives := 3

	var guess int = 0
	for guess != number && lives > 0 {
		lives--
		fmt.Println("Enter guess: ")
		fmt.Scan(&guess)
		if guess < number {
			fmt.Println("Guess higher!")
		} else {
			if guess > number {
				fmt.Println("Guess lower !")
			} else {
				fmt.Println("You win!")
			}
		}
	}
	fmt.Println("You ran out of moves!")
}

func impossibleMode() {

	number := rand.Intn(100)
	lives := 1

	var guess int = 0
	for guess != number && lives > 0 {
		lives--
		fmt.Println("Enter guess: ")
		fmt.Scan(&guess)
		if guess == number {
			fmt.Println("You win. Cheater.")
		}
	}
	fmt.Println("L.")
}

func main() {
	var mode string
	fmt.Println("Please enter mode from the list: " +
		"baby\n" +
		"easy\n" +
		"normal\n" +
		"hard\n" +
		"extreme\n" +
		"impossible\n" +
		"Input:")
	fmt.Scan(&mode)
	switch mode {
	case "baby":
		babyMode()
		break
	case "normal":
		normalMode()
		break
	case "easy":
		easyMode()
		break
	case "hard":
		hardMode()
		break
	case "extreme":
		extremeMode()
		break
	case "impossible":
		impossibleMode()
		break

	default:
		fmt.Println("Unknown gamemode.")
	}

}
