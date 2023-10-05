package main

func majorityElement(nums []int) []int {
	firstCandidate, secondCandidate, firstCount, secondCount := 0, 1, 0, 0

	for _, v := range nums {
		if v == firstCandidate {
			firstCount++
		} else if v == secondCandidate {
			secondCount++
		} else if firstCount == 0 {
			firstCandidate = v
			firstCount = 1
		} else if secondCount == 0 {
			secondCandidate = v
			secondCount = 1
		} else {
			firstCount--
			secondCount--
		}
	}

	firstCount, secondCount = 0, 0

	for _, v := range nums {
		if v == firstCandidate {
			firstCount++
		}
		if v == secondCandidate {
			secondCount++
		}
	}

	var result []int

	if firstCount > len(nums)/3 {
		result = append(result, firstCandidate)
	}

	if secondCount > len(nums)/3 {
		result = append(result, secondCandidate)
	}

	return result
}
