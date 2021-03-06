package bubble_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	t.Run("Simple bubble sort", func(t *testing.T) {
		input := []int{7, 6, 4, 5, 10, 3, 2,1}
		expected := []int{1, 2, 3, 4, 5, 6, 7, 10}
		actual := BubbleSort(input)
		assert.Equal(t, expected, actual)
	})
}
