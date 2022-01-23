func moveZeroes(nums []int) {
	k := 0
	for i, _ := range nums {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}
	if k < len(nums) {
		for j := k; j < len(nums); j++ {
			nums[j] = 0
		}
	}

}
