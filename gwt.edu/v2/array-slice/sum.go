package array_slice

func Sum(a []int) int {
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}

func SumAll(xs ...[]int) []int {
	accumulator := make([]int, len(xs))
	for i, v := range xs {
		accumulator[i] = Sum(v)
	}
	return accumulator
}
