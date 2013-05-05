package problems

import (
		"fmt"
		"math"
		"strings"
		"strconv"
)

func Problem16() {
	var res int64

	prod := math.Pow(2, 1000)
	digits := strings.Split((fmt.Sprintf("%f", prod)), "")

	for i := 0; i < len(digits); i++ {
		if num, err := strconv.ParseInt(digits[i], 10, 64); err == nil {
			res += num
		}
	}

	fmt.Println( fmt.Sprintf("Result is %v", res) )
}