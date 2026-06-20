# 3755 — Find Maximum Balanced Xor Subarray Length

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findMaximumBalancedXorSubarrayLength(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3755: Find Maximum Balanced XOR Subarray Length
// https://leetcode.com/problems/find-maximum-balanced-xor-subarray-length/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findMaximumBalancedXorSubarrayLength(nums []int) int {
	type state struct {
		xor  int
		diff int
	}
  // Membuat map (HashMap) — pencarian O(1)
	first := make(map[state]int)
	// Initial state before first element
	first[state{xor: 0, diff: 0}] = -1

	prefixXor := 0
	diff := 0
	maxLen := 0

	for i, v := range nums {
		prefixXor ^= v
		if v%2 == 1 {
			diff++
		} else {
			diff--
		}

		key := state{xor: prefixXor, diff: diff}
		if pos, ok := first[key]; ok {
			if i-pos > maxLen {
				maxLen = i - pos
			}
		} else {
			first[key] = i
		}
	}

	return maxLen
}

func main() {
	fmt.Println(findMaximumBalancedXorSubarrayLength([]int{3, 1, 3, 2, 0}))
	fmt.Println(findMaximumBalancedXorSubarrayLength([]int{3, 2, 8, 5, 4, 14, 9, 15}))
	fmt.Println(findMaximumBalancedXorSubarrayLength([]int{1, 2, 3, 4, 5}))
}
```
