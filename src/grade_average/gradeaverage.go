package grade_average

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

func GradeAverage(inputPoint []int) float64 {
	var four, threeDotFive, three, thwoDotFive, two, oneDotFive, one, zeroDotFive, zero float64
	for index, _ := range inputPoint {
		if inputPoint[index] >= 85 && inputPoint[index] <= 100 {
			four += 4.0
		}
		if inputPoint[index] >= 80 && inputPoint[index] < 85 {
			threeDotFive += 3.5
		}
		if inputPoint[index] >= 75 && inputPoint[index] < 80 {
			three += 3.0
		}
		if inputPoint[index] >= 70 && inputPoint[index] < 75 {
			thwoDotFive += 2.5
		}
		if inputPoint[index] >= 65 && inputPoint[index] < 70 {
			two += 2.0
		}
		if inputPoint[index] >= 60 && inputPoint[index] < 65 {
			oneDotFive += 1.5
		}
		if inputPoint[index] >= 55 && inputPoint[index] < 60 {
			one += 1.0
		}
		if inputPoint[index] >= 50 && inputPoint[index] < 55 {
			zeroDotFive += 0.5
		}
		if inputPoint[index] >= 0 && inputPoint[index] < 50 {
			zero += 0.0
		}

	}
	averge := (four + threeDotFive + three + thwoDotFive + two + oneDotFive + one + zeroDotFive + zero) / float64(len(inputPoint))
	return averge
}
