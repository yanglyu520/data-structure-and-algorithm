func twoSum(nums []int, target int) []int {
    for i,v := range nums{
        for j:=i+1;j<len(nums);j++ {
            if nums[j] == target - v{
                return []int{i,j}
            }
        }
    }
    return nil
    
}
