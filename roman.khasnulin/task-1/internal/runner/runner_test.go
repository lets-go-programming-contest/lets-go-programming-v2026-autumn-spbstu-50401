package runner

import (
	"errors"
	"strings"
	"testing"

	"github.com/Spider-has/task-1/pkg/calculator"
)

type ioRunTest struct {
	name          string
	input         string
	expectedOut   string
	expectedError error
}

func TestRun(t *testing.T) {
	testCases := []ioRunTest{
		{name: "multiply positive test", input: "10\n5\n*\n", expectedOut: "50\n", expectedError: nil},
		{name: "sum positive test", input: "10\n5\n+\n", expectedOut: "15\n", expectedError: nil},
		{name: "subtract positive test", input: "10\n5\n-\n", expectedOut: "5\n", expectedError: nil},
		{name: "divide positive test", input: "10\n5\n/\n", expectedOut: "2\n", expectedError: nil},
		{name: "invalid first operand negative test", input: "ABC\n5\n*", expectedOut: "", expectedError: NewOperandError(1, "ABC")},
		{name: "invalid second operand negative test", input: "10\nABC\n*", expectedOut: "", expectedError: NewOperandError(2, "ABC")},
		{name: "invalid operator negative test", input: "10\n5\n%", expectedOut: "", expectedError: calculator.ErrInvalidOperator},
		{name: "divide by zero error negative test", input: "10\n0\n/", expectedOut: "", expectedError: calculator.ErrDivByZero},
		{name: "empty input negative test", input: "", expectedOut: "", expectedError: NewOperandError(1, "")},
		{name: "only first operand input negative test", input: "10\n", expectedOut: "", expectedError: NewOperandError(2, "")},
		{name: "only first and second operand input negative test", input: "10\n5\n", expectedOut: "", expectedError: calculator.ErrInvalidOperator},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			input := strings.NewReader(testCase.input)
			output := strings.Builder{}
			err := Run(input, &output)
			if !errors.Is(err, testCase.expectedError) {
				t.Fatalf("Expect error: %v, got error: %v", testCase.expectedError, err)
			}

			if output.String() != testCase.expectedOut {
				t.Errorf("Expected result: %s, got: %s", testCase.expectedOut, output.String())
			}
		})
	}
}
