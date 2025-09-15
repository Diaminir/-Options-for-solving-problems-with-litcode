func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	res := make([]int, len(digits)+1)
	res[0] = 1
	return res
}

func main() {
	digits1 := []int{1, 2, 3}
	fmt.Println(plusOne(digits1))

	digits2 := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits2))

	digits3 := []int{9}
	fmt.Println(plusOne(digits3))

	digits4 := []int{9, 9, 9}
	fmt.Println(plusOne(digits4))
}
