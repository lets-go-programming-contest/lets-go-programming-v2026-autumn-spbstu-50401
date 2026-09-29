package runner

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/Spider-has/task-1/pkg/calculator"
)

func Run(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	first, err := readOperand(scanner, 1)
	if err != nil {
		return err
	}

	second, err := readOperand(scanner, 2)
	if err != nil {
		return err
	}

	operator, err := readOperator(scanner)
	if err != nil {
		return err
	}

	result, err := calculator.Calculate(first, second, operator)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(output, result)
	return err
}

func readOperand(scanner *bufio.Scanner, pos int) (int, error) {
	if !scanner.Scan() {
		return 0, NewOperandError(pos, "")
	}

	str := scanner.Text()
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, NewOperandError(pos, str)
	}
	return num, nil
}

func readOperator(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", calculator.ErrInvalidOperator
	}

	str := scanner.Text()
	return str, nil
}
