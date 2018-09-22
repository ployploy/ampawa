package add

import "strconv"

func Addition(inputNumber1 []int, inputNumber2 []int) string {
	var plus int
	var tod int
	var number, result string

	for index := len(inputNumber1) - 1; index >= 0; index-- {
		plus = inputNumber1[index] + inputNumber2[index] + tod
		sufer := plus - 10
		if sufer >= 0 {
			tod = 1
			result = strconv.Itoa(sufer)
		}
		if sufer < 0 {
			tod = 0
			result = strconv.Itoa(plus)
		}
		number = result + number

	}
	if tod == 1 {
		number = "1" + number
	}
	return number
}
