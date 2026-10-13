package sorting

// QuickSort 快速排序（升序，原地排序）。
// 思路：分治。选一个基准（pivot），把数组分成“小于基准”和“大于等于基准”两部分，
// 再递归排序这两部分。这里使用 Lomuto 分区。
// 时间复杂度：平均 O(n log n)，最坏（如已有序且选尾元素为基准）O(n^2)；空间复杂度 O(log n)（递归栈），不稳定。
func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

// partition 以区间最后一个元素为基准做 Lomuto 分区，返回基准最终所在下标。
// 返回后：[lo, p-1] 均 < 基准，[p+1, hi] 均 >= 基准。
func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}
