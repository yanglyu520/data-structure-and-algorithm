func removeDuplicates(nums []int) int {
    k := 0
    for i,_ :=range nums{
        if nums[i] != nums[k]{
            nums[k+1]=nums[i]
            k++
        }
    }
    return k+1
}
