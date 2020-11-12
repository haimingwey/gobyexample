package main

/**
An interesting property of arrays is that the size is encoded in its type.
If you try to pass an [4]int into a function that expects [5]int, it won't compile.
They are different types so it's just the same as trying to pass a string into a function that wants an int.
*/
func SumByArray(nums [5]int) int {
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum
}

func SumBySlices(nums []int) int {
	sum := 0
	for _, number := range nums {
		sum += number
	}
	return sum
}

func SumAll(nums ...[]int) []int {
	var sums []int
	for _, numbers := range nums {
		sums = append(sums, SumBySlices(numbers))
	}

	return sums
}

func SumAllTails(nums ...[]int) []int {
	var sums []int
	for _, numbers := range nums {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, SumBySlices(tail))
		}
	}
	return sums
}
