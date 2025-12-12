package policy

import (
	"reflect"
	"testing"
)

const sample = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`

func TestParseSample(t *testing.T) {
	got, err := Parse(sample)
	if err != nil {
		t.Fatalf("Parse(sample) unexpected error: %v", err)
	}

	want := []Entry{
		{Min: 1, Max: 3, Letter: 'a', Password: "abcde"},
		{Min: 1, Max: 3, Letter: 'b', Password: "cdefg"},
		{Min: 2, Max: 9, Letter: 'c', Password: "ccccccccc"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Parse(sample) = %#v, want %#v", got, want)
	}
}

func TestParseIgnoresEmptyLines(t *testing.T) {
	input := "\n\n1-1 z: z\n\n"
	got, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse unexpected error: %v", err)
	}
	want := []Entry{{Min: 1, Max: 1, Letter: 'z', Password: "z"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Parse = %#v, want %#v", got, want)
	}
}

func TestParseInvalidLine(t *testing.T) {
	_, err := Parse("not-a-line")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != ErrInvalidLine && err != ErrInvalidNumber {
		t.Fatalf("expected ErrInvalidLine or ErrInvalidNumber, got %v", err)
	}
}

func TestValidCountRule(t *testing.T) {
	entries, _ := Parse(sample)

	if !ValidCountRule(entries[0]) {
		t.Fatalf("expected first entry valid under count rule")
	}
	if ValidCountRule(entries[1]) {
		t.Fatalf("expected second entry invalid under count rule")
	}
	if !ValidCountRule(entries[2]) {
		t.Fatalf("expected third entry valid under count rule")
	}
}

func TestValidPositionRule(t *testing.T) {
	entries, _ := Parse(sample)

	if !ValidPositionRule(entries[0]) {
		t.Fatalf("expected first entry valid under position rule")
	}
	if ValidPositionRule(entries[1]) {
		t.Fatalf("expected second entry invalid under position rule")
	}
	if ValidPositionRule(entries[2]) {
		t.Fatalf("expected third entry invalid under position rule (both positions match)")
	}
}

func TestCountValid(t *testing.T) {
	entries, _ := Parse(sample)

	gotA := CountValid(entries, ValidCountRule)
	if gotA != 2 {
		t.Fatalf("CountValid count-rule = %d, want 2", gotA)
	}

	gotB := CountValid(entries, ValidPositionRule)
	if gotB != 1 {
		t.Fatalf("CountValid position-rule = %d, want 1", gotB)
	}
}

func TestValidPositionRuleOutOfRange(t *testing.T) {
	e := Entry{Min: 10, Max: 20, Letter: 'a', Password: "abc"}
	if ValidPositionRule(e) {
		t.Fatalf("expected out-of-range positions to be non-matching")
	}
}
