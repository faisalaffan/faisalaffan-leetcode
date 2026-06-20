# 3676 — Count Bowl Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countBowlSubarrays(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3676: Count Bowl Subarrays
// https://leetcode.com/problems/count-bowl-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func countBowlSubarrays(nums []int) int64 {
	var ans int64 = 0
	var stack []int

	for _, num := range nums {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			if len(stack) >= 2 {
				ans++
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	return ans
}

func main() {
	fmt.Println(countBowlSubarrays([]int{1, 3, 5, 4, 2}))
	fmt.Println(countBowlSubarrays([]int{3, 1, 2, 4}))
	fmt.Println(countBowlSubarrays([]int{1, 2, 3, 4}))
}
```
