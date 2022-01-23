func findShortestSubArray(nums []int) int {
    nums_degree := findDegreeOfArray (nums)
    md := transformArrayToHashmap(nums)
    //find the smallest length of a contiguous subarray of nums
     ans := len(nums)
     for v,_ := range md {
        if md[v][2] == nums_degree {
            if ans > md[v][1] - md[v][0] + 1 {
                ans = md[v][1] - md[v][0] + 1
            }
        }
    }
    
    return ans
}

func transformArrayToHashmap  (nums []int) map[int][]int{
    //create a hashmap to record all the index of an item occurs
    m := make(map[int][]int)
    for i,v := range nums{
        m[v] = append(m[v], i)
    }
    
    //create a new hashmap to record the left index, right index, and duplicated times of an itme
    md := make(mapint][]int)
    for v,_ := range m {
        j := len(m[v])
       md[v] = append(md[v], m[v][0], m[v][j-1],j)
    }
    
    return md
    
}

func findDegreeOfArray (nums []int) int {
    md := transformArrayToHashmap(nums)
    //find the degree of a given array
    nums_degree := 0
    for v,_ := range md {
        if md[v][2] > nums_degree {
            nums_degree = md[v][2]
        }
    }
    
    return nums_degree
    
}

