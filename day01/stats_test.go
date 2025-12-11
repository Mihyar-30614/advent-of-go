package stats

import (
    "math"
    "testing"
)

func TestSum(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {"empty slice", []int{}, 0},
        {"single element", []int{42}, 42},
        {"multiple positive", []int{1, 2, 3, 4}, 10},
        {"with negatives", []int{-5, 10, -3}, 2},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Sum(tt.nums)
            if got != tt.want {
                t.Fatalf("Sum(%v) = %d, want %d", tt.nums, got, tt.want)
            }
        })
    }
}

func TestMin(t *testing.T) {
    tests := []struct {
        name      string
        nums      []int
        want      int
        wantError bool
    }{
        {"empty slice", []int{}, 0, true},
        {"single element", []int{42}, 42, false},
        {"multiple positive", []int{1, 2, 3, 4}, 1, false},
        {"with negatives", []int{-5, 10, -3}, -5, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Min(tt.nums)
            if tt.wantError {
                if err == nil {
                    t.Fatalf("Min(%v) expected error, got nil", tt.nums)
                }
                if err != ErrEmptySlice {
                    t.Fatalf("Min(%v) expected ErrEmptySlice, got %v", tt.nums, err)
                }
            } else {
                if err != nil {
                    t.Fatalf("Min(%v) unexpected error: %v", tt.nums, err)
                }
                if got != tt.want {
                    t.Fatalf("Min(%v) = %d, want %d", tt.nums, got, tt.want)
                }
            }
        })
    }
}

func TestMax(t *testing.T) {
    tests := []struct {
        name      string
        nums      []int
        want      int
        wantError bool
    }{
        {"empty slice", []int{}, 0, true},
        {"single element", []int{42}, 42, false},
        {"multiple positive", []int{1, 2, 3, 4}, 4, false},
        {"with negatives", []int{-5, 10, -3}, 10, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Max(tt.nums)
            if tt.wantError {
                if err == nil {
                    t.Fatalf("Max(%v) expected error, got nil", tt.nums)
                }
                if err != ErrEmptySlice {
                    t.Fatalf("Max(%v) expected ErrEmptySlice, got %v", tt.nums, err)
                }
            } else {
                if err != nil {
                    t.Fatalf("Max(%v) unexpected error: %v", tt.nums, err)
                }
                if got != tt.want {
                    t.Fatalf("Max(%v) = %d, want %d", tt.nums, got, tt.want)
                }
            }
        })
    }
}

func almostEqual(a, b, tolerance float64) bool {
    return math.Abs(a-b) <= tolerance
}

func TestAverage(t *testing.T) {
    tests := []struct {
        name      string
        nums      []int
        want      float64
        wantError bool
    }{
        {"empty slice", []int{}, 0, true},
        {"single element", []int{42}, 42, false},
        {"multiple positive", []int{1, 2, 3, 4}, 2.5, false},
        {"with negatives", []int{-5, 5, -5, 5}, 0, false},
    }

    const tolerance = 1e-9

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Average(tt.nums)
            if tt.wantError {
                if err == nil {
                    t.Fatalf("Average(%v) expected error, got nil", tt.nums)
                }
                if err != ErrEmptySlice {
                    t.Fatalf("Average(%v) expected ErrEmptySlice, got %v", tt.nums, err)
                }
            } else {
                if err != nil {
                    t.Fatalf("Average(%v) unexpected error: %v", tt.nums, err)
                }
                if !almostEqual(got, tt.want, tolerance) {
                    t.Fatalf("Average(%v) = %f, want %f", tt.nums, got, tt.want)
                }
            }
        })
    }
}

