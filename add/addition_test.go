package add

import "testing"

func Test_Addition_Input1_48_And_Input2_16_Should_Be_64(t *testing.T) {
	expected := "64"
	input1 := []int{4, 8}
	input2 := []int{1, 6}

	actualResult := Addition(input1, input2)

	if expected != actualResult {
		t.Errorf("expect %s but got it %s", expected, actualResult)
	}

}

func Test_Addition_Input1_921_And_Input2_251_Should_Be_1172(t *testing.T) {
	expected := "1172"
	input1 := []int{9, 2, 1}
	input2 := []int{2, 5, 1}

	actualResult := Addition(input1, input2)

	if expected != actualResult {
		t.Errorf("expect %s but got it %s", expected, actualResult)
	}
}

func Test_Addition_Input1_78_And_Input2_22_Should_Be_100(t *testing.T) {
	expected := "100"
	input1 := []int{7, 8}
	input2 := []int{2, 2}

	actualResult := Addition(input1, input2)

	if expected != actualResult {
		t.Errorf("expect %s but got it %s", expected, actualResult)
	}
}

func Test_Addition_Input1_893880539957604781622135818781054154760092_And_Input2_458875111557430000227050141224806041518870_Should_Be_XX(t *testing.T) {
	expected := "1352755651515034781849185960005860196278962"
	input1 := []int{8, 9, 3, 8, 8, 0, 5, 3, 9, 9, 5, 7, 6, 0, 4, 7, 8, 1, 6, 2, 2, 1, 3, 5, 8, 1, 8, 7, 8, 1, 0, 5, 4, 1, 5, 4, 7, 6, 0, 0, 9, 2}
	input2 := []int{4, 5, 8, 8, 7, 5, 1, 1, 1, 5, 5, 7, 4, 3, 0, 0, 0, 0, 2, 2, 7, 0, 5, 0, 1, 4, 1, 2, 2, 4, 8, 0, 6, 0, 4, 1, 5, 1, 8, 8, 7, 0}

	actualResult := Addition(input1, input2)

	if expected != actualResult {
		t.Errorf("expect %s but got it %s", expected, actualResult)
	}
}
