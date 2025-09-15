func removeDuplicates(nums []int) int {
  mapDupl := make(map[int]int)
	for _, v := range nums {
		mapDupl[v]++
	}
	nums = nums[:0]
	for k, _ := range mapDupl {
		nums = append(nums, k)
	}
	sort.Ints(nums)
	return len(nums)
}

func main() {
	slcNum1 := []int{1,1,2}
	fmt.Println(removeDuplicates(slcNum1))

	slcNum2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeDuplicates(slcNum2))
}
