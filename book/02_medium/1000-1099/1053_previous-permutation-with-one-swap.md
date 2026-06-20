# 1053 — Previous Permutation With One Swap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func prevPermOpt1(arr []int) []int
```

> **💡 Hint:** Find rightmost pair where arr[i] > arr[i+1].

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1053: Previous Permutation With One Swap
// https://leetcode.com/problems/previous-permutation-with-one-swap/
// Difficulty: Medium
//
// Approach: Find rightmost pair where arr[i] > arr[i+1].
//           Swap with the largest smaller element to the right.
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(prevPermOpt1([]int{3, 2, 1}))    // [3,1,2]
	fmt.Println(prevPermOpt1([]int{1, 1, 5}))    // [1,1,5]
	fmt.Println(prevPermOpt1([]int{1, 9, 4, 6, 7})) // [1,7,4,6,9]
}

func prevPermOpt1(arr []int) []int {
	i := len(arr) - 2
	for i >= 0 && arr[i] <= arr[i+1] {
		i--
	}

	if i < 0 {
		return arr
	}

	// Find rightmost smaller than arr[i]
	j := len(arr) - 1
	for arr[j] >= arr[i] {
		j--
	}
	// Skip duplicates
	for arr[j] == arr[j-1] {
		j--
	}

	arr[i], arr[j] = arr[j], arr[i]
	return arr
}
```
