package main

import "fmt"

//Checks if a string is a palindrome
func checkPalindrome(inputString string) bool {
	i := 0
	j := len(inputString) - 1
	for i < j {
		if inputString[i] != inputString[j] {
			return false
		}
		i++
		j--
	}
	return true
}
func main() {
	s := "aabaa"
	v := checkPalindrome(s)
	fmt.Println(v)
}
