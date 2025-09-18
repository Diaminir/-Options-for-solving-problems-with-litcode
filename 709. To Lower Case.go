func toLowerCase(s string) string {
	str := []rune{}
	for _, v := range s {
		if v >= 65 && v <= 90 {
			str = append(str, v+32)
		} else {
			str = append(str, v)
		}
	}
	return string(str)
}

func main() {
	s1 := "Hello"
	fmt.Println(toLowerCase(s1))

	s2 := "here"
	fmt.Println(toLowerCase(s2))

	s3 := "LOVELY"
	fmt.Println(toLowerCase(s3))
}
