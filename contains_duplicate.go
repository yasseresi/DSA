package arrays_and_hashing

func containsDuplicate(nums []int) bool {
	for _, b := range nums {
		count := 0
		for _, d := range nums {
			if d == b {
				count++
			}
		}
		if count >= 2 {
			return true
		}
	}
	return false
}
