# 0907 — Sum Of Subarray Minimums

## Deskripsi

**Soal:** [0907. Sum Of Subarray Minimums](https://leetcode.com/problems/sum-of-subarray-minimums/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #907: Sum of Subarray Minimums
// https://leetcode.com/problems/sum-of-subarray-minimums/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(SumOfSubarrayMinimums([]int{3, 1, 2, 4}))
	fmt.Println(SumOfSubarrayMinimums([]int{11, 81, 94, 43, 3}))
	fmt.Println(SumOfSubarrayMinimums([]int{71, 55, 82, 55}))
}

// Time: O(n) | Space: O(n)
func SumOfSubarrayMinimums(arr []int) int {
	const mod = 1_000_000_007
	n := len(arr)

  // Membuat slice untuk menyimpan hasil
	prevSmaller := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	nextSmaller := make([]int, n)

	for i := 0; i < n; i++ {
		prevSmaller[i] = -1
		nextSmaller[i] = n
	}

	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			nextSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = nil
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			prevSmaller[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		left := i - prevSmaller[i]
		right := nextSmaller[i] - i
		ans = (ans + arr[i]*left*right) % mod
	}

	return ans
}
```
