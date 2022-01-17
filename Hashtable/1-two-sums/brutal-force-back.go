func twoSum(nums []int, target int) []int {
    result := []int{}
    for i,v := range nums{
        fmt.Println(i)
        for j:=i+1;j<len(nums);j++ {
            fmt.Println(i,j,target,v,target-v, nums[j])
            if nums[j] == target - v{
                result = append(result, i,j)
            }
        }
    }
    return result
    
}