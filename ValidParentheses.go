package codewar

import "fmt"

//Write a function that takes a string of parentheses,
//and determines if the order of the parentheses is valid.
//The function should return true if the string is valid, and false if it's invalid.
func ValidParentheses(parens string) bool {
	var leftside rune = 40
	var rightside rune = 41
	leftsideCount := 0
	rightsideCount := 0
	for _, eachstr := range parens {
		switch eachstr {
		case leftside:
			leftsideCount++
		case rightside:
			rightsideCount++
		default:
			fmt.Println("unexpected input:", string(eachstr))
		}

		if leftsideCount < rightsideCount {
			return false
		}
	}

	return leftsideCount == rightsideCount
}
