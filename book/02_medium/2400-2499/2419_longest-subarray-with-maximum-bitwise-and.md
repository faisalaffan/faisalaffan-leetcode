# 2419 — Longest Subarray With Maximum Bitwise And

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestSubarray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2419: Longest Subarray With Maximum Bitwise AND
// https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Max AND in any subarray is just the max element. Find longest consecutive
// subarray where all elements equal the max element.

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 2, 3, 3, 2, 2})) // 2
	fmt.Println(longestSubarray([]int{1, 2, 3, 4}))        // 1
}

func longestSubarray(nums []int) int {
	mx := 0
	for _, v := range nums {
		if v > mx {
			mx = v
		}
	}
	ans, cur := 0, 0
	for _, v := range nums {
		if v == mx {
			cur++
		} else {
			cur = 0
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
```
