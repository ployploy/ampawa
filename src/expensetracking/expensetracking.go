package expensetracking

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ExpenseAVGpreDay(inputString string) float64 {
	var sum int
	expenses := strings.Split(inputString, "\n")

	for _, expense := range expenses {
		numbers := strings.Split(expense, " ")
		numbers = numbers[1:]
		fmt.Print(numbers)
		for _, number := range numbers {
			number = number[1:]
			numberInt, _ := strconv.Atoi(number)
			sum += numberInt
		}
	}

	days := len(expenses)
	totalNumber := float64(sum) / float64(days)
	return math.Round(totalNumber*100) / 100
}
