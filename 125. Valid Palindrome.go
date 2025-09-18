func isPalindrome(s string) bool {
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	str := reg.ReplaceAllString(strings.ToLower(s), "")
	revR := []rune(str)
	slices.Reverse(revR)
	rev := string(revR)
	for i := 0; i < len(str); i++ {
		if rev[i] != str[i] {
			return false
		}
	}
	return true
}

func main() {
	s1 := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s1))

	s2 := "race a car"
	fmt.Println(isPalindrome(s2))

	s3 := " "
	fmt.Println(isPalindrome(s3))

	s4 := "0P"
	fmt.Println(isPalindrome(s4))
}
