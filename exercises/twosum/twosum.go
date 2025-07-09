package twosum

// twoSum returns indices of the two numbers such that they add up to target.
func twoSum(nums []int, target int) []int {

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
