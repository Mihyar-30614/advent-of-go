# Day 04 – Password Policy Parser

## Problem Description

You are given a multi-line text input. Each line describes a password policy and a password.
Your job is to parse the input and determine how many passwords are valid under **two different rules**.

This challenge focuses on:

- Parsing structured text
- Using structs
- Error handling
- Writing clear, testable logic

---

## Package

All code for this challenge must belong to the following package:

```go
package policy
```

---

## Error Handling

You must define the following errors:

```go
var (
    ErrInvalidLine   = errors.New("invalid line")
    ErrInvalidNumber = errors.New("invalid number")
)
```

Return:

- `ErrInvalidLine` when a line does not match the expected format.
- `ErrInvalidNumber` when numeric fields cannot be parsed.

---

## Input Format

Each non-empty line has the form:

```
min-max letter: password
```

Where:

- `min` and `max` are integers (base 10)
- `letter` is a single ASCII letter (`a`–`z` or `A`–`Z`)
- `password` is a non-empty string (may contain letters/digits; treat as raw text)

Examples of valid lines:

- `1-3 a: abcde`
- `1-3 b: cdefg`
- `2-9 c: ccccccccc`

Empty lines should be **ignored** (skipped) during parsing.

---

## Data Type

Define the following struct:

```go
type Entry struct {
    Min      int
    Max      int
    Letter   byte
    Password string
}
```

Notes:

- Use `byte` for `Letter` (the ASCII character).
- Preserve `Password` exactly as provided.

---

## Function 1: Parse

### Signature

```go
func Parse(input string) ([]Entry, error)
```

### What this function must do

- Split `input` into lines.
- Ignore empty lines.
- For each non-empty line:
  1. Parse `min` and `max` from the `min-max` part
  2. Parse `letter` from the `letter:` part (must be exactly 1 character before `:`)
  3. Parse `password` after the space following `letter:`
- If any line cannot be parsed, return:
  - `nil` and `ErrInvalidLine`, OR
  - `nil` and `ErrInvalidNumber` (if numbers fail to parse)
- Otherwise return the slice of parsed `Entry`.

---

## Rule A: Count-Based Validity

A password is valid if the policy letter appears in the password **at least `Min` times** and **at most `Max` times**.

---

## Function 2: ValidCountRule

### Signature

```go
func ValidCountRule(e Entry) bool
```

### What this function must do

- Count how many times `e.Letter` appears in `e.Password`
- Return `true` if `Min <= count <= Max`, else return `false`

---

## Rule B: Position-Based Validity

A password is valid if the policy letter appears **in exactly one** of the two positions:

- position `Min`
- position `Max`

Positions are **1-based** (the first character is position 1).

This is an XOR rule: one must match, but not both.

If either position is out of range for the password length, treat it as “does not match”.

---

## Function 3: ValidPositionRule

### Signature

```go
func ValidPositionRule(e Entry) bool
```

### What this function must do

- Let `p1 = e.Min` and `p2 = e.Max`
- Check whether `Password[p1-1] == Letter` (if `p1` is in range)
- Check whether `Password[p2-1] == Letter` (if `p2` is in range)
- Return `true` if exactly one check matches, else `false`

---

## Function 4: CountValid

### Signature

```go
func CountValid(entries []Entry, rule func(Entry) bool) int
```

### What this function must do

- Apply `rule` to each entry
- Return how many entries return `true`

---

## Notes

- You may only use the Go standard library.
- Your implementation must pass all tests in `policy_test.go`.

---

## Running Tests

From the `day04` directory:

```bash
go test -v
```

