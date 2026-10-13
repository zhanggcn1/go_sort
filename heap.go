package sorting

// HeapSort 堆排序（升序，原地排序）。
// 思路：先把数组建成大顶堆（父节点 >= 子节点），此时堆顶是最大值；
// 然后把堆顶交换到数组末尾，缩小堆范围并重新调整堆，重复直到堆为空。
// 时间复杂度：平均/最坏/最好均为 O(n log n)；空间复杂度 O(1)，不稳定。
func HeapSort(a []int) {
	n := len(a)

	// 建堆：从最后一个非叶子节点开始自底向上调整
	for i := n/2 - 1; i >= 0; i-- {
		heapify(a, n, i)
	}

	// 依次取出堆顶最大值放到末尾
	for i := n - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		heapify(a, i, 0)
	}
}

// heapify 对以 root 为根的子树做下沉调整，使其满足大顶堆性质（堆大小为 n）。
func heapify(a []int, n, root int) {
	largest := root
	left := 2*root + 1
	right := 2*root + 2

	if left < n && a[left] > a[largest] {
		largest = left
	}
	if right < n && a[right] > a[largest] {
		largest = right
	}
	if largest != root {
		a[root], a[largest] = a[largest], a[root]
		heapify(a, n, largest)
	}
}
