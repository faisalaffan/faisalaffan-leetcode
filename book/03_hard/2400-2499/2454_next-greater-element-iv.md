# 2454 — Next Greater Element Iv

## Deskripsi

**Soal:** [2454. Next Greater Element Iv](https://leetcode.com/problems/next-greater-element-iv/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO), Monotonic Stack (tumpukan monoton)

## Solusi Go

```go
package main

// LeetCode #2454: Next Greater Element IV (Second Greater Element)
// https://leetcode.com/problems/next-greater-element-iv/
// Difficulty: Hard
//
// Two monotonic stacks. stack2 holds elements that have already found their
// first greater element; when we find a greater element for stack2, that's
// the answer for those elements. stack1 holds elements still looking for
// their first greater element.
// Time O(N) | Space O(N)

import "fmt"

func main() {
	// Example 1
	fmt.Println(secondGreaterElement([]int{2, 4, 0, 9, 6}))
	// Example 2
	fmt.Println(secondGreaterElement([]int{3, 3}))
	// Strictly increasing
	fmt.Println(secondGreaterElement([]int{1, 2, 3, 4, 5}))
}

func secondGreaterElement(nums []int) []int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

  // Membuat slice untuk menyimpan hasil
	stack1 := make([]int, 0, n)
  // Membuat slice untuk menyimpan hasil
	stack2 := make([]int, 0, n)
  // Membuat slice untuk menyimpan hasil
	buf := make([]int, 0, n)

	for i := 0; i < n; i++ {
		v := nums[i]

		// Process stack2 first: current element is the 2nd greater
		for len(stack2) > 0 && nums[stack2[len(stack2)-1]] < v {
			ans[stack2[len(stack2)-1]] = v
			stack2 = stack2[:len(stack2)-1]
		}

		// Process stack1: move elements that found their 1st greater to buffer
		for len(stack1) > 0 && nums[stack1[len(stack1)-1]] < v {
			buf = append(buf, stack1[len(stack1)-1])
			stack1 = stack1[:len(stack1)-1]
		}

		// Push buffer to stack2 in reverse order (maintain original order)
		for k := len(buf) - 1; k >= 0; k-- {
			stack2 = append(stack2, buf[k])
		}
		buf = buf[:0]

		stack1 = append(stack1, i)
	}

	return ans
}
```
