func twoSums(nums []int, target int) []int {
  for i, left := range nums {
    for j, right := range nums && i != j {
      if left + right == target {
        return []int{i,j}
      }
    }
  }
  return nil
}
