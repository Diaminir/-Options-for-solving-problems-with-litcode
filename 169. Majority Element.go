func majorityElement(nums []int) int {
mapNums := make(map[int]int)
	for _, v := range nums {
		mapNums[v]++
	}
	var max int
	for k, v := range mapNums {
		if v > mapNums[max] {
			max = k
		}
	}
	return max
}

func main() {
	s1 := []int{3, 2, 3}
	fmt.Println(majorityElement(s1))

	s2 := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(s2))
}
