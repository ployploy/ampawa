package sum

import (
	"testing"
)

func Test_Summation_Input_2_3_4_5_Should_Be_14(t *testing.T) {
	expected := 14
	input := []int{2, 3, 4, 5}

	actualResult := Summation(input)

	if expected != actualResult {
		t.Errorf("expect %d but got it %d", expected, actualResult)
	}
}

func Test_Summation_Input_28_35_42_Should_Be_105(t *testing.T) {
	expected := 105
	input := []int{28, 35, 42}

	actualResult := Summation(input)

	if expected != actualResult {
		t.Errorf("expect %d but got it %d", expected, actualResult)
	}
}

func Test_Summation_Input_sum_Should_Be_474872753(t *testing.T) {
	expected := 635589218
	input := Readfile("inputdata.txt")

	actualResult := Summation(input)

	if expected != actualResult {
		t.Errorf("expect %d but got it %d", expected, actualResult)
	}
}

func Test_Summation_Input_9388053_9957604_7816221_Should_Be_(t *testing.T) {
	expected := 27161878
	input := []int{9388053, 9957604, 7816221}

	actualResult := Summation(input)

	if expected != actualResult {
		t.Errorf("expect %d but got it %d", expected, actualResult)
	}

}
