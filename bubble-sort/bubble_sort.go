package bubble_sort

import "fmt"

func BubbleSort(numbers []int) []int  {

	for i := 0; i<len(numbers); i++ {
		numbers, didSwap := sweep(numbers, i)
		if !didSwap {
			return numbers
		}
		fmt.Printf("sort phase: %d - %+v\n", i, numbers)
	}
	return numbers
}

func sweep(numbers []int, previousPhases int) ([]int, bool) {
	n  := len(numbers)
	first_index := 0
	second_index := 1
	didSwap := false

	for second_index < (n - previousPhases) {

		firstNumber := numbers[first_index]
		secondNumber := numbers[second_index]
		if firstNumber > secondNumber {
			//swap numbers
			numbers[first_index] = secondNumber
			numbers[second_index] = firstNumber
			didSwap = true
		}

		first_index++
		second_index++
	}
	return numbers, didSwap
}