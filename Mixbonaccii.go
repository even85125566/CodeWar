package codewar

func Mixbonacci(pattern []string, length int) []int64 {
	resultlist := []int64{}
	patternCountMap := make(map[string]int)
	if length == 0 || len(pattern) == 0 {
		return resultlist
	}
	for i := 0; i < length; i++ {
		r := i % len(pattern)
		thePattern := pattern[r]
		if _, ok := patternCountMap[thePattern]; !ok {
			patternCountMap[thePattern] = 0
		}
		result := Transform(pattern[r], patternCountMap[thePattern])
		patternCountMap[thePattern]++
		resultlist = append(resultlist, result)
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
		return 0
	case 2:
		return 0
	case 3:
		return 1
	default:
		return Tetranacci(number-1) + Tetranacci(number-2) + Tetranacci(number-3) + Tetranacci(number-4)

	}

}
