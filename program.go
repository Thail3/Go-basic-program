package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter\nyour name: ")
	var name string

	fmt.Scan(&name)
	fmt.Printf("Hello, %v , Welcome to the game!\n", name)

	fmt.Printf("Enter\nyour age: ")
	var age int
	fmt.Scan(&age)
	if age > 10 {
		fmt.Printf("You can play!")
	} else {
		fmt.Printf("You can't play!")
	}

	
	score := 0
	num_questions := 2


	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer + " " + answer2 == "RTX 3090" {
		fmt.Println("Correct answer")
		score++
	} else {
		fmt.Printf("Incorrect")
	}

	var cores int
	fmt.Printf("How many cores does the Ryzen 9 have? ")
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct answer")
		score++
	} else {
		fmt.Printf("Incorrect")
	}

	fmt.Printf("Your score %v out of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v%%.", percent)
}
