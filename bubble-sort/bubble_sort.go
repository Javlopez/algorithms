package bubble_sort

import "fmt"

func BubbleSort(numbers []int) []int  {

	for i := 0; i<len(numbers); i++ {
		numbers = sweep(numbers)
		fmt.Printf("sort phase: %d - %+v\n", i, numbers)
	}
	return numbers
}

func sweep(numbers []int) []int {
	n  := len(numbers)
	first_index := 0
	second_index := 1

	for second_index < n {

		firstNumber := numbers[first_index]
		secondNumber := numbers[second_index]
		if firstNumber > secondNumber {
			//swap numbers
			numbers[first_index] = secondNumber
			numbers[second_index] = firstNumber
		}

		first_index++
		second_index++
	}
	return numbers
}