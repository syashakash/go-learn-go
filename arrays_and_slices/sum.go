package arrays_and_slices

func Sum(numbers [5]int) (sum int) {
	for _, number  := range numbers {
		sum += number
	}
	return
}