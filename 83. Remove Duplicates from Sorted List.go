func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}

func main() {
	nums1_1 := []int{1, 2, 3, 0, 0, 0}
	m1 := 3
	nums2_1 := []int{2, 5, 6}
	n1 := 3
	merge(nums1_1, m1, nums2_1, n1)

	nums1_2 := []int{1}
	m2 := 1
	nums2_2 := []int{}
	n2 := 0
	merge(nums1_2, m2, nums2_2, n2)

	nums1_3 := []int{0}
	m3 := 0
	nums2_3 := []int{1}
	n3 := 1
	merge(nums1_3, m3, nums2_3, n3)
}
