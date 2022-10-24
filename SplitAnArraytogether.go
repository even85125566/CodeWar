package codewar

//You will receive an array as parameter that contains 1 or more integers and a number n.

func SplitAndAdd(numbers []int, n int) []int {
	if n == 0 {
		return numbers
	}

	var alist []int
	var blist []int
	var clist []int
	for i := 0; i < len(numbers)/2; i++ {
		alist = append(alist, numbers[i])
	}
	for i := len(numbers) / 2; i < len(numbers); i++ {
		blist = append(blist, numbers[i])
	}
	for i := 0; i < len(blist); i++ {
		if len(blist) != len(alist) {
			if i == 0 {
				clist = append(clist, blist[0])
				continue
			}
			clist = append(clist, blist[i]+alist[i-1])
		} else {
			clist = append(clist, alist[i]+blist[i])
		}

	}
	return SplitAndAdd(clist, n-1)
}
