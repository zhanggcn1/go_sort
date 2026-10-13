package sorting

// BubbleSort 冒泡排序（升序，原地排序）。
// 思路：每一轮从左到右比较相邻元素，若逆序则交换，使当前未排序区间的最大值“冒泡”到末尾。
// 时间复杂度：平均/最坏 O(n^2)，最好（已有序）O(n)；空间复杂度 O(1)，稳定。
func BubbleSort(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapped = true
			}
		}
		// 某一轮没有任何交换说明数组已有序，提前结束
		if !swapped {
			break
		}
	}
}
