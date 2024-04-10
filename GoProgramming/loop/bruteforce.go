package loop

func FindBruteForceAlgorithm(setOfArray []int, number int) bool {
	for i := 0; i < len(setOfArray); i++ {
		if setOfArray[i] == number {
			return true
		}
	}
	return false
}}
