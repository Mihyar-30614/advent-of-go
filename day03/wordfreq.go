package wordfreq

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

var ErrInvalidN = errors.New("invalid n")
var re = regexp.MustCompile(`[a-z0-9]+`)

type WordCount struct {
	Word  string
	Count int
}

func Tokenize(text string) []string {
	lowercaseText := strings.ToLower(text)
	words := re.FindAllString(lowercaseText, -1)

	if words == nil {
		return []string{}
	}

	return words
}

func CountWords(words []string) map[string]int {
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}
	return freq
}

func TopN(counts map[string]int, n int) ([]WordCount, error) {
	if n < 0 {
		return nil, ErrInvalidN
	}

	// var list []WordCount
	list := make([]WordCount, 0, len(counts))
	for w, c := range counts {
		list = append(list, WordCount{Word: w, Count: c})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].Count == list[j].Count {
			return list[i].Word < list[j].Word
		}

		return list[i].Count > list[j].Count
	})

	if n > len(list) {
		n = len(list)
	}

	return list[:n], nil
}
