package problems

import (
		"fmt"
		"strings"
		"strconv"
)

var ones, teens, tens, hundreds, thousands map[int]string

func Problem17() {
	ones = map[int]string {
		1: "one", 
		2: "two", 
		3: "three", 
		4: "four", 
		5: "five", 
		6: "six", 
		7: "seven", 
		8: "eight", 
		9: "nine",
	}

	teens = map[int]string {
		10: "ten", 
		11: "eleven", 
		12: "twelve", 
		13: "thirteen", 
		14: "fourteen", 
		15: "fifteen", 
		16: "sixteen", 
		17: "seventeen", 
		18: "eighteen", 
		19: "nineteen",
	}

	tens = map[int]string {
		2: "twenty",
		3: "thirty",
		4: "forty",
		5: "fifty",
		6: "sixty",
		7: "seventy",
		8: "eighty",
		9: "ninety",
	}

	hundreds = map[int]string {
		1: "onehundred",
		2: "twohundred",
		3: "threehundred",
		4: "fourhundred",
		5: "fivehundred",
		6: "sixhundred",
		7: "sevenhundred",
		8: "eighthundred",
		9: "ninehundred",
	}

	thousands = map[int]string {
		1: "onethousand",
	}

	res := 0

	for i := 1; i <= 1000; i++ {
		if i < 10 {
			res += len(ones[i])
		} else if i >= 10 && i < 20 {
			res += len(teens[i])
		} else if i >= 20 && i < 100 {
			res += getLengthForTens( i )
		} else if i >= 100 && i < 1000 {
			res += getLengthForHundreds( i )
		} else {
			res += len( thousands[1] )
		}
	} 

	fmt.Println( res )
}

func getLengthForTens( i int ) ( res int ) {
	parts := strings.Split((fmt.Sprintf("%d", i)), "")

	ten := parseInt( parts[0])
	res += len( tens[ int(ten) ] )

	one := parseInt( parts[1])
	res += len( ones[ int(one) ] )

	return
}

func getLengthForHundreds( i int ) ( res int ) {
	parts := strings.Split((fmt.Sprintf("%d", i)), "")

	hundred := parseInt( parts[0])
	res += len( hundreds[ int(hundred) ] )

	if parseInt( parts[1] + parts[2]) != 0 {
		res += len( "and" )
	}

	if parts[1] == "1" {
		teen := parseInt( parts[1] + parts[2])
		res += len( teens[ int(teen) ] )
	} else {
		ten := parseInt( parts[1])
		res += len( tens[ int(ten) ] )

		one := parseInt( parts[2])
		res += len( ones[ int(one) ] )
	}

	return
}

func parseInt(str string) (res int64) {
	res, _ = strconv.ParseInt( str, 10, 32)
	return
}