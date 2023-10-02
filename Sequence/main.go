package main

func IsSequence(part string, sentence string) bool {
	sequenceIndex := 0

	for i := 0; i < len(sentence); i++ {
		if sequenceIndex == len(part){
			return true
		}
		if sentence[i] == part[sequenceIndex] {
			sequenceIndex++
		}
	}
	return false
}