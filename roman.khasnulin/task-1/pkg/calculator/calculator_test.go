package calculator

import (
	"errors"
	"math"
	"testing"
)

type testCase[T Number] struct {
	name        string
	n1, n2      T
	op          string
	expected    T
	expectedErr error
}

func TestCalculateInt(t *testing.T) {
	testCases := []testCase[int]{
		{name: "sum positive test", n1: 1, n2: 2, op: "+", expected: 3, expectedErr: nil},
		{name: "subtract positive test", n1: 3, n2: 2, op: "-", expected: 1, expectedErr: nil},
		{name: "multiple positive test", n1: 5, n2: 6, op: "*", expected: 30, expectedErr: nil},
		{name: "divide positive test", n1: 11, n2: 5, op: "/", expected: 2, expectedErr: nil},
		{name: "divide negatives positive test", n1: -7, n2: 2, op: "/", expected: -3, expectedErr: nil},
		{name: "substract negatives positive test", n1: -5, n2: -6, op: "-", expected: 1, expectedErr: nil},
		{name: "zero divide positive test", n1: 0, n2: 5, op: "/", expected: 0, expectedErr: nil},
		{name: "zero divide negative test", n1: 11, n2: 0, op: "/", expected: 0, expectedErr: ErrDivByZero},
		{name: "wrong operator negative test", n1: 1, n2: 2, op: "%", expected: 0, expectedErr: ErrInvalidOperator},
		{name: "wrong long operator negative test", n1: 1, n2: 2, op: "%++//sincos", expected: 0, expectedErr: ErrInvalidOperator},
		{name: "wrong empty operator negative test", n1: 1, n2: 2, op: "", expected: 0, expectedErr: ErrInvalidOperator},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Calculate(tc.n1, tc.n2, tc.op)
			if !errors.Is(err, tc.expectedErr) {
				t.Fatalf("Expect error: %v, got error: %v", tc.expectedErr, err)
			}
			if tc.expectedErr != nil {
				return
			}
			if got != tc.expected {
				t.Errorf("Expected result: %d, got: %d", tc.expected, got)
			}
		})
	}
}

func TestCalculateFloat64(t *testing.T) {
	testCases := []testCase[float64]{
		{name: "sum positive test", n1: 1.0, n2: 2.0, op: "+", expected: 3.0, expectedErr: nil},
		{name: "subtract positive test", n1: 5.75, n2: 3.25, op: "-", expected: 2.5, expectedErr: nil},
		{name: "multiple positive test", n1: 1.25, n2: 4, op: "*", expected: 5, expectedErr: nil},
		{name: "divide positive test", n1: 13.0, n2: 4.0, op: "/", expected: 3.25, expectedErr: nil},
		{name: "sum small nums positive test", n1: 0.1, n2: 0.2, op: "+", expected: 0.3, expectedErr: nil},
		{name: "divide negative test", n1: 5, n2: 0, op: "/", expected: 0, expectedErr: ErrDivByZero},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Calculate(tc.n1, tc.n2, tc.op)
			if !errors.Is(err, tc.expectedErr) {
				t.Fatalf("Expect error: %v, got error: %v", tc.expectedErr, err)
			}
			if tc.expectedErr != nil {
				return
			}
			if math.Abs(got-tc.expected) > 1e-9 {
				t.Errorf("Expected result: %v, got: %v", tc.expected, got)
			}
		})
	}
}
