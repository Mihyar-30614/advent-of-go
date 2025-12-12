package policy

import "errors"

var (
	ErrInvalidLine   = errors.New("invalid line")
	ErrInvalidNumber = errors.New("invalid number")
)

type Entry struct {
	Min      int
	Max      int
	Letter   byte
	Password string
}

func Parse(input string) ([]Entry, error) {
	// TODO: implement
	return nil, nil
}

func ValidCountRule(e Entry) bool {
	// TODO: implement
	return false
}

func ValidPositionRule(e Entry) bool {
	// TODO: implement
	return false
}

func CountValid(entries []Entry, rule func(Entry) bool) int {
	// TODO: implement
	return 0
}
