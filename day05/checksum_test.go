package checksum

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	input := "5 1 9 5\n7\t5\t3\n\n2 4 6 8\n"
	got, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse unexpected error: %v", err)
	}
	want := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3},
		{2, 4, 6, 8},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Parse = %#v, want %#v", got, want)
	}
}

func TestParseInvalidNumber(t *testing.T) {
	_, err := Parse("1 2 x\n")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != ErrInvalidNumber {
		t.Fatalf("expected ErrInvalidNumber, got %v", err)
	}
}

func TestParseInvalidLine(t *testing.T) {
	_, err := Parse("   \n")
	// whitespace-only lines are ignored, so this is OK:
	if err != nil {
		t.Fatalf("expected nil error for whitespace-only input, got %v", err)
	}

	_, err = Parse("5 6\n---\n7 8\n")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != ErrInvalidNumber && err != ErrInvalidLine {
		t.Fatalf("expected ErrInvalidLine or ErrInvalidNumber, got %v", err)
	}
}

func TestRangeChecksumExample(t *testing.T) {
	input := `5 1 9 5
7 5 3
2 4 6 8
`
	sheet, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse unexpected error: %v", err)
	}
	got := RangeChecksum(sheet)
	if got != 18 {
		t.Fatalf("RangeChecksum = %d, want 18", got)
	}
}

func TestDivisibleChecksumExample(t *testing.T) {
	input := `5 9 2 8
9 4 7 3
3 8 6 5
`
	sheet, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse unexpected error: %v", err)
	}
	got, err := DivisibleChecksum(sheet)
	if err != nil {
		t.Fatalf("DivisibleChecksum unexpected error: %v", err)
	}
	if got != 9 {
		t.Fatalf("DivisibleChecksum = %d, want 9", got)
	}
}

func TestDivisibleChecksumNoPair(t *testing.T) {
	sheet := [][]int{
		{5, 7, 11},
	}
	_, err := DivisibleChecksum(sheet)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != ErrNoDivisiblePair {
		t.Fatalf("expected ErrNoDivisiblePair, got %v", err)
	}
}

func TestDivisibleChecksumEmpty(t *testing.T) {
	got, err := DivisibleChecksum(nil)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}
