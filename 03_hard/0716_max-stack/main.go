package main

import (
	"fmt"
	"math"
)

// LeetCode #716: Max Stack
// https://leetcode.com/problems/max-stack/
// Difficulty: Hard
//
// Design a stack that supports push, pop, top, peekMax, and popMax.
// Uses two stacks: one for values, one for max tracking.
// popMax is O(n) using a temporary buffer.

type MaxStack struct {
	stack []int
	max   []int
}

func Constructor() MaxStack {
	return MaxStack{stack: []int{}, max: []int{}}
}

func (ms *MaxStack) Push(x int) {
	ms.stack = append(ms.stack, x)
	if len(ms.max) == 0 || x >= ms.max[len(ms.max)-1] {
		ms.max = append(ms.max, x)
	}
}

func (ms *MaxStack) Pop() int {
	if len(ms.stack) == 0 {
		return math.MinInt32
	}
	val := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	if val == ms.max[len(ms.max)-1] {
		ms.max = ms.max[:len(ms.max)-1]
	}
	return val
}

func (ms *MaxStack) Top() int {
	if len(ms.stack) == 0 {
		return math.MinInt32
	}
	return ms.stack[len(ms.stack)-1]
}

func (ms *MaxStack) PeekMax() int {
	if len(ms.max) == 0 {
		return math.MinInt32
	}
	return ms.max[len(ms.max)-1]
}

func (ms *MaxStack) PopMax() int {
	maxVal := ms.PeekMax()
	var buf []int
	for ms.Top() != maxVal {
		buf = append(buf, ms.Pop())
	}
	ms.Pop()
	for i := len(buf) - 1; i >= 0; i-- {
		ms.Push(buf[i])
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Standard example
	ms := Constructor()
	ms.Push(5)
	ms.Push(1)
	ms.Push(5)
	fmt.Println(ms.Top())     // 5
	fmt.Println(ms.PopMax())  // 5
	fmt.Println(ms.Top())     // 1
	fmt.Println(ms.PeekMax()) // 5
	fmt.Println(ms.Pop())     // 1
	fmt.Println(ms.Top())     // 5

	fmt.Println("---")

	// Push/pop sequence
	ms2 := Constructor()
	ms2.Push(2)
	ms2.Push(1)
	ms2.Push(3)
	ms2.Push(2)
	fmt.Println(ms2.PeekMax()) // 3
	fmt.Println(ms2.PopMax())  // 3
	fmt.Println(ms2.PeekMax()) // 2
	fmt.Println(ms2.Pop())     // 2
	fmt.Println(ms2.PeekMax()) // 2
	fmt.Println(ms2.Pop())     // 1
	fmt.Println(ms2.PeekMax()) // 2
	fmt.Println(ms2.Pop())     // 2

	fmt.Println("---")

	// Single element
	ms3 := Constructor()
	ms3.Push(42)
	fmt.Println(ms3.PopMax())  // 42
	fmt.Println(ms3.PeekMax()) // -2147483648 (empty sentinel)
	fmt.Println(ms3.Pop())     // -2147483648 (empty sentinel)
}
