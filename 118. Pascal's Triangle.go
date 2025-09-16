func generate(numRows int) [][]int {
	slSL := [][]int{}
	for i := 0; i < numRows; i++ {
		sl := make([]int, i+1)
		sl[0], sl[i] = 1, 1
		slSL = append(slSL, sl)
		for j := 1; j < i; j++ {
			slSL[i][j] = slSL[i-1][j-1] + slSL[i-1][j]
			fmt.Println(slSL)
		}
	}
	return slSL
}

func main() {
	numRows1 := 5
	fmt.Println(generate(numRows1))

	numRows2 := 1
	fmt.Println(generate(numRows2))
}
