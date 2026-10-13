package sorting

// InsertionSort 插入排序（升序，原地排序）。
// 思路：像整理扑克牌一样，从第 2 个元素开始，把每个元素插入到左侧已排序区间的正确位置。
// 时间复杂度：平均/最坏 O(n^2)，最好（已有序）O(n)；空间复杂度 O(1)，稳定。
func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		// 把大于 key 的元素依次后移
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}
