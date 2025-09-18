func titleToNumber(columnTitle string) int {
	mapNums := map[byte]int{
		'A': 1, 'B': 2, 'C': 3, 'D': 4,
		'E': 5, 'F': 6, 'G': 7, 'H': 8,
		'I': 9, 'J': 10, 'K': 11, 'L': 12,
		'M': 13, 'N': 14, 'O': 15, 'P': 16,
		'Q': 17, 'R': 18, 'S': 19, 'T': 20,
		'U': 21, 'V': 22, 'W': 23, 'X': 24,
		'Y': 25, 'Z': 26}
	var sum int
	for i := 0; len(columnTitle) > i; i++ {
		sum += (int(math.Pow(26, float64(len(columnTitle)-i-1)))) * mapNums[columnTitle[i]]
	}
	return int(sum)
}

func main() {
	columnTitle1 := "A"
	fmt.Println(titleToNumber(columnTitle1))

	columnTitle2 := "AB"
	fmt.Println(titleToNumber(columnTitle2))

	columnTitle3 := "ZY"
	fmt.Println(titleToNumber(columnTitle3))
}
