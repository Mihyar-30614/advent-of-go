# Day 02 – Floor Instructions

## Problem Description

You are in a building and receive a string of instructions that move you between floors.

Each character in the instruction string represents a movement:

- `'('` means move **up** one floor
- `')'` means move **down** one floor

You always start on **floor 0**.

---

## Package

All code for this challenge must belong to the following package:

```go
package floors
```

---

## Error Handling

You must define the following error:

```go
var ErrInvalidInstruction = errors.New("invalid instruction")
```

This error must be returned whenever the instruction string contains a character
other than `'('` or `')'`.

Invalid characters must **not** be ignored.

---

## Function 1: `FinalFloor`

### Signature

```go
func FinalFloor(instructions string) (int, error)
```

### What this function must do

- Start at floor `0`.
- Read the `instructions` string from left to right.
- For each character:
  - If the character is `'('`, increase the floor by `1`.
  - If the character is `')'`, decrease the floor by `1`.
  - If the character is anything else:
    - Stop processing
    - Return `0` and `ErrInvalidInstruction`
- If the instruction string is empty:
  - Return `0` and `nil`
- If all characters are valid:
  - Return the final floor and `nil`

### Examples

| Input       | Output |
|------------|--------|
| `""`       | `0`    |
| `"("`      | `1`    |
| `")"`      | `-1`   |
| `"()()"`   | `0`    |
| `"(()())"` | `0`    |
| `"(a)"`    | error  |

---

## Function 2: `FirstBasementPosition`

### Signature

```go
func FirstBasementPosition(instructions string) (int, error)
```

### What this function must do

- Start at floor `0`.
- Process the instruction string **one character at a time**.
- Track the position of each character using **1-based indexing**:
  - The first character has position `1`.
- For each character:
  - If the character is `'('`, increase the floor by `1`.
  - If the character is `')'`, decrease the floor by `1`.
  - If the character is anything else:
    - Stop processing
    - Return `0` and `ErrInvalidInstruction`
- Immediately after applying each instruction:
  - If the floor becomes `-1`, return the current position.
- If the end of the string is reached and the basement was never entered:
  - Return `0` and `nil`
- If the instruction string is empty:
  - Return `0` and `nil`

### Examples

| Input     | Explanation                              | Output |
|----------|------------------------------------------|--------|
| `")"`    | First move reaches basement               | `1`    |
| `"()())"`| Basement reached at 5th instruction       | `5`    |
| `"((("`  | Basement is never reached                | `0`    |
| `"())x"` | Invalid instruction encountered           | error  |

---

## Notes

- Positions are **1-based**, not zero-based.
- Prefer early returns when an error occurs.
- Your implementation must pass all tests in `floors_test.go`.

---

## Running Tests

From the `day02` directory:

```bash
go test
```
