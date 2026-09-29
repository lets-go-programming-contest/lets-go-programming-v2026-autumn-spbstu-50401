package runner

type OperandError struct {
	pos  int
	read string
}

func NewOperandError(pos int, read string) *OperandError {
	return &OperandError{
		pos:  pos,
		read: read,
	}
}

func (op *OperandError) Error() string {
	switch op.pos {
	case 1:
		return "Invalid first operand"
	case 2:
		return "Invalid second operand"
	default:
		return "Invalide operand"
	}
}
