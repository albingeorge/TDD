package sum

func Sum(input []int) int {
	sum := 0
	for _, v := range input {
		sum += v
	}
	return sum
}
