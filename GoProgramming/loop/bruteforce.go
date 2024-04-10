package loop
import "fmt"

func FindBruteForceAlgorithm(setOfArray []int, number int) {
	var iterationNumber int = 1
	for i := 0; i < len(setOfArray); i++ {

		if setOfArray[i] == number {
			fmt.Println("Researched number: ", setOfArray[i], " ", "Iteratıon number: ", iterationNumber)
		}
		iterationNumber++
	}
}
