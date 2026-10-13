// 演示程序：生成随机数组，用五种排序算法分别排序并打印。
package main

import (
	"fmt"
	"math/rand"

	"github.com/zhanggcn1/go_sort"
)

func main() {
	rng := rand.New(rand.NewSource(2026))
	original := make([]int, 12)
	for i := range original {
		original[i] = rng.Intn(100)
	}

	fmt.Printf("原始数组: %v\n\n", original)

	// 每种排序都基于原始数组的副本，避免相互影响
	for _, s := range []struct {
		name string
		fn   func([]int)
	}{
		{"冒泡排序 BubbleSort", sorting.BubbleSort},
		{"插入排序 InsertionSort", sorting.InsertionSort},
		{"归并排序 MergeSort", sorting.MergeSort},
		{"快速排序 QuickSort", sorting.QuickSort},
		{"堆排序 HeapSort", sorting.HeapSort},
	} {
		a := append([]int(nil), original...)
		s.fn(a)
		fmt.Printf("%-24s -> %v\n", s.name, a)
	}
}
