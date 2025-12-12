package wordfreq

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []string
	}{
		{"empty", "", []string{}},
		{"only separators", "!!! --- ...", []string{}},
		{"simple", "Hello, world!", []string{"hello", "world"}},
		{"hyphens split", "Go-go-go", []string{"go", "go", "go"}},
		{"numbers and letters", "  123!!abc ", []string{"123", "abc"}},
		{"mixed case", "Mix3d CASE and NUM83R5", []string{"mix3d", "case", "and", "num83r5"}},
		{"underscores split", "go_lang", []string{"go", "lang"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tokenize(tt.text)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Tokenize(%q) = %#v, want %#v", tt.text, got, tt.want)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  map[string]int
	}{
		{"empty", []string{}, map[string]int{}},
		{"single", []string{"go"}, map[string]int{"go": 1}},
		{"multiple", []string{"go", "go", "lang"}, map[string]int{"go": 2, "lang": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWords(tt.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("CountWords(%v) = %#v, want %#v", tt.words, got, tt.want)
			}
		})
	}
}

func TestTopN(t *testing.T) {
	counts := map[string]int{
		"go":   3,
		"lang": 2,
		"is":   2,
		"fun":  1,
	}

	t.Run("n negative", func(t *testing.T) {
		_, err := TopN(counts, -1)
		if err == nil {
			t.Fatalf("TopN expected error, got nil")
		}
		if err != ErrInvalidN {
			t.Fatalf("TopN expected ErrInvalidN, got %v", err)
		}
	})

	t.Run("n zero", func(t *testing.T) {
		got, err := TopN(counts, 0)
		if err != nil {
			t.Fatalf("TopN unexpected error: %v", err)
		}
		if len(got) != 0 {
			t.Fatalf("TopN(n=0) expected empty slice, got %#v", got)
		}
	})

	t.Run("top 2 with tie-break", func(t *testing.T) {
		got, err := TopN(counts, 2)
		if err != nil {
			t.Fatalf("TopN unexpected error: %v", err)
		}
		want := []WordCount{
			{"go", 3},
			{"is", 2}, // "is" < "lang" alphabetically
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("TopN(counts,2) = %#v, want %#v", got, want)
		}
	})

	t.Run("n bigger than unique words", func(t *testing.T) {
		got, err := TopN(counts, 10)
		if err != nil {
			t.Fatalf("TopN unexpected error: %v", err)
		}
		want := []WordCount{
			{"go", 3},
			{"is", 2},
			{"lang", 2},
			{"fun", 1},
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("TopN(counts,10) = %#v, want %#v", got, want)
		}
	})
}

func TestIntegration(t *testing.T) {
	text := "Go go GO... lang? go; fun. IS is."
	words := Tokenize(text)
	counts := CountWords(words)
	got, err := TopN(counts, 3)
	if err != nil {
		t.Fatalf("TopN unexpected error: %v", err)
	}

	// Tokenized words should be: go go go lang go fun is is
	// counts: go=4, is=2, fun=1, lang=1
	want := []WordCount{
		{"go", 4},
		{"is", 2},
		{"fun", 1}, // fun and lang tie at 1; "fun" < "lang"
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("integration TopN = %#v, want %#v", got, want)
	}
}
