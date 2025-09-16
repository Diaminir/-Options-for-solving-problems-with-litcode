func getRow(rowIndex int) []int {
	slSL := [][]int{}
	numRows := rowIndex + 1
	for i := 0; i < numRows; i++ {
		sl := make([]int, i+1)
		sl[0], sl[i] = 1, 1
		slSL = append(slSL, sl)
		for j := 1; j < i; j++ {
			slSL[i][j] = slSL[i-1][j-1] + slSL[i-1][j]
		}
	}
	return slSL[rowIndex]
}

func main() {
	rowIndex1 := 3
	fmt.Println(getRow(rowIndex1))

	rowIndex2 := 0
	fmt.Println(getRow(rowIndex2))

	rowIndex3 := 1
	fmt.Println(getRow(rowIndex3))
}
