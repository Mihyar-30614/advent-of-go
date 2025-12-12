# Day 03 – Word Frequency

## Problem Description

You are given a block of text. Your task is to analyze this text and determine
which words appear most frequently.

This challenge focuses on:

- String processing
- Maps
- Sorting
- Deterministic output rules

---

## Package

All code for this challenge must belong to the following package:

```go
package wordfreq
```

---

## Error Handling

You must define the following error:

```go
var ErrInvalidN = errors.New("invalid n")
```

This error must be returned when the value `n` passed to `TopN` is negative.

---

## Definitions

### What is a “word”?

A **word** is defined as a contiguous sequence of:

- Letters (`A–Z`, `a–z`)
- Digits (`0–9`)

All other characters are considered **separators**.

Examples:

- `"Hello, world!"` → `["hello", "world"]`
- `"Go-go-go"` → `["go", "go", "go"]`
- `"go_lang"` → `["go", "lang"]`
- `"123!!abc"` → `["123", "abc"]`
- `"..."` → `[]`

---

### Normalization Rules

- All words must be converted to **lowercase**
- Empty words must be ignored
- Word order must be preserved during tokenization

---

## Data Type

You must define the following type:

```go
type WordCount struct {
    Word  string
    Count int
}
```

---

## Function 1: Tokenize

### Signature

```go
func Tokenize(text string) []string
```

### What this function must do

- Convert the input string to lowercase
- Split the string into words using the definition above
- Return the words in the order they appear
- If no words exist, return an empty slice

### Examples

| Input                       | Output                             |
|----------------------------|------------------------------------|
| `""`                       | `[]`                               |
| `"Hello, world!"`          | `["hello","world"]`                |
| `"Go-go-go"`               | `["go","go","go"]`                 |
| `"  123!!abc "`            | `["123","abc"]`                    |
| `"Mix3d CASE and NUM83R5"` | `["mix3d","case","and","num83r5"]` |

---

## Function 2: CountWords

### Signature

```go
func CountWords(words []string) map[string]int
```

### What this function must do

- Count how many times each word appears in the input slice
- Return a map where:
  - key = word
  - value = count
- If the input slice is empty, return an empty map

### Examples

Input:

```go
[]string{"go", "go", "lang"}
```

Output:

```go
map[string]int{
    "go":   2,
    "lang": 1,
}
```

---

## Function 3: TopN

### Signature

```go
func TopN(counts map[string]int, n int) ([]WordCount, error)
```

### What this function must do

- If `n < 0`, return `nil` and `ErrInvalidN`
- If `n == 0`, return an empty slice and `nil`
- Sort words using the following rules:
  1. Higher `Count` comes first
  2. If counts are equal, sort alphabetically by `Word` (ascending)
- Return the top `n` results
- If `n` is greater than the number of unique words, return all results

### Examples

Given:

```go
counts := map[string]int{
    "go":   3,
    "lang": 2,
    "is":   2,
    "fun":  1,
}
```

| n  | Output |
|----|--------|
| 0  | `[]` |
| 2  | `[{"go",3},{"is",2}]` |
| 10 | `[{"go",3},{"is",2},{"lang",2},{"fun",1}]` |

---

## Notes

- Only the Go standard library may be used
- Sorting is required for `TopN`
- Output must be deterministic
- Your implementation must pass all tests in `wordfreq_test.go`

---

## Running Tests

From the `day03` directory:

```bash
go test -v
```

