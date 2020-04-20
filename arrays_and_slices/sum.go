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

func SumAllTails(numbersToSumOfTails ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSumOfTails {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
