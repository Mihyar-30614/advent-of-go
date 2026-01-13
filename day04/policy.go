package policy

import (
	"errors"
	"strconv"
	"strings"
)

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
	var entries []Entry

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			return nil, ErrInvalidLine
		}

		rangeParts := strings.Split(parts[0], "-")
		if len(rangeParts) != 2 {
			return nil, ErrInvalidNumber
		}
		min, err1 := strconv.Atoi(rangeParts[0])
		max, err2 := strconv.Atoi(rangeParts[1])
		if err1 != nil || err2 != nil {
			return nil, ErrInvalidNumber
		}

		if len(parts[1]) != 2 || parts[1][1] != ':' {
			return nil, ErrInvalidLine
		}

		letter := parts[1][0]
		password := parts[2]

		entries = append(entries, Entry{
			Min:      min,
			Max:      max,
			Letter:   letter,
			Password: password,
		})
	}

	return entries, nil
}

func ValidCountRule(e Entry) bool {
	count := strings.Count(e.Password, string(e.Letter))
	return count >= e.Min && count <= e.Max
}

func ValidPositionRule(e Entry) bool {
	// convert 1-based positions to 0-based positions
	idx1 := e.Min - 1
	idx2 := e.Max - 1

	// Check if positions matches and within bounds
	match1 := idx1 >= 0 && idx1 <= len(e.Password) && e.Password[idx1] == e.Letter
	match2 := idx2 >= 0 && idx2 <= len(e.Password) && e.Password[idx2] == e.Letter

	// XOR returns true if exactly one matches
	return match1 != match2
}

func CountValid(entries []Entry, rule func(Entry) bool) int {
	count := 0
	for _, e := range entries {
		if rule(e) {
			count++
		}
	}
	return count
}
