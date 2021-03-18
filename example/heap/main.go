package main

import "fmt"

func main() {
	glassOfWater := answers()
	// fortyTwo is a candidate for deletion
	fmt.Println(*glassOfWater)
}

func answers() *string {
	answerMap := make(map[string]*string)
	fortyTwo := "42"
	answerMap["What do you get if you multiply six by nine?"] = &fortyTwo
	glassOfWater := "Ask a glass of water."
	answerMap["What's so unpleasant about being drunk?"] = &glassOfWater
	return &glassOfWater
}
