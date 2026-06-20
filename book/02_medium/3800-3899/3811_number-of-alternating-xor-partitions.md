# 3811 — Number Of Alternating Xor Partitions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func NumberOfAlternatingXorPartitions(nums []int, target1 int, target2 int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP, Prefix Sum

**Waktu:** O(N)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3811: Number of Alternating XOR Partitions
// https://leetcode.com/problems/number-of-alternating-xor-partitions/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: DP with prefix XOR and two hash maps to track alternating pattern.

import "fmt"

const MOD = 1000000007

func NumberOfAlternatingXorPartitions(nums []int, target1 int, target2 int) int {
  // HashMap: O(1) lookup
	cnt1 := make(map[int]int) // ends with target1
  // HashMap: O(1) lookup
	cnt2 := make(map[int]int) // ends with target2

	cnt2[0] = 1
	pre := 0
	ans := 0

	for _, x := range nums {
		pre ^= x

		a := cnt2[pre^target1] // ways block XOR = target1
		b := cnt1[pre^target2] // ways block XOR = target2

		ans = (a + b) % MOD

		cnt1[pre] = (cnt1[pre] + a) % MOD
		cnt2[pre] = (cnt2[pre] + b) % MOD
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(NumberOfAlternatingXorPartitions([]int{2, 3, 1, 4}, 1, 5)) // Expected: 1

	// Example 2
	fmt.Println(NumberOfAlternatingXorPartitions([]int{1, 0, 0}, 1, 0)) // Expected: 3

	// Example 3
	fmt.Println(NumberOfAlternatingXorPartitions([]int{1, 2, 3}, 1, 2))
}
```
