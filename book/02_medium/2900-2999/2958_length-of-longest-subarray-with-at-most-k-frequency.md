# 2958 — Length Of Longest Subarray With At Most K Frequency

## Deskripsi

**Soal:** [2958. Length Of Longest Subarray With At Most K Frequency](https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2958: Length of Longest Subarray With at Most K Frequency
// https://leetcode.com/problems/length-of-longest-subarray-with-at-most-k-frequency/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2))
	fmt.Println(maxSubarrayLength([]int{1, 2, 1, 2, 1, 2, 1, 2}, 1))
	fmt.Println(maxSubarrayLength([]int{5, 5, 5, 5, 5, 5, 5}, 4))
}

func maxSubarrayLength(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	for i, j, n := 0, 0, len(nums); i < n; i++ {
		cnt[nums[i]]++
		for ; cnt[nums[i]] > k; j++ {
			cnt[nums[j]]--
		}
		if i-j+1 > ans {
			ans = i - j + 1
		}
	}
	return
}
```
