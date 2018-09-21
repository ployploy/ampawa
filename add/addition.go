package add

import "strconv"

func Addition(inputNumber1 []int, inputNumber2 []int) string {
	var result int
	var tod int
	var number, s string

	for index := len(inputNumber1) - 1; index >= 0; index-- {
		result = inputNumber1[index] + inputNumber2[index] + tod
		sufer := result - 10
		if sufer >= 0 {
			tod = 1
			s = strconv.Itoa(sufer)
		}
		if sufer < 0 {
			tod = 0
			s = strconv.Itoa(result)
		}
		number = s + number

	}
	if tod == 1 {
		number = "1" + number
	}
	return number
}
