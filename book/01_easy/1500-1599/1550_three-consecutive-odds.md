# 1550 — Three Consecutive Odds

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func threeConsecutiveOdds(arr []int) bool

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1550: Three Consecutive Odds
// https://leetcode.com/problems/three-consecutive-odds/
// Difficulty: Easy
//
// LeetCode submission: func threeConsecutiveOdds(arr []int) bool

import "fmt"

func main() {
	fmt.Println(ThreeConsecutiveOdds([]int{2, 6, 4, 1}))  // false
	fmt.Println(ThreeConsecutiveOdds([]int{1, 2, 34, 3, 4, 5, 7, 23, 12})) // true
}

// Time: O(n), Space: O(1)
func ThreeConsecutiveOdds(arr []int) bool {
	count := 0
	for _, v := range arr {
		if v%2 == 1 {
			count++
			if count == 3 {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}
```
