package grade_average

import "testing"

func Test_GradeAverage_InputInput_25_4_Should_Be_0(t *testing.T) {
	expected := 0.0
	inputPoint := []int{25, 4}

	actualResult := GradeAverage(inputPoint)

	if expected != actualResult {
		t.Errorf("expect %.3f but got it %.3f", expected, actualResult)
	}
}

func Test_GradeAverage_InputInput_ListPoints_Should_Be_1_Dot_375(t *testing.T) {
	expected := 1.375
	inputPoint := Readfile("inputdata.txt")

	actualResult := GradeAverage(inputPoint)

	if expected != actualResult {
		t.Errorf("expect %.3f but got it %.3f", expected, actualResult)
	}
}
