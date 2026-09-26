package calculator

type sentinelError string

func (s sentinelError) Error() string {
	return string(s)
}

const (
	ErrDivByZero       = sentinelError("Division by zero")
	ErrInvalidOperator = sentinelError("Invalid operation")
)
