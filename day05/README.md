# Day 05 – Spreadsheet Checksums

## Problem Description

You are given a multi-line “spreadsheet” of integers. Each line represents a row, and each row contains one or more integers separated by whitespace (spaces and/or tabs).

Your job is to compute **two different checksums** from this spreadsheet.

This challenge focuses on:

- Parsing text input into structured data
- Nested loops
- Error handling
- Writing clean, testable Go functions

---

## Package

All code for this challenge must belong to the following package:

```go
package checksum
```

---

## Error Handling

You must define the following errors:

```go
var (
    ErrInvalidLine      = errors.New("invalid line")
    ErrInvalidNumber    = errors.New("invalid number")
    ErrNoDivisiblePair  = errors.New("no divisible pair")
)
```

Return:

- `ErrInvalidLine` when a non-empty line contains no numbers.
- `ErrInvalidNumber` when any value cannot be parsed as an integer.
- `ErrNoDivisiblePair` when a row has **no** pair where one number evenly divides the other (for the divisible checksum).

Empty lines should be **ignored** during parsing.

---

## Function 1: Parse

### Signature

```go
func Parse(input string) ([][]int, error)
```

### What this function must do

- Split `input` into lines.
- Ignore empty/whitespace-only lines.
- For each non-empty line:
  - Split the line by whitespace.
  - Parse each token as a base-10 integer.
  - If parsing fails, return `nil` and `ErrInvalidNumber`.
  - If the line contains no numbers (after splitting), return `nil` and `ErrInvalidLine`.
- Return the spreadsheet as a slice of rows (`[][]int`).

### Examples

Input:

```
5 1 9 5
7 5 3
2 4 6 8
```

Output:

```go
[][]int{
    {5, 1, 9, 5},
    {7, 5, 3},
    {2, 4, 6, 8},
}
```

---

## Checksum A: Range Checksum

For each row:

- Find the largest value (`max`)
- Find the smallest value (`min`)
- Row value = `max - min`

The **Range Checksum** is the sum of row values across all rows.

---

## Function 2: RangeChecksum

### Signature

```go
func RangeChecksum(sheet [][]int) int
```

### What this function must do

- For each row in `sheet`:
  - Compute `max(row) - min(row)`
  - Add it to a running total
- Return the total
- If `sheet` is empty, return `0`.
- You may assume each row is non-empty if it came from `Parse`.

### Examples

For:

```
5 1 9 5     -> max 9, min 1, diff 8
7 5 3       -> max 7, min 3, diff 4
2 4 6 8     -> max 8, min 2, diff 6
```

Range checksum = `8 + 4 + 6 = 18`

---

## Checksum B: Divisible Checksum

For each row, find the **only** pair of numbers where:

- One number evenly divides the other (remainder is 0)

Then:

- Row value = `larger / smaller`

The **Divisible Checksum** is the sum of row values across all rows.

Rules:

- Each row is guaranteed (by the problem) to contain **at most one** divisible pair.
- If a row contains **no** divisible pair, return an error.

---

## Function 3: DivisibleChecksum

### Signature

```go
func DivisibleChecksum(sheet [][]int) (int, error)
```

### What this function must do

- For each row:
  - Find a pair `(a, b)` with `a != b` where either `a % b == 0` or `b % a == 0`
  - Add `max(a,b)/min(a,b)` to a running total
  - If no such pair exists in a row, return `0` and `ErrNoDivisiblePair`
- Return the total and `nil` if all rows contain a divisible pair.
- If `sheet` is empty, return `0` and `nil`.

### Examples

For:

```
5 9 2 8     -> divisible pair (8,2) -> 8/2 = 4
9 4 7 3     -> divisible pair (9,3) -> 9/3 = 3
3 8 6 5     -> divisible pair (6,3) -> 6/3 = 2
```

Divisible checksum = `4 + 3 + 2 = 9`

---

## Notes

- You may only use the Go standard library.
- Keep your code simple and readable.
- Your implementation must pass all tests in `checksum_test.go`.

---

## Running Tests

From the `day05` directory:

```bash
go test -v
```

