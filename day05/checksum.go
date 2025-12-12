package checksum

import "errors"

var (
	ErrInvalidLine     = errors.New("invalid line")
	ErrInvalidNumber   = errors.New("invalid number")
	ErrNoDivisiblePair = errors.New("no divisible pair")
)

func Parse(input string) ([][]int, error) {
	// TODO: implement
	return nil, nil
}

func RangeChecksum(sheet [][]int) int {
	// TODO: implement
	return 0
}

func DivisibleChecksum(sheet [][]int) (int, error) {
	// TODO: implement
	return 0, nil
}
