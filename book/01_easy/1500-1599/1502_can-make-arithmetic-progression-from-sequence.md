# 1502 — Can Make Arithmetic Progression From Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func canMakeArithmeticProgression(arr []int) bool

import (
	"fmt"
	"sort"
)

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1502: Can Make Arithmetic Progression From Sequence
// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/
// Difficulty: Easy
//
// LeetCode submission: func canMakeArithmeticProgression(arr []int) bool

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CanMakeArithmeticProgressionFromSequence([]int{3, 5, 1})) // true
	fmt.Println(CanMakeArithmeticProgressionFromSequence([]int{1, 2, 4})) // false
}

// Time: O(n log n), Space: O(1)
func CanMakeArithmeticProgressionFromSequence(arr []int) bool {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}
```
