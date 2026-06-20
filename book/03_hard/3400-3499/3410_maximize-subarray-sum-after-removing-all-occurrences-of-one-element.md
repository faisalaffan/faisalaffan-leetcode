# 3410 — Maximize Subarray Sum After Removing All Occurrences Of One Element

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximizeSubarraySumAfterRemovingAllOccurrencesOfOneElement(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3410: Maximize Subarray Sum After Removing All Occurrences of One Element
// https://leetcode.com/problems/maximize-subarray-sum-after-removing-all-occurrences-of-one-element/
// Difficulty: Hard
//
// Kadane variant: skip all occurrences of one value.
// For each distinct value, run Kadane treating that value as 0.
// Also compute standard Kadane (no removal).

import "fmt"

func main() {
	fmt.Println(MaximizeSubarraySumAfterRemovingAllOccurrencesOfOneElement([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func MaximizeSubarraySumAfterRemovingAllOccurrencesOfOneElement(nums []int) int64 {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

	// Standard Kadane (no removal)
	best := int64(nums[0])
	cur := int64(0)
	for _, v := range nums {
		cur += int64(v)
		if cur > best {
			best = cur
		}
		if cur < 0 {
			cur = 0
		}
	}

	// Try removing each distinct value
  // Membuat map (HashMap) — pencarian O(1)
	vals := make(map[int]bool)
	for _, v := range nums {
		vals[v] = true
	}

	for skipVal := range vals {
		cur = 0
		for _, v := range nums {
			if v == skipVal {
				continue
			}
			cur += int64(v)
			if cur > best {
				best = cur
			}
			if cur < 0 {
				cur = 0
			}
		}
	}
	return best
}
```
