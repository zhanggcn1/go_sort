package sorting

// MergeSort 归并排序（升序，原地排序，借助 O(n) 额外空间）。
// 思路：分治。把数组递归对半分成子数组，各自排好后，再把两个有序子数组合并成一个有序数组。
// 时间复杂度：平均/最坏/最好均为 O(n log n)；空间复杂度 O(n)，稳定。
func MergeSort(a []int) {
	tmp := make([]int, len(a))
	mergeSort(a, tmp, 0, len(a)-1)
}

func mergeSort(a, tmp []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSort(a, tmp, lo, mid)
	mergeSort(a, tmp, mid+1, hi)
	merge(a, tmp, lo, mid, hi)
}

// merge 合并两个有序区间 [lo, mid] 与 [mid+1, hi]。
func merge(a, tmp []int, lo, mid, hi int) {
	i, j, k := lo, mid+1, lo
	for i <= mid && j <= hi {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = a[i]
		i++
		k++
	}
	for j <= hi {
		tmp[k] = a[j]
		j++
		k++
	}
	copy(a[lo:hi+1], tmp[lo:hi+1])
}
