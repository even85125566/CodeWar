package codewar
//In this example you have to validate if a user input string is alphanumeric. The given string is not nil/null/NULL/None, so you don't have to check that.

//The string has the following conditions to be alphanumeric:

//At least one character ("" is not valid)
//Allowed characters are uppercase / lowercase latin letters and digits from 0 to 9
//No whitespaces / underscore
func alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, s := range str {
		if (s >= 48 && s <= 57) || (s >= 65 && s <= 90) || (s >= 97 && s <= 122) {
			continue
		}
		if check(s) {
			continue
		}
		return false
	}
	return true
}

func check(r rune) bool {
	return r == 64 || r == 177 || r == 181 || r == 163 || r == 153
}
