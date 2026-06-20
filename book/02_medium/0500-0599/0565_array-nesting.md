# 0565 — Array Nesting

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ArrayNesting(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (reuses input array as visited marker)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #565: Array Nesting
// https://leetcode.com/problems/array-nesting/
// Difficulty: Medium
// Time: O(n)
// Space: O(1) (reuses input array as visited marker)

import "fmt"

func main() {
	fmt.Println(ArrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))
	fmt.Println(ArrayNesting([]int{0, 1, 2}))
}

func ArrayNesting(nums []int) int {
	maxLen := 0
	visited := make([]bool, len(nums))

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		count := 0
		cur := i
		for !visited[cur] {
			visited[cur] = true
			cur = nums[cur]
			count++
		}
		if count > maxLen {
			maxLen = count
		}
	}

	return maxLen
}
```
