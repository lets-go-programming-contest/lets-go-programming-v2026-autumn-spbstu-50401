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

func Calculate[T Number](first, second T, operator string) (T, error) {
	switch operator {
	case "+":
		return sum(first, second), nil
	case "-":
		return subtract(first, second), nil
	case "*":
		return multiply(first, second), nil
	case "/":
		if second == 0 {
			return 0, ErrDivByZero
		}
		return divide(first, second), nil
	default:
		return 0, ErrInvalidOperator
	}
}

func sum[T Number](first T, second T) T {
	return first + second
}

func subtract[T Number](first T, second T) T {
	return first - second
}

func multiply[T Number](first T, second T) T {
	return first * second
}

func divide[T Number](first T, second T) T {
	return first / second
}
