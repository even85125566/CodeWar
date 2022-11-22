package codewar

func Mixbonacci(pattern []string, length int) []int64 {
	orilength := len(pattern)
	newpattern := []string{}
	resultlist := []int64{}
	patternCountMap := make(map[string]int)
	for i := 0; i < length/orilength; i++ {
		newpattern = append(newpattern, pattern...)
	}

	for _, eachpattern := range newpattern {

		if _, ok := patternCountMap[eachpattern]; !ok {
			patternCountMap[eachpattern] = 0
		}
		result := Transform(eachpattern, patternCountMap[eachpattern])
		resultlist = append(resultlist, result)
		patternCountMap[eachpattern]++
	}
	return resultlist
}

func Transform(pattern string, patternCount int) int64 {
	switch pattern {
	case "fib":
		return int64(Fibonacci(patternCount))
	case "pad":
		return int64(Padovan(patternCount))
	case "jac":
		return int64(Jacobstahi(patternCount))
	case "pel":
		return int64(Pell(patternCount))
	case "tri":
		return int64(Tribonacci(patternCount))
	case "tet":
		return int64(Tetranacci(patternCount))
	default:
		return 0
	}
}

func Fibonacci(number int) int {
	if number < 2 {
		return number
	}

	return Fibonacci(number-1) + Fibonacci(number-2)
}
func Pell(number int) int {
	if number < 2 {
		return number
	}

	return 2*Pell(number-1) + Pell(number-2)

}
func Padovan(number int) int {

	switch number {
	case 0:
		return 1
	case 1, 2:
		return 0

	}
	return Padovan(number-2) + Padovan(number-3)
}

func Jacobstahi(number int) int {
	if number < 2 {
		return number
	}

	return Jacobstahi(number-1) + 2*Jacobstahi(number-2)
}

func Tribonacci(number int) int {
	if number < 2 {
		return 0
	}
	if number == 2 {
		return 1
	}

	return Tribonacci(number-1) + Tribonacci(number-2) + Tribonacci(number-3)

}
func Tetranacci(number int) int {

	switch number {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	case 3:
		return 2
	default:
		return Tetranacci(number-1) + Tetranacci(number-2) + Tetranacci(number-3) + Tetranacci(number-4)

	}

}
