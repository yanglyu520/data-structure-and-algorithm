func isAnagram(s string, t string) bool {
	md1 := transformArrayToHashmap(s)
	md2 := transformArrayToHashmap(t)
	fmt.Println(md1)

	return reflect.DeepEqual(md1, md2)

}

//func translate a string into a hashmap of rune digit, and duplicate times of each digit
func transformArrayToHashmap(s string) map[rune]int {
	md := make(map[rune]int)

	for _, c := range s {
		_, exists := md[c]
		if exists {
			md[c]++
		} else {
			md[c] = 1
		}
	}
	return md
}

