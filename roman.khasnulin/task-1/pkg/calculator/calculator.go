package calculator

type sentinelError string

func (s sentinelError) Error() string {
	return string(s)
}

const (
	ErrDivByZero       = sentinelError("Division by zero")
	ErrInvalidOperator = sentinelError("Invalid operation")
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Calculate[T Number](first T, second T, operand byte) (int, error) {
	return 0, nil
}
