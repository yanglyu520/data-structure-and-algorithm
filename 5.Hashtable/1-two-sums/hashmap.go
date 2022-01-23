func twoSum(nums []int, target int) []int {
    //create a hashmap
    hm := make(map[int]int)
    
    //go through the array, while saving the ones inside a set
    for i, v := range nums {
        //check if this key (target-v) exists in hm map
        j, found := hm[target - v]
        if found {
            return []int{i, j}
        }
        //store to a set
        hm[v] = i   
  }
    return nil
}
