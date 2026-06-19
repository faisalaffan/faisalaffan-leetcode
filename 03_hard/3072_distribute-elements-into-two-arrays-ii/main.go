package main

// LeetCode #3072: Distribute Elements Into Two Arrays II
// https://leetcode.com/problems/distribute-elements-into-two-arrays-ii/
// Difficulty: Hard
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

// Fenwick Tree (Binary Indexed Tree) for counting elements greater than a value.
type BIT struct {
	tree []int
}

func NewBIT(size int) *BIT {
	return &BIT{tree: make([]int, size+1)}
}

func (b *BIT) Update(idx, delta int) {
	idx++
	for idx < len(b.tree) {
		b.tree[idx] += delta
		idx += idx & -idx
	}
}

func (b *BIT) Query(idx int) int {
	idx++
	sum := 0
	for idx > 0 {
		sum += b.tree[idx]
		idx -= idx & -idx
	}
	return sum
}

// QueryRange returns count of elements in [l, r].
func (b *BIT) QueryRange(l, r int) int {
	if l > r {
		return 0
	}
	return b.Query(r) - b.Query(l-1)
}

func ResultArray(nums []int) []int {
	n := len(nums)

	// Coordinate compression
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)
	coord := make(map[int]int)
	for i, v := range sorted {
		if _, ok := coord[v]; !ok {
			coord[v] = i
		}
	}

	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}
	bit1 := NewBIT(n)
	bit2 := NewBIT(n)

	bit1.Update(coord[nums[0]], 1)
	bit2.Update(coord[nums[1]], 1)

	for i := 2; i < n; i++ {
		idx := coord[nums[i]]

		// Count elements > nums[i] in each array
		// greater = total - (elements <= nums[i])
		greater1 := len(arr1) - bit1.Query(idx)
		greater2 := len(arr2) - bit2.Query(idx)

		if greater1 > greater2 {
			arr1 = append(arr1, nums[i])
			bit1.Update(idx, 1)
		} else if greater2 > greater1 {
			arr2 = append(arr2, nums[i])
			bit2.Update(idx, 1)
		} else {
			// equal greater count: put in smaller array, or arr1 if same size
			if len(arr1) <= len(arr2) {
				arr1 = append(arr1, nums[i])
				bit1.Update(idx, 1)
			} else {
				arr2 = append(arr2, nums[i])
				bit2.Update(idx, 1)
			}
		}
	}

	return append(arr1, arr2...)
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", ResultArray([]int{2, 1, 3, 3}))
	// Expected: [2, 1, 3, 3]

	// Test case 2
	fmt.Println("Test 2:", ResultArray([]int{5, 4, 3, 8}))
	// Expected: [5, 3, 4, 8]

	// Test case 3
	fmt.Println("Test 3:", ResultArray([]int{1, 2, 3, 4, 5}))
	// Expected: [1, 3, 2, 4, 5] or similar valid distribution
}
