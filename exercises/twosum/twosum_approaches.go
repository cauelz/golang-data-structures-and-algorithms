package twosum

import "sort"

// TwoSumHashMap - Hash Map Approach (Your current implementation)
// Time: O(n), Space: O(n)
func TwoSumHashMap(nums []int, target int) []int {
	sum := make(map[int]int, len(nums))

	for index, value := range nums {
		complement := target - value

		if index2, ok := sum[complement]; ok {
			return []int{index2, index}
		}

		sum[value] = index
	}

	return nil
}

// TwoSumBruteForce - Brute Force Approach
// Time: O(n²), Space: O(1)
func TwoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// TwoSumTwoPointers - Two Pointers Approach (requires sorted array)
// Time: O(n log n), Space: O(n) - because we need to preserve original indices
func TwoSumTwoPointers(nums []int, target int) []int {
	// Create a slice of pairs (value, original_index) to preserve indices
	type pair struct {
		value int
		index int
	}
	
	pairs := make([]pair, len(nums))
	for i, num := range nums {
		pairs[i] = pair{value: num, index: i}
	}
	
	// Sort by value
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})
	
	// Two pointers approach
	left, right := 0, len(pairs)-1
	
	for left < right {
		sum := pairs[left].value + pairs[right].value
		
		if sum == target {
			return []int{pairs[left].index, pairs[right].index}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	
	return nil
}

// TwoSumBinarySearch - Binary Search Approach
// Time: O(n log n), Space: O(n) - because we need to preserve original indices
func TwoSumBinarySearch(nums []int, target int) []int {
	// Create a slice of pairs (value, original_index) to preserve indices
	type pair struct {
		value int
		index int
	}
	
	pairs := make([]pair, len(nums))
	for i, num := range nums {
		pairs[i] = pair{value: num, index: i}
	}
	
	// Sort by value
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})
	
	// For each element, use binary search to find complement
	for i, p := range pairs {
		complement := target - p.value
		
		// Binary search for complement
		left, right := i+1, len(pairs)-1
		for left <= right {
			mid := left + (right-left)/2
			
			if pairs[mid].value == complement {
				return []int{p.index, pairs[mid].index}
			} else if pairs[mid].value < complement {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	
	return nil
}

// TwoSumOptimized - Optimized version with early termination and validation
// Time: O(n), Space: O(n)
func TwoSumOptimized(nums []int, target int) []int {
	// Input validation
	if len(nums) < 2 {
		return nil
	}
	
	sum := make(map[int]int, len(nums))
	
	for index, value := range nums {
		complement := target - value
		
		// Check if complement exists
		if index2, ok := sum[complement]; ok {
			return []int{index2, index}
		}
		
		// Store current value and its index
		sum[value] = index
	}
	
	return nil
} 