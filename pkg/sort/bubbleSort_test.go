package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{1, 20, 3, 4, 5}
	BubbleSort(nums)
	fmt.Println(nums)
}
