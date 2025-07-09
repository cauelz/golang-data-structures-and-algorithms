package containsduplicate

// ContainsDuplicate returns true if any value appears at least twice in the array.
func ContainsDuplicate(nums []int) bool {
	
	mapper := make(map[int]int)

	for _, value := range nums {

		if _, ok := mapper[value]; ok {
			return true
		}

		mapper[value]++

	}

	return false
} 