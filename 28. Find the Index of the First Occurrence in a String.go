func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	haystack1 := "sadbutsad"
	needle1 := "sad"
	fmt.Println(strStr(haystack1, needle1))
  
	haystack2 := "leetcode"
	needle2 := "leeto"
	fmt.Println(strStr(haystack2, needle2))
}
  
