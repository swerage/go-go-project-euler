package problems

import "fmt"

func Problem14() {
	longestChain, longestStart, length, i := int64(0), int64(0), int64(0), int64(0)
	var chain map[int64]int64

	for i = 13; i < 1000000; i++ {
		chain = getChainFor(int64(i))
		length = int64(len(chain))
		chain = nil

		fmt.Println(i)
		if length > longestChain {
			longestChain = length
			longestStart = int64(i)
		}
	}

	fmt.Println("longest", longestStart)
}

func getChainFor(n int64) map[int64]int64 {
	chain := make(map[int64]int64)
	chain[0] = n

	for n != int64(1) {
		if n % int64(2) == 0 {
			n = even(n)
		} else {
			n = odd(n)
		}
		chain[int64(len(chain))] = n
	} 

	return chain
}

func odd(n int64) int64 {
	var v = int64(3) * n + int64(1)
	return v
}

func even(n int64) int64 {
	var v = n / int64(2)
	return v
}