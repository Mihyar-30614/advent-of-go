package stats

import "errors"

var ErrEmptySlice = errors.New("empty slice")

func Sum(nums []int) int {
	total := 0
	for _, number := range nums {
		total += number
	}
	return total
}

func Min(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, ErrEmptySlice
	}
	minVal := nums[0]
	for _, number := range nums {
		if number < minVal {
			minVal = number
		}
	}
	return minVal, nil
}

func Max(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, ErrEmptySlice
	}
	maxVal := nums[0]
	for _, number := range nums {
		if number > maxVal {
			maxVal = number
		}
	}
	return maxVal, nil
}

func Average(nums []int) (float64, error) {
	if len(nums) == 0 {
		return 0.0, ErrEmptySlice
	}
	avg := float64(Sum(nums)) / float64(len(nums))
	return avg, nil
}
