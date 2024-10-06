package main

import "fmt"

func Sum(array []int) (sum int) {
	for _, number := range array {
		sum += number
	}

	return
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	fmt.Println(lengthOfNumbers)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	fmt.Println(sums)
	return sums
}

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

func main() {

}
