package checksum

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidLine     = errors.New("invalid line")
	ErrInvalidNumber   = errors.New("invalid number")
	ErrNoDivisiblePair = errors.New("no divisible pair")
)

func Parse(input string) ([][]int, error) {

	lines := strings.Split(input, "\n")
	out := make([][]int, 0, len(lines))

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			// empty lines are ignored
			continue
		}

		fields := strings.Fields(trimmed)
		if len(fields) == 0 {
			return nil, ErrInvalidLine
		}

		row := make([]int, 0, len(fields))
		for _, tok := range fields {
			n, err := strconv.Atoi(tok)
			if err != nil {
				return nil, ErrInvalidNumber
			}
			row = append(row, n)
		}

		out = append(out, row)
	}
	return out, nil
}

func RangeChecksum(sheet [][]int) int {
	total := 0

	for _, row := range sheet {
		min, max := row[0], row[0]

		for _, v := range row[1:] {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
		total += max - min
	}
	return total
}

func DivisibleChecksum(sheet [][]int) (int, error) {
	total := 0

	for _, row := range sheet {
		found := false
		rowVal := 0

		for i := 0; i < len(row) && !found; i++ {
			for j := 1; j < len(row); j++ {
				a, b := row[i], row[j]

				if a == b {
					continue
				}

				if b != 0 && a%b == 0 {
					rowVal = a / b
					found = true
					break
				}

				if a != 0 && b%a == 0 {
					rowVal = b / a
					found = true
					break
				}

			}
		}
		if !found {
			return 0, ErrNoDivisiblePair
		}
		total += rowVal
	}
	return total, nil
}
