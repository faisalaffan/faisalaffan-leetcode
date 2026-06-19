package main

// LeetCode #900: RLE Iterator
// https://leetcode.com/problems/rle-iterator/
// Difficulty: Medium

import "fmt"

type RLEIterator struct {
	encoding []int
	idx      int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{encoding: encoding, idx: 0}
}

func (this *RLEIterator) Next(n int) int {
	for this.idx < len(this.encoding) {
		if this.encoding[this.idx] >= n {
			this.encoding[this.idx] -= n
			return this.encoding[this.idx+1]
		}
		n -= this.encoding[this.idx]
		this.idx += 2
	}
	return -1
}

func main() {
	// Test: encoding = [3,8,0,9,2,5]
	obj := Constructor([]int{3, 8, 0, 9, 2, 5})
	fmt.Println(obj.Next(2))  // 8
	fmt.Println(obj.Next(1))  // 8
	fmt.Println(obj.Next(1))  // 5
	fmt.Println(obj.Next(2))  // 5

	fmt.Println("---")

	// Test: encoding = [2,1,3,2]
	obj2 := Constructor([]int{2, 1, 3, 2})
	fmt.Println(obj2.Next(3))  // 2
	fmt.Println(obj2.Next(2))  // -1
}
