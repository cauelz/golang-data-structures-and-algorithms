package twosum

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// Test cases for all approaches
var testCases = []struct {
	name     string
	nums     []int
	target   int
	expected []int
}{
	{
		name:     "Basic case",
		nums:     []int{2, 7, 11, 15},
		target:   9,
		expected: []int{0, 1},
	},
	{
		name:     "Numbers at end",
		nums:     []int{3, 2, 4},
		target:   6,
		expected: []int{1, 2},
	},
	{
		name:     "Same numbers",
		nums:     []int{3, 3},
		target:   6,
		expected: []int{0, 1},
	},
	{
		name:     "Large numbers",
		nums:     []int{1000000, 2000000, 3000000, 4000000},
		target:   5000000,
		expected: []int{1, 3},
	},
	{
		name:     "Negative numbers",
		nums:     []int{-1, -2, -3, -4, -5},
		target:   -8,
		expected: []int{2, 4},
	},
	{
		name:     "No solution",
		nums:     []int{1, 2, 3, 4},
		target:   10,
		expected: nil,
	},
}

// Helper function to check if result is valid
func isValidResult(nums []int, target int, result []int) bool {
	if result == nil {
		return false
	}
	if len(result) != 2 {
		return false
	}
	if result[0] < 0 || result[0] >= len(nums) || result[1] < 0 || result[1] >= len(nums) {
		return false
	}
	if result[0] == result[1] {
		return false
	}
	return nums[result[0]]+nums[result[1]] == target
}

// Test all approaches with the same test cases
func TestAllApproaches(t *testing.T) {
	approaches := []struct {
		name string
		fn   func([]int, int) []int
	}{
		{"HashMap", TwoSumHashMap},
		{"BruteForce", TwoSumBruteForce},
		{"TwoPointers", TwoSumTwoPointers},
		{"BinarySearch", TwoSumBinarySearch},
		{"Optimized", TwoSumOptimized},
	}

	for _, approach := range approaches {
		t.Run(approach.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					result := approach.fn(tc.nums, tc.target)
					
					if tc.expected == nil {
						if result != nil {
							t.Errorf("Expected nil, got %v", result)
						}
					} else {
						if !isValidResult(tc.nums, tc.target, result) {
							t.Errorf("Invalid result: %v for nums=%v, target=%d", result, tc.nums, tc.target)
						}
					}
				})
			}
		})
	}
}

// Benchmark tests to compare performance
func BenchmarkApproaches(b *testing.B) {
	// Create a large test case
	largeNums := make([]int, 10000)
	for i := range largeNums {
		largeNums[i] = i
	}
	target := 19998 // Should find indices 9998 and 10000

	approaches := []struct {
		name string
		fn   func([]int, int) []int
	}{
		{"HashMap", TwoSumHashMap},
		{"BruteForce", TwoSumBruteForce},
		{"TwoPointers", TwoSumTwoPointers},
		{"BinarySearch", TwoSumBinarySearch},
		{"Optimized", TwoSumOptimized},
	}

	for _, approach := range approaches {
		b.Run(approach.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				approach.fn(largeNums, target)
			}
		})
	}
}

// Performance comparison test
func TestPerformanceComparison(t *testing.T) {
	// Create test data of different sizes
	testSizes := []int{100, 1000, 5000}
	
	for _, size := range testSizes {
		t.Run(fmt.Sprintf("Size_%d", size), func(t *testing.T) {
			nums := make([]int, size)
			for i := range nums {
				nums[i] = i
			}
			target := size * 2 - 2 // Should find indices size-2 and size-1
			
			approaches := []struct {
				name string
				fn   func([]int, int) []int
			}{
				{"HashMap", TwoSumHashMap},
				{"BruteForce", TwoSumBruteForce},
				{"TwoPointers", TwoSumTwoPointers},
				{"BinarySearch", TwoSumBinarySearch},
				{"Optimized", TwoSumOptimized},
			}
			
			for _, approach := range approaches {
				t.Run(approach.name, func(t *testing.T) {
					start := time.Now()
					result := approach.fn(nums, target)
					duration := time.Since(start)
					
					if !isValidResult(nums, target, result) {
						t.Errorf("Invalid result: %v", result)
					}
					
					t.Logf("%s took %v for size %d", approach.name, duration, size)
				})
			}
		})
	}
}

// Edge case tests
func TestEdgeCases(t *testing.T) {
	edgeCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Empty array",
			nums:     []int{},
			target:   0,
			expected: nil,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			target:   2,
			expected: nil,
		},
		{
			name:     "Two elements no solution",
			nums:     []int{1, 2},
			target:   5,
			expected: nil,
		},
		{
			name:     "All same numbers",
			nums:     []int{1, 1, 1, 1},
			target:   2,
			expected: []int{0, 1}, // Should return first two
		},
	}

	approaches := []struct {
		name string
		fn   func([]int, int) []int
	}{
		{"HashMap", TwoSumHashMap},
		{"BruteForce", TwoSumBruteForce},
		{"TwoPointers", TwoSumTwoPointers},
		{"BinarySearch", TwoSumBinarySearch},
		{"Optimized", TwoSumOptimized},
	}

	for _, approach := range approaches {
		t.Run(approach.name, func(t *testing.T) {
			for _, tc := range edgeCases {
				t.Run(tc.name, func(t *testing.T) {
					result := approach.fn(tc.nums, tc.target)
					
					if tc.expected == nil {
						if result != nil {
							t.Errorf("Expected nil, got %v", result)
						}
					} else {
						if !reflect.DeepEqual(result, tc.expected) {
							t.Errorf("Expected %v, got %v", tc.expected, result)
						}
					}
				})
			}
		})
	}
} 