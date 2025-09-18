func singleNumber(nums []int) int {
    mapNums := make(map[int]int)
	for _, v := range nums {
		mapNums[v]++
	}
	var min int
	for k, v := range mapNums {
		if v < mapNums[min] {
			min = k
		}
		if min == 0 {
			min = k
		}
	}
	return min
}

// с помощью побитового сравнения(не мое решение)
// func singleNumber(nums []int) int {
//     result := 0
//     for _, num := range nums {
//         result ^= num
//     }
    
//     return result
// }

func main() {
	s1 := []int{2, 2, 1}
	fmt.Println(singleNumber(s1))

	s2 := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(s2))

	s3 := []int{1}
	fmt.Println(singleNumber(s3))
}
