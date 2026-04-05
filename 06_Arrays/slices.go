package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// + useful methods

// Definition: A dynamically-sized, flexible view into an underlying array. Contains pointer, length, and capacity.
// Analogy: A slice is like a window on a spreadsheet — you see part of the data, can resize the window, but the spreadsheet (underlying array) is what holds the data.
// Slice Header (24 bytes on 64-bit):
// ┌────────────────────────────┐
// │  ptr  → points to array    │  8 bytes
// │  len  = number of elements │  8 bytes
// │  cap  = max before resize  │  8 bytes
// └────────────────────────────┘

// s := make([]int, 3, 5)

// Slice Header          Underlying Array
// ┌─────────┐          ┌───┬───┬───┬───┬───┐
// │ ptr ─────────────  │ 0 │ 0 │ 0 │ . │ . │
// │ len = 3 │          └───┴───┴───┴───┴───┘
// │ cap = 5 │            ← len=3 →← extra →
// └─────────┘             ──── cap=5 ──────

func main() {
	// uninitialized slice is nil(null)
	var nums []int
	fmt.Println(nums == nil) //true
	fmt.Println(len(nums)) // 0

	// if you dont wan nil
	var nums1 = make([]int, 0, 5) // make->new
	fmt.Println(cap(nums1))  //cap is max no of items you can add
	fmt.Println(nums1 == nil) // false



	nums2 := []int{}
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)

	nums2[0] = 3
	nums2[1] = 5
	fmt.Println(nums2)
	fmt.Println(cap(nums2))  // output: 4 *******(capacity doubles when exceeded)
	fmt.Println(len(nums2))  // output: 2


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

	fmt.Println(slices.Equal(nums11, nums22)) // false as 3!=4

	var nums12 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums12)

}
