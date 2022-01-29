package utils

import "testing"

/*
func Test_IsPrime(t *testing.T) {
	//Arrange
	no := int32(61)
	expected := true

	//Act
	actual := IsPrime(no)

	//Assert
	if actual != expected {
		t.Errorf("Checking %d is Prime, exepected %v, got %v", no, expected, actual)
	}
}
*/

type testCase struct {
	name     string
	input    int32
	expected bool
	actual   bool
}

func Test_IsPrime(t *testing.T) {
	testCases := []testCase{
		testCase{name: "Testing_Prime:31", input: -1, expected: false},
		testCase{name: "Testing_Prime:31", input: 2, expected: true},
		testCase{name: "Testing_Prime:31", input: 31, expected: true},
		testCase{name: "Testing_Prime:61", input: 61, expected: true},
		testCase{name: "Testing_Prime:71", input: 71, expected: true},
		testCase{name: "Testing_Prime:73", input: 73, expected: true},
		testCase{name: "Testing_Prime:64", input: 64, expected: false},
		testCase{name: "Testing_Prime:83", input: 83, expected: true},
		testCase{name: "Testing_Prime:12", input: 12, expected: false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//Arrange
			no := tc.input
			tc.actual = IsPrime(no)

			//Assert
			if tc.actual != tc.expected {
				t.Errorf("Checking (odd) %d is Prime, exepected %v, got %v", tc.input, tc.expected, tc.actual)
			}

		})
	}
}
