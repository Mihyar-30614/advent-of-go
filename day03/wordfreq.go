package wordfreq

import "errors"

var ErrInvalidN = errors.New("invalid n")

type WordCount struct {
	Word  string
	Count int
}

func Tokenize(text string) []string {
	// TODO: implement
	return nil
}

func CountWords(words []string) map[string]int {
	// TODO: implement
	return nil
}

func TopN(counts map[string]int, n int) ([]WordCount, error) {
	// TODO: implement
	return nil, nil
}
