package sort

func BubbleSort(nums []int) {
	l := len(nums)
	isChanded := false
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isChanded = true
			}
		}
		if !isChanded {
			break
		}
	}
}
