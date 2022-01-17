func twoSum(nums []int, target int) []int {
    for i,v := range nums {
        for j := 0; j<i;j++ {
            if nums[j] == target - v {
                return []int{i,j}                
            }
        }
    }
    return nil
}
