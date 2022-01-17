package twosums

import fmt

func twoSum(nums []int, target int) []int {
    result := []int{}
    for i,v := range nums {
        fmt.Println(i)
				//this is to iterate j through the front 
        for j := 0; j<i;j++ {
					  //this is to print out the result
            fmt.Println(i,j,target,v, target -v, nums[j])
            if nums[j] == target - v {
                result = append(result,i,j)                
            }
        }
    }
    return result
}
