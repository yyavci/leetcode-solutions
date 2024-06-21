package main

import "fmt"

func main() {
	res := isPalindrome(121)
	fmt.Println(res)

	res2 := isPalindrome(-121)
	fmt.Println(res2)

	res3 := isPalindrome(10)
	fmt.Println(res3)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := fmt.Sprint(x)
	var res = ""
	for i := len(s) - 1; i >= 0; i-- {
		c := string(s[i])
		res += c
	}
	return res == s
}
