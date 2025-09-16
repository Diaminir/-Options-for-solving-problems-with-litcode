func climbStairs(n int) int {
	sl := []int{0, 1}
	for range n {
		f := sl[len(sl)-1] + sl[len(sl)-2]
		sl = append(sl, f)
	}
	return sl[len(sl)-1]
}

func main() {
	n1 := 2
	fmt.Println(climbStairs(n1))

	n2 := 3
	fmt.Println(climbStairs(n2))

	n3 := 4
	fmt.Println(climbStairs(n3))
}
