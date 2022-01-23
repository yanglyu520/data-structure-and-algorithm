func twoSum(numbers []int, target int) []int {
	j := len(numbers) - 1
	for i, v := range numbers {
		for j > i && numbers[j]+v > target {
			j--
		}
		if j > i && v+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}
