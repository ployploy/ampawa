package grade_summary

import (
	"testing"
)

func Test_GradeSummary_Input_25_4_Should_Be_F_2(t *testing.T) {
	expected := Grade{
		APlus: 0,
		A:     0,
		BPlus: 0,
		B:     0,
		CPlus: 0,
		C:     0,
		DPlus: 0,
		D:     0,
		F:     2,
	}
	inputPoint := []int{25, 4}

	actualResult := GradeSummary(inputPoint)

	if expected != actualResult {
		t.Errorf("expect %d but got it %d", expected, actualResult)
	}
}

func Test_GradeSummary_Input_25_4_57_73_1_Should_Be_B_1_DPlus_1_F_3(t *testing.T) {
	expected := Grade{
		APlus: 0,
		A:     0,
		BPlus: 0,
		B:     1,
		CPlus: 0,
		C:     0,
		DPlus: 1,
		D:     0,
		F:     3,
	}
	inputPoint := []int{25, 4, 57, 73, 1}

	actualResult := GradeSummary(inputPoint)

	if expected != actualResult {
		t.Errorf("expect %d but got it %d", expected, actualResult)
	}
}

func Test_GradeSummary_Input_sum_Should_Be_B_1_DPlus_1_F_3(t *testing.T) {
	expected := Grade{
		APlus: 1,
		A:     0,
		BPlus: 0,
		B:     1,
		CPlus: 0,
		C:     1,
		DPlus: 1,
		D:     0,
		F:     1,
	}
	inputPoint := []int{
		30,
		56,
		72,
		89,
		60,
	}

	actualResult := GradeSummary(inputPoint)

	if expected != actualResult {
		t.Errorf("expect %d but got it %d", expected, actualResult)
	}
}
