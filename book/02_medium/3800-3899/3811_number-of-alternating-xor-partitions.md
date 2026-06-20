# 3811 — Number Of Alternating Xor Partitions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfAlternatingXorPartitions(nums []int, target1 int, target2 int) int
```

> **💡 Hint:** DP with prefix XOR and two hash maps to track alternating pattern.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat map (HashMap) — pencarian O(1)
	cnt1 := make(map[int]int) // ends with target1
  // Membuat map (HashMap) — pencarian O(1)
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
