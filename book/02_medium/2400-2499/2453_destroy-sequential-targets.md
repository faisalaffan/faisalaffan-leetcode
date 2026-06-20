# 2453 — Destroy Sequential Targets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func destroyTargets(nums []int, space int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2453: Destroy Sequential Targets
// https://leetcode.com/problems/destroy-sequential-targets/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// Group nums by nums[i] % space. The group with max size gives max targets.
// Pick smallest nums[i] from that group.

import "fmt"

func main() {
	fmt.Println(destroyTargets([]int{3, 7, 8, 1, 1, 5}, 2)) // 1
	fmt.Println(destroyTargets([]int{1, 3, 5, 2, 4, 6}, 2)) // 1
}

func destroyTargets(nums []int, space int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	minVal := make(map[int]int)

	for _, v := range nums {
		rem := v % space
		freq[rem]++
		if _, ok := minVal[rem]; !ok || v < minVal[rem] {
			minVal[rem] = v
		}
	}

	maxFreq, ans := 0, 0
	for rem, f := range freq {
		if f > maxFreq || (f == maxFreq && minVal[rem] < ans) {
			maxFreq = f
			ans = minVal[rem]
		}
	}
	return ans
}
```
