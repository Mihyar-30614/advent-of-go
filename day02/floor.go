package floors

import "errors"

var ErrInvalidInstruction = errors.New("invalid instruction")

func FinalFloor(instructions string) (int, error) {
	if len(instructions) == 0 {
		return 0, nil
	}

	finalFloor := 0
	for _, char := range instructions {

		switch char {
		case '(':
			finalFloor++
		case ')':
			finalFloor--
		default:
			return 0, ErrInvalidInstruction
		}

	}
	return finalFloor, nil
}

func FirstBasementPosition(instructions string) (int, error) {
	// pass 1: validate
	for _, ch := range instructions {
		if ch != '(' && ch != ')' {
			return 0, ErrInvalidInstruction
		}
	}

	// pass 2: compute
	floor := 0
	for i, ch := range instructions {
		if ch == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			return i + 1, nil
		}
	}
	return 0, nil
}
