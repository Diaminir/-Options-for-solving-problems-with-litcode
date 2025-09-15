func searchInsert(nums []int, target int) int {
	num := map[int]int{}
	for i, _ := range nums {
		if nums[i] == target {
			num[0] = i
		}
	}
	if _, ok := num[0]; !ok {
		for i, _ := range nums {
			if target > nums[len(nums)-1] {
				num[0] += len(nums)
				return num[0]
			}
			if target > nums[i] && target < nums[i+1] {
				num[0] += i + 1
			}
		}
	}
	return num[0]
}

func main() {
	nums1 := []int{1, 3, 5, 6}
	target1 := 5
	fmt.Println(searchInsert(nums1, target1))

	nums2 := []int{1, 3, 5, 6}
	target2 := 2
	fmt.Println(searchInsert(nums2, target2))

	nums3 := []int{1, 3, 5, 6}
	target3 := 7
	fmt.Println(searchInsert(nums3, target3))
}
