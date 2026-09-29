package runner

type OperandError struct {
	Pos  int
	Read string
}

func NewOperandError(pos int, read string) *OperandError {
	return &OperandError{
		Pos:  pos,
		Read: read,
	}
}

func (op *OperandError) Error() string {
	switch op.Pos {
	case 1:
		return "Invalid first operand"
	case 2:
		return "Invalid second operand"
	default:
		return "Invalid operand"
	}
}

func (op *OperandError) Is(target error) bool {
	t, ok := target.(*OperandError)
	if !ok {
		return false
	}
	return op.Pos == t.Pos && op.Read == t.Read
}
