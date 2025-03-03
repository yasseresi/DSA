package arrays_and_hashing

func twoSum(nums []int, target int) []int {
	list := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				list[0] = i
				list[1] = j
			}
		}
	}
	return list
}
