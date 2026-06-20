# 1243 — Array Transformation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func transformArray(arr []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2) worst case  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1243: Array Transformation
// https://leetcode.com/problems/array-transformation/
// Difficulty: Easy [Paid]
// Time: O(n^2) worst case | Space: O(n)

import "fmt"

func main() {
	fmt.Println(transformArray([]int{6, 2, 3, 4})) // [6,3,3,4]
	fmt.Println(transformArray([]int{1, 6, 3, 4, 3, 5})) // [1,4,4,4,4,5]
}

// LeetCode submission: transformArray
func transformArray(arr []int) []int {
	if len(arr) <= 2 {
		return append([]int{}, arr...)
	}
	for {
		changed := false
  // Alokasi slice integer
		next := make([]int, len(arr))
		copy(next, arr)
		for i := 1; i < len(arr)-1; i++ {
			if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
				next[i]++
				changed = true
			} else if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				next[i]--
				changed = true
			}
		}
		if !changed {
			break
		}
		arr = next
	}
	return arr
}
```
