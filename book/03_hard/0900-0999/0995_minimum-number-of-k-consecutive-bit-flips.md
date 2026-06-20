# 0995 — Minimum Number Of K Consecutive Bit Flips

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minKBitFlips(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** BFS

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **BFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #995: Minimum Number of K Consecutive Bit Flips
// https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips/
// Difficulty: Hard
//
// Approach: Greedy + queue (or flip-tracking with a boolean array).
//   Maintain a queue of flip start indices. When processing position i:
//   - Pop from queue front if the flip ended (queue[0]+k == i).
//   - The current effective value = nums[i] ^ (len(queue)%2).
//   - If it's 0, we must flip: push i into the queue.

import "fmt"

func main() {
	fmt.Println(minKBitFlips([]int{0, 1, 0}, 1))    // 2
	fmt.Println(minKBitFlips([]int{1, 1, 0}, 2))    // -1
	fmt.Println(minKBitFlips([]int{0, 0, 0, 1, 0, 1, 1, 0}, 3)) // 3
}

func minKBitFlips(nums []int, k int) int {
	n := len(nums)
	flipped := make([]bool, n)
	curFlips := 0
	ans := 0

	for i := 0; i < n; i++ {
		if i >= k && flipped[i-k] {
			curFlips--
		}
		if curFlips%2 == 0 && nums[i] == 0 || curFlips%2 == 1 && nums[i] == 1 {
			if i+k > n {
				return -1
			}
			flipped[i] = true
			curFlips++
			ans++
		}
	}
	return ans
}
```
