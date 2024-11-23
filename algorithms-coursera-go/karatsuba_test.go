package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	x        string
	y        string
	expected string
}

func TestKaratsuba(t *testing.T) {
	x := "3141592653589793238462643383279502884197169399375105820974944592"
	y := "2718281828459045235360287471352662497757247093699959574966967627"
	expected := "8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184"

	testCases := []testCase{
		{x, y, expected},
		{"0", "0", "0"},
		{"1", "1", "1"},
		{"1", "0", "0"},
		{"0", "1", "0"},
		{"10", "10", "100"},
		{"10", "1", "10"},
		{"1", "10", "10"},
	}

	for _, tc := range testCases {
		result := karatsubaMultiplication(tc.x, tc.y)
		if result != tc.expected {
			t.Errorf("Case: %s, Expected %s, got %s", tc, tc.expected, result)
		} else {
			fmt.Println("Case succeeded:", tc, "Result:", result)
		}
	}
}
