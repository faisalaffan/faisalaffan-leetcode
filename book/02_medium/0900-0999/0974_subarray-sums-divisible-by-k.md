# 0974 — Subarray Sums Divisible By K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func subarraysDivByK(nums []int, k int) int
```

> **💡 Hint:** Prefix sum + modulo counting

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(k)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #974: Subarray Sums Divisible by K
// https://leetcode.com/problems/subarray-sums-divisible-by-k/
// Difficulty: Medium
//
// Approach: Prefix sum + modulo counting
// Time: O(n)
// Space: O(k)

import "fmt"

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5)) // 7
	fmt.Println(subarraysDivByK([]int{5}, 9))                  // 0
	fmt.Println(subarraysDivByK([]int{-1, 2, 9}, 2))           // 2
}

func subarraysDivByK(nums []int, k int) int {
  // Membuat map (HashMap) — pencarian O(1)
	modCount := make(map[int]int)
	modCount[0] = 1
	prefixSum := 0
	result := 0

	for _, n := range nums {
		prefixSum += n
		mod := prefixSum % k
		if mod < 0 {
			mod += k
		}
		result += modCount[mod]
		modCount[mod]++
	}

	return result
}
```
