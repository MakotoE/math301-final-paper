package main

func main() {
	doStuff()
	// fortyTwo is a candidate for deletion
}

func doStuff() {
	answerMap := make(map[string]*int, 100)
	getAnswer(answerMap)
}

func getAnswer(answerMap map[string]*int) {
	fortyTwo := 42
	answerMap["What do you get if you multiply six by nine?"] = &fortyTwo
}

/*
answerMap -> "What ... nine?" -> &fortyTwo
*/
