func removeElement(nums []int, val int) int {
	mapDupl := make(map[int]int)
	for i, v := range nums {
		mapDupl[i] = v
	}
	nums = nums[:0]
	for k, v := range mapDupl {
		if mapDupl[k] == val {
			delete(mapDupl, k)
		} else {
			nums = append(nums, v)
		}
	}
	sort.Ints(nums)
	return len(nums)
}

func main() {
	val1 := 2
	slcNum1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(slcNum1, val1))

	val2 := 2
	slcNum2 := []int{3, 2, 2, 3}
	fmt.Println(removeElement(slcNum2, val2))
}
