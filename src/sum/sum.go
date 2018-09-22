package sum

func Summation(input []int) int {
	var result int
	for i, _ := range input {
		result += input[i]
	}
	return result
}
