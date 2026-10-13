package sorting

import (
	"math/rand"
	"slices"
	"sort"
	"testing"
)

var sorts = map[string]func([]int){
	"BubbleSort":    BubbleSort,
	"InsertionSort": InsertionSort,
	"MergeSort":     MergeSort,
	"QuickSort":     QuickSort,
	"HeapSort":      HeapSort,
}

// TestEdgeCases 覆盖空数组、单元素、已有序、逆序、重复元素等边界情况。
func TestEdgeCases(t *testing.T) {
	cases := [][]int{
		{},
		{42},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{3, 1, 2, 1, 3},
		{2, 2, 2, 2},
	}
	for name, fn := range sorts {
		for i, c := range cases {
			input := append([]int(nil), c...)
			fn(input)
			if !sort.IntsAreSorted(input) {
				t.Errorf("%s case %d: 结果 %v 未排序", name, i, input)
			}
		}
	}
}

// TestRandomAgainstReference 用随机数据与标准库 sort.Ints 的结果逐项比对。
func TestRandomAgainstReference(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for name, fn := range sorts {
		for trial := 0; trial < 200; trial++ {
			n := rng.Intn(200)
			input := make([]int, n)
			for i := range input {
				input[i] = rng.Intn(1000)
			}
			want := append([]int(nil), input...)
			sort.Ints(want)
			fn(input)
			if !slices.Equal(input, want) {
				t.Fatalf("%s trial %d: got %v, want %v", name, trial, input, want)
			}
		}
	}
}

// TestLargeRandom 大数组冒烟测试（5000 个随机整数）。
func TestLargeRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	n := 5000
	for name, fn := range sorts {
		input := make([]int, n)
		for i := range input {
			input[i] = rng.Intn(1_000_000)
		}
		want := append([]int(nil), input...)
		sort.Ints(want)
		fn(input)
		if !slices.Equal(input, want) {
			t.Fatalf("%s large test: 排序结果与参考不一致", name)
		}
	}
}
