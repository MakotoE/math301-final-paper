package main

func main() {
	_ = answers()
	// fortyTwo is a candidate for deletion
}

func answers() *string {
	answerMap := make(map[string]*string, 100)
	fortyTwo := "42"
	answerMap["What do you get if you multiply six by nine?"] = &fortyTwo
	glassOfWater := "Ask a glass of water."
	answerMap["What's so unpleasant about being drunk?"] = &glassOfWater
	return &glassOfWater
}
