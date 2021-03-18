package main

import (
	"os"
)

func main() {
	question := "What do you get if you multiply six by nine?"
	theAnswer := getAnswer(question)
	// answer variable deleted here
	os.Exit(theAnswer)
}

func getAnswer(question string) int {
	answer := 42
	return answer
}
