package must

import (
	"errors"
	"testing"
)

func TestMust(t *testing.T) {
	exampleFunc := func() (int, error) {
		return 42, nil
	}

	errorFunc := func() (int, error) {
		return 42, errors.New("an error occurred")
	}

	funcNoError := func() int {
		return 42
	}

	funcPanic := func() int {
		panic("42")
	}

	errorFuncInt := func(x int) (int, error) {
		return 42 * x, errors.New("another error occurred")
	}

	multiply := func(x, y float64) (float64, error) {
		return x * y, errors.New("yet another error occurred")
	}

	type testStruct struct {
		A int
		B string
	}

	incrementStruct := func(a testStruct) (testStruct, error) {
		a.A++
		a.B += "!"
		return a, errors.New("some error!")
	}

	if Must(exampleFunc) != 42 {
		t.Error("Must(exampleFunc) should return 42")
	}
	if Must(errorFunc) != 42 {
		t.Error("Must(errorFunc) should return 42")
	}
	if Must(funcNoError) != 42 {
		t.Error("Must(funcNoError) should return 42")
	}
	if Must(errorFuncInt, 2) != 84 {
		t.Error("Must(errorFunc) should return 84")
	}
	if Must(multiply, 21.0, 2.0) != 42.0 {
		t.Error("Must(errorFunc) should return 42.0")
	}
	if Must(funcPanic) != nil {
		t.Error("Must(funcPanic) should return nil")
	}
	// structs do need to be cast
	testStruct2 := Must(incrementStruct, testStruct{A: 41, B: "Hello"}).(testStruct)
	if testStruct2.A != 42 {
		t.Error("Must(incrementStruct) should return 42")
	}
}
