# 3048 — Earliest Second To Mark Indices I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func earliestSecondToMarkIndices(nums []int, changeIndices []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m log m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3048: Earliest Second to Mark Indices I
// https://leetcode.com/problems/earliest-second-to-mark-indices-i/
// Difficulty: Medium
// Time: O(m log m) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(earliestSecondToMarkIndices([]int{2, 2, 0}, []int{2, 2, 2, 2, 3, 2, 2, 1}))
	fmt.Println(earliestSecondToMarkIndices([]int{1, 3}, []int{1, 1, 1, 2, 1, 1, 1}))
	fmt.Println(earliestSecondToMarkIndices([]int{0, 1}, []int{2, 2, 2}))
}

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	ans := sort.Search(m+1, func(t int) bool {
		if t == 0 {
			return false
		}
  // Alokasi slice integer
		last := make([]int, n+1)
		for s, idx := range changeIndices[:t] {
			last[idx] = s
		}
		for i := 1; i <= n; i++ {
			if last[i] == 0 && i != changeIndices[0] {
				return false
			}
		}
		decrement := 0
		marked := 0
		for s, idx := range changeIndices[:t] {
			if last[idx] == s {
				if decrement < nums[idx-1] {
					return false
				}
				decrement -= nums[idx-1]
				marked++
			} else {
				decrement++
			}
		}
		return marked == n
	})
	if ans > m {
		return -1
	}
	return ans
}
```
