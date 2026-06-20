# 2422 — Merge Operations To Turn Array Into A Palindrome

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minMerges(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2422: Merge Operations to Turn Array Into a Palindrome
// https://leetcode.com/problems/merge-operations-to-turn-array-into-a-palindrome/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Two pointers, merge smaller side towards larger.

import "fmt"

func main() {
	fmt.Println(minMerges([]int{1, 4, 1, 3}))        // 1 (merge 4+1 to 5, then [1,5,3])
	fmt.Println(minMerges([]int{1, 2, 3, 4, 5, 1})) // 2
}

func minMerges(nums []int) int {
	ops := 0
	i, j := 0, len(nums)-1
	left, right := nums[i], nums[j]

	for i < j {
		if left < right {
			i++
			left += nums[i]
			ops++
		} else if left > right {
			j--
			right += nums[j]
			ops++
		} else {
			i++
			j--
			if i < j {
				left, right = nums[i], nums[j]
			}
		}
	}
	return ops
}
```
