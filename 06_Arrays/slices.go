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
	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var nums1 = make([]int, 0, 5) // make->new
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

	var nums12 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums12)


	
	// Creation
	var s1 = []int{1, 2, 3}     // Literal
	var s2 = make([]int, 5)     // len=5, cap=5
	var s3 = make([]int, 3, 10) // len=3, cap=10
	var s4 []int             // nil slice (len=0, cap=0)

	// Slicing
	a := []int{0, 1, 2, 3, 4, 5}
	b := a[1:4] // [1, 2, 3]       (shares underlying array!)
	c := a[:3]  // [0, 1, 2]
	d := a[2:]  // [2, 3, 4, 5]
	e := a[:]   // full copy reference

	// Full slice expression (controls capacity)
	f := a[1:4:5] // len=3, cap=4   [low:high:max]

	// Append
	s := []int{1, 2}
	s = append(s, 3)        // [1, 2, 3]
	s = append(s, 4, 5, 6)  // [1, 2, 3, 4, 5, 6]
	s = append(s, other...) // append another slice

	// Copy
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	n := copy(dst, src) // n = 3 (elements copied)

	// Delete element at index i
	s = append(s[:i], s[i+1:]...)

	// Insert at index i
	s = append(s[:i], append([]T{val}, s[i:]...)...)

	// Slices package (Go 1.21+)
	slices.Sort(s)
	slices.Contains(s, 42)
	slices.Index(s, 42)

}
