# 1299 — Replace Elements With Greatest Element On Right Side

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func replaceElements(arr []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) excluding output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1299: Replace Elements with Greatest Element on Right Side
// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/
// Difficulty: Easy
// Time: O(n) | Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1})) // [18,6,6,6,1,-1]
	fmt.Println(replaceElements([]int{400}))                 // [-1]
}

// LeetCode submission: replaceElements
func replaceElements(arr []int) []int {
  // Alokasi slice integer
	ans := make([]int, len(arr))
	maxRight := -1
	for i := len(arr) - 1; i >= 0; i-- {
		ans[i] = maxRight
		if arr[i] > maxRight {
			maxRight = arr[i]
		}
	}
	return ans
}
```
