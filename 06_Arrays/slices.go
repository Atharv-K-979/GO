package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var nums1 = make([]int, 0, 5)                     // make->new
	// capacity -> maximum numbers of elements can fit
	fmt.Println(cap(nums1))
	fmt.Println(nums1 == nil)

	nums2 := []int{}

	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)

	nums2[0] = 3
	nums2[1] = 5
	fmt.Println(nums2)
	fmt.Println(cap(nums2))
	fmt.Println(len(nums2))

	var nums3 = make([]int, 0, 5)
	nums = append(nums, 2)
	var nums4 = make([]int, len(nums))

	// // copy function
	copy(nums3, nums4)
	fmt.Println(nums3, nums4)

	// slice operator

	var nums5 = []int{1, 2, 3, 4, 5}
	fmt.Println(nums5[0:1])
	fmt.Println(nums5[:2])
	fmt.Println(nums5[1:])

	// slices
	var nums11 = []int{1, 2, 3}
	var nums22 = []int{1, 2, 4}

	fmt.Println(slices.Equal(nums11, nums22))

	var nums12= [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums12)

}