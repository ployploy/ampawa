package sum

func Summation(input []int) int {
	var result int
	for index, _ := range input {
		if index+1 <= len(input)-1 {
			input[index+1] = input[index] + input[index+1]
			result = input[index+1]
		}

	}
	return result
}
