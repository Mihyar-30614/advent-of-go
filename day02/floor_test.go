package floors

import "testing"

func TestFinalFloor(t *testing.T) {
	tests := []struct {
		name         string
		instructions string
		wantFloor    int
		wantError    bool
	}{
		{"empty", "", 0, false},
		{"simple up", "(", 1, false},
		{"simple down", ")", -1, false},
		{"mixed 1", "()()", 0, false},
		{"mixed 2", "(()())", 2, false},
		{"back to zero", ")(", 0, false},
		{"more downs", ")))", -3, false},
		{"invalid char", "(a)", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FinalFloor(tt.instructions)
			if tt.wantError {
				if err == nil {
					t.Fatalf("FinalFloor(%q) expected error, got nil", tt.instructions)
				}
				if err != ErrInvalidInstruction {
					t.Fatalf("FinalFloor(%q) expected ErrInvalidInstruction, got %v", tt.instructions, err)
				}
			} else {
				if err != nil {
					t.Fatalf("FinalFloor(%q) unexpected error: %v", tt.instructions, err)
				}
				if got != tt.wantFloor {
					t.Fatalf("FinalFloor(%q) = %d, want %d", tt.instructions, got, tt.wantFloor)
				}
			}
		})
	}
}

func TestFirstBasementPosition(t *testing.T) {
	tests := []struct {
		name         string
		instructions string
		wantPos      int
		wantError    bool
	}{
		{"empty", "", 0, false},
		{"never basement", "(((", 0, false},
		{"first char basement", ")", 1, false},
		{"later basement", "()())", 5, false},
		{"back and forth no basement", "()()()", 0, false},
		{"invalid char", "())x", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstBasementPosition(tt.instructions)
			if tt.wantError {
				if err == nil {
					t.Fatalf("FirstBasementPosition(%q) expected error, got nil", tt.instructions)
				}
				if err != ErrInvalidInstruction {
					t.Fatalf("FirstBasementPosition(%q) expected ErrInvalidInstruction, got %v", tt.instructions, err)
				}
			} else {
				if err != nil {
					t.Fatalf("FirstBasementPosition(%q) unexpected error: %v", tt.instructions, err)
				}
				if got != tt.wantPos {
					t.Fatalf("FirstBasementPosition(%q) = %d, want %d", tt.instructions, got, tt.wantPos)
				}
			}
		})
	}
}
