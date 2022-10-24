package codewar

//Your task, is to create NxN multiplication table, of size provided in parameter.

func MultiplicationTable(size int) [][]int {
	var theResult [][]int
	if size == 0 {
		theResult := [][]int{}

		return theResult
	}
	for i := 1; i <= size; i++ {
		var testlist []int
		for j := 1; j <= size; j++ {

			testlist = append(testlist, i*j)

		}

		theResult = append(theResult, testlist)
	}
	return theResult
}
