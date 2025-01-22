package array_slice

func Sum(a []int) int {
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}
