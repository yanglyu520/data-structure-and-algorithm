func merge(nums1 []int, m int, nums2 []int, n int) {
	//use a new list to store the final result list
	numsnew := []int{}
	i := 0
	j := 0
	//traverse through both lists using 2 pointers and compare between the 2 lists
	for i < m || j < n {
		//condition to keep nums1[i]
		if (i < m && j < n && nums1[i] <= nums2[j]) || j >= n {
			numsnew = append(numsnew, nums1[i])
			i++
		} else {
			//condition to keep num2[j]
			numsnew = append(numsnew, nums2[j])
			j++
		}
	}
	//transfer the new list to nums1
	for k, _ := range numsnew {
		nums1[k] = numsnew[k]
	}

}




