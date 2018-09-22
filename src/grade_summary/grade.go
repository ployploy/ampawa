package grade_summary

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Grade struct {
	APlus int
	A     int
	BPlus int
	B     int
	CPlus int
	C     int
	DPlus int
	D     int
	F     int
}

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

func GradeSummary(inputPoint []int) Grade {
	var countAPlus, countA, countBPlus, countB, countCPlus, countC, countDPlus, countD, countF int
	for index, _ := range inputPoint {
		if inputPoint[index] >= 85 && inputPoint[index] <= 100 {
			countAPlus++
		}
		if inputPoint[index] >= 80 && inputPoint[index] < 85 {
			countA++
		}
		if inputPoint[index] >= 75 && inputPoint[index] < 80 {
			countBPlus++
		}
		if inputPoint[index] >= 70 && inputPoint[index] < 75 {
			countB++
		}
		if inputPoint[index] >= 65 && inputPoint[index] < 70 {
			countCPlus++
		}
		if inputPoint[index] >= 60 && inputPoint[index] < 65 {
			countC++
		}
		if inputPoint[index] >= 55 && inputPoint[index] < 60 {
			countDPlus++
		}
		if inputPoint[index] >= 50 && inputPoint[index] < 55 {
			countD++
		}
		if inputPoint[index] >= 0 && inputPoint[index] < 50 {
			countF++
		}

	}
	return Grade{
		APlus: countAPlus,
		A:     countA,
		BPlus: countBPlus,
		B:     countB,
		CPlus: countCPlus,
		C:     countC,
		DPlus: countDPlus,
		D:     countD,
		F:     countF,
	}
}
