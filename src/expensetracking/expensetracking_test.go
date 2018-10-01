package expensetracking

import (
	"io/ioutil"
	"testing"
)

func Test_ExpenseAVGpreDay_Input_7_Days_Should_Be_Total_Price_185dot71(t *testing.T) {
	expected := 185.71
	file, _ := ioutil.ReadFile("dataex.txt")
	inputString := string(file)
	actual := ExpenseAVGpreDay(inputString)

	if expected != actual {
		t.Errorf("expect %2.f but it got %2.f", expected, actual)
	}
}

func Test_ExpenseAVGpreDay_Input_31_Days_Should_Be_Total_Price_211dot87(t *testing.T) {
	expected := 211.87
	file, _ := ioutil.ReadFile("data.txt")
	inputString := string(file)
	actual := ExpenseAVGpreDay(inputString)

	if expected != actual {
		t.Errorf("expect %2.f but it got %2.f", expected, actual)
	}
}
