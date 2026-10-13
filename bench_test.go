package sorting

import (
	"math/rand"
	"testing"
)

var benchData []int

// benchInput 返回一份长度为 n 的随机排列副本，避免各基准互相影响。
func benchInput(n int) []int {
	if len(benchData) != n {
		benchData = rand.New(rand.NewSource(7)).Perm(n)
	}
	return append([]int(nil), benchData...)
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := benchInput(1000)
		BubbleSort(a)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := benchInput(1000)
		InsertionSort(a)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := benchInput(1000)
		MergeSort(a)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := benchInput(1000)
		QuickSort(a)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := benchInput(1000)
		HeapSort(a)
	}
}
