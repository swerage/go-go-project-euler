package problems

import (
		"fmt"
		"time"
		"os"
)

func Problem12() {
	var triangleNumber, numberOfFactors, highestNumberOfFactors int

	i := 1
	start := time.Now()


	for highestNumberOfFactors < 500 {
		triangleNumber = getTriangleNumberFor(i)
		numberOfFactors = factorize(triangleNumber)

		if numberOfFactors > highestNumberOfFactors {
			highestNumberOfFactors = numberOfFactors
		}
		i += 1
		fmt.Println(i)
	}

	end := time.Now().Sub(start)

	fmt.Printf("Triangle number %d, number of factors %d. Time %v ", 
		triangleNumber, numberOfFactors, end.Minutes())
}

func getTriangleNumberFor(num int) (result int) {
	for i := 0; i <= num; i++ {
		result += i
	}
	return
}

func factorize(triNum int ) (factors int ) {
	if (triNum == 2147483647) {
		os.Exit(1)
	}

	factors = 0
	for i := 1; i <= triNum; i++ {
		if triNum % i == 0 {
			factors += 1
		}
	}
	return
}