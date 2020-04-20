package arrays_and_slices

func Sum(numbers []int) (sum int) {
	for _, number  := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int){
	// lengthOfNumbers := len(numbersToSum)
	// sums = make([]int, lengthOfNumbers)

	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		sums = append(sums, Sum(numbers))
	}

	return
}
