# 1535 — Find The Winner Of An Array Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func GetWinner(arr []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1535: Find the Winner of an Array Game
// https://leetcode.com/problems/find-the-winner-of-an-array-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))
	fmt.Println(GetWinner([]int{3, 2, 1}, 10))
	fmt.Println(GetWinner([]int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}, 1000000000))
}

func GetWinner(arr []int, k int) int {
	// Time: O(N), Space: O(1)
	if k == 0 {
		return 0
	}

	// If k >= n, the maximum element wins
	n := len(arr)
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}
	if k >= n {
		return maxVal
	}

	current := arr[0]
	wins := 0

	for i := 1; i < n; i++ {
		if current > arr[i] {
			wins++
		} else {
			current = arr[i]
			wins = 1
		}

		if wins == k {
			return current
		}
	}

	// If we've gone through the whole array, the max element wins
	return maxVal
}
```
