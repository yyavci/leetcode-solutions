package main

import "fmt"

func main() {
	res := romanToInt("III") //3
	fmt.Printf("III -> %d\n", res)

	res = romanToInt("LVIII") //58
	fmt.Printf("LVIII -> %d\n", res)

	res = romanToInt("LIVII") //56
	fmt.Printf("LIVII -> %d\n", res)

	res = romanToInt("MCMXCIV") //1994
	fmt.Printf("MCMXCIV -> %d\n", res)
}

/*
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
*/

func romanToInt(s string) int {

	var m = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	res := 0
	last := 0

	for i := len(s) - 1; i >= 0; i-- {

		val := m[string(s[i])]
		if val < last && val > 0 {
			//subtract
			res -= val
		} else {
			//add
			res += val
		}
		last = val
	}

	return res

}
