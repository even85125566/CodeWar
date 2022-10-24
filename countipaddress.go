package codewar

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)


//Implement a function that receives two IPv4 addresses, and returns the number of addresses between them (including the first one, excluding the last one).
//All inputs will be valid IPv4 addresses in the form of strings. The last address will always be greater than the first one.


func IpsBetween(start, end string) int {
	return int(IPTest(start, end))

}
func IPTest(start string, end string) (number float64) {
	startString := strings.Split(start, ".")
	endString := strings.Split(end, ".")

	startIP := getNumber(startString)
	endIP := getNumber(endString)
	for i := 0; i < 4; i++ {

		number += math.Pow(float64(256), float64(3-i)) * float64(endIP[i]-startIP[i])

	}
	return

}
func getNumber(stringSlice []string) (numberslice []int) {
	for _, eachstring := range stringSlice {
		eachNumber, err := strconv.Atoi(eachstring)
		if err != nil {
			fmt.Println(err)
			return
		}
		numberslice = append(numberslice, eachNumber)
	}

	return

}
