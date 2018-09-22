package sum

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Readfile(path string) []int {
	var inputs []int
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		inputs = append(inputs, number)
	}

	return inputs

}

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
