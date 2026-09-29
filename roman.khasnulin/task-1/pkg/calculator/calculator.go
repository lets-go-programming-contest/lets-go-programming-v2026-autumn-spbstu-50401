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

type binaryOp[T Number] func(first T, second T) (T, error)

func Calculate[T Number](first, second T, operator string) (T, error) {
	operate, err := selectOperation[T](operator)
	if err != nil {
		return 0, err
	}
	return operate(first, second)
}

func selectOperation[T Number](operator string) (binaryOp[T], error) {
	switch operator {
	case "+":
		return add, nil
	case "-":
		return subtract, nil
	case "*":
		return multiply, nil
	case "/":
		return divide, nil
	default:
		return nil, ErrInvalidOperator
	}
}

func add[T Number](first, second T) (T, error) {
	return first + second, nil
}

func subtract[T Number](first, second T) (T, error) {
	return first - second, nil
}

func multiply[T Number](first, second T) (T, error) {
	return first * second, nil
}

func divide[T Number](first, second T) (T, error) {
	if second == 0 {
		return 0, ErrDivByZero
	}
	return first / second, nil
}
