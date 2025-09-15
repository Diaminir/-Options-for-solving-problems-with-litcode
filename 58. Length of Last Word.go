func lengthOfLastWord(s string) int {
	return len(strings.Fields(s)[len(strings.Fields(s))-1])
}

func main() {
	s1 := "Hello World"
	fmt.Println(lengthOfLastWord(s1))

	s2 := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s2))

	s3 := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s3))
}
