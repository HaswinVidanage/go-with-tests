package arrays

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func GetDiff(a []int, marked []int) int {
	set := make(map[int]bool)
	for _, i := range marked {
		set[i] = true
	}

	totalUnreadCount := len(a)

	for _, i := range a {
		if set[i] == true {
			totalUnreadCount -= 1
		}
	}

	return totalUnreadCount
}
