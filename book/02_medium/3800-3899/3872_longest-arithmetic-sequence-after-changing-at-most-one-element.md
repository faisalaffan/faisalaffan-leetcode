# 3872 — Longest Arithmetic Sequence After Changing At Most One Element

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestArithmeticSequenceAfterChangingAtMostOneElement(nums []int) int
```

> **💡 Hint:** Compute prefix (arithmetic ending at i) and suffix (arithmetic starting at i).

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Prefix Sum

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3872: Longest Arithmetic Sequence After Changing At Most One Element
// https://leetcode.com/problems/longest-arithmetic-sequence-after-changing-at-most-one-element/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Compute prefix (arithmetic ending at i) and suffix (arithmetic starting at i).
// For each position, try changing it to connect left and right arithmetic sequences.

import "fmt"

func LongestArithmeticSequenceAfterChangingAtMostOneElement(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

  // Alokasi slice integer
	pref := make([]int, n)
	pref[0] = 1
	pref[1] = 2
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			pref[i] = pref[i-1] + 1
		} else {
			pref[i] = 2
		}
	}

  // Alokasi slice integer
	suff := make([]int, n)
	suff[n-1] = 1
	suff[n-2] = 2
	for i := n - 3; i >= 0; i-- {
		if nums[i+2]-nums[i+1] == nums[i+1]-nums[i] {
			suff[i] = suff[i+1] + 1
		} else {
			suff[i] = 2
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if pref[i] > ans {
			ans = pref[i]
		}
	}

	if ans < n {
		ans = max(ans, 1+suff[1])
	}
	if ans < n {
		ans = max(ans, 1+pref[n-2])
	}

	for i := 1; i < n-1; i++ {
		if (nums[i+1]-nums[i-1])%2 != 0 {
			continue
		}
		d := (nums[i+1] - nums[i-1]) / 2

		leftLen := 1
		if i >= 2 && nums[i-1]-nums[i-2] == d {
			leftLen = pref[i-1]
		}

		rightLen := 1
		if i <= n-3 && nums[i+2]-nums[i+1] == d {
			rightLen = suff[i+1]
		}

		total := leftLen + 1 + rightLen
		if total > ans {
			ans = total
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(LongestArithmeticSequenceAfterChangingAtMostOneElement([]int{9, 7, 5, 10, 1})) // Expected: 5

	// Example 2
	fmt.Println(LongestArithmeticSequenceAfterChangingAtMostOneElement([]int{1, 2, 6, 7})) // Expected: 3

	// Extra
	fmt.Println(LongestArithmeticSequenceAfterChangingAtMostOneElement([]int{1, 2, 3, 4})) // Expected: 4
	fmt.Println(LongestArithmeticSequenceAfterChangingAtMostOneElement([]int{1}))          // Expected: 1
}
```
