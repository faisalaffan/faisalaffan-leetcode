# 3420 — Count Non Decreasing Subarrays After K Operations

## Deskripsi

**Soal:** [3420. Count Non Decreasing Subarrays After K Operations](https://leetcode.com/problems/count-non-decreasing-subarrays-after-k-operations/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser)

## Solusi Go

```go
package main

// LeetCode #3420: Count Non-Decreasing Subarrays After K Operations
// https://leetcode.com/problems/count-non-decreasing-subarrays-after-k-operations/
// Difficulty: Hard
//
// Reverse array, monotonic deque tracking cost, sliding window shrink when cost > k.

import "fmt"

func main() {
	fmt.Println(CountNonDecreasingSubarraysAfterKOperations([]int{3, 2, 1, 4}, 2))
}

func CountNonDecreasingSubarraysAfterKOperations(nums []int, k int) int64 {
	n := len(nums)
	// Reverse array: now we want non-increasing in reversed = non-decreasing in original
  // Membuat slice untuk menyimpan hasil
	rev := make([]int, n)
	for i, v := range nums {
		rev[n-1-i] = v
	}

	type pair struct{ val, cnt int }
	var dq []pair
	cost := int64(0)
	ans := int64(0)
	j := 0

	for i := 0; i < n; i++ {
		cnt := 1
		for len(dq) > 0 && dq[len(dq)-1].val <= rev[i] {
			top := dq[len(dq)-1]
			dq = dq[:len(dq)-1]
			cost -= int64(top.val-rev[i]) * int64(top.cnt)
			cnt += top.cnt
		}
		dq = append(dq, pair{rev[i], cnt})

		for cost > int64(k) {
			// Shrink window from left
			first := &dq[0]
			if first.cnt > 1 {
				first.cnt--
				cost -= int64(dq[0].val - rev[j])
			} else {
				dq = dq[1:]
			}
			j++
		}
		ans += int64(i - j + 1)
	}
	return ans
}
```
