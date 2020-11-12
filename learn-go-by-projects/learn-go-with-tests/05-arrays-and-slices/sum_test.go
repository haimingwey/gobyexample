package main

import (
	"reflect"
	"testing"
)

// try run for code coverage : go test -cover

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := SumByArray(numbers)
		want := 15

		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3}
		got := SumBySlices(nums)
		want := 6

		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})

	t.Run("sum all", func(t *testing.T) {
		nums1 := []int{1, 2, 3}
		nums2 := []int{1, 2}
		nums3 := []int{0, 1}
		got := SumAll(nums1, nums2, nums3)
		want := []int{6, 3, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d but got %d", want, got)
		}
	})

}

// 累加切片中除第一个元素的其他元素
func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d but got %d", want, got)
		}
	}
	t.Run("make the sums of slices", func(t *testing.T) {
		got := SumAllTails([]int{0, 2}, []int{1, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
