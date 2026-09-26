package calculator

import (
	"errors"
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
		{"sum positive test", 1, 2, "+", 3, nil},
		{"subtract positive test", 3, 2, "-", 1, nil},
		{"multiple positive test", 5, 6, "*", 30, nil},
		{"divide positive test", 11, 5, "/", 2, nil},
		{"zero divide positive test", 0, 5, "/", 0, nil},
		{"zero divide negative test", 11, 0, "/", 0, ErrDivByZero},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Calculate(tc.n1, tc.n2, tc.op)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expect error: %v, got error: %v", tc.expectedErr, err)
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
