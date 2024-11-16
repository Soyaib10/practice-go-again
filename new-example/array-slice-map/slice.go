package main

import (
	"fmt"
	"slices"
	"sort"
)

func modify(a []int) {
	a[0] = 2345
}

func slice() {
	nums := make([]int, 3)
	for i := 0; i < 3; i++ {
		nums[i] = i + 1
	}

	modify(nums)
	fmt.Print(nums)
}

func functions() {
	// copy- two different entity
	src := []int{3, 4}
	dest := make([]int, len(src))
	copy(dest, src[:2])
	fmt.Println(dest)
	dest[1] = 5423
	fmt.Println(src, dest)

	// slicing
	nums := []int{1, 3, 11, 4, 5, 5, 4, 3, 2, 2, 2, 1}
	sub := nums[1:4]
	fmt.Println(nums, sub)
	sub[2] = 245
	fmt.Println(nums, sub)

	// remove
	// single element
	idx := 2
	nums = append(nums[:idx], nums[idx+1:]...)
	fmt.Println(nums)
	// multiple element
	start, end := 2, 4
	nums = append(nums[0:start], nums[end:]...)
	fmt.Println(nums)

	// reverse
	slices.Reverse(nums)
	fmt.Println(nums)

	// sort
	sort.Ints(nums)
	fmt.Println(nums) // ascending
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // descending
	fmt.Println(nums)

	sort.Ints(nums[2:6])

	// custom slice


	// make slice empty
	nums = nums[:0]
	fmt.Println(nums)

}

func exercise() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	// sort
	sort.Ints(s[0:n])
	fmt.Println(s)

	// reverse
	slices.Reverse(s)
	fmt.Println(s)

	// remove 2nd element
	s = append(s[:2], s[2 + 1:]...)
	fmt.Println(s)

	if (len(s) == 0) {
		fmt.Println("slice is empty.")
	} else {
		fmt.Println("not empty")
	}

    for index, value := range s[0:n-1] {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

}

