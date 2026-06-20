# 1437 — Check If All 1S Are At Least Length K Places Away

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func kLengthApart(nums []int, k int) bool

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

// LeetCode #1437: Check If All 1's Are at Least Length K Places Away
// https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/
// Difficulty: Easy
//
// LeetCode submission: func kLengthApart(nums []int, k int) bool

import "fmt"

func main() {
	fmt.Println(CheckIfAllOneSAreAtLeastLengthKPlacesAway([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2)) // true
	fmt.Println(CheckIfAllOneSAreAtLeastLengthKPlacesAway([]int{1, 0, 0, 1, 0, 1}, 2))       // false
}

// Time: O(n), Space: O(1)
func CheckIfAllOneSAreAtLeastLengthKPlacesAway(nums []int, k int) bool {
	prev := -k - 1
	for i, v := range nums {
		if v == 1 {
			if i-prev-1 < k {
				return false
			}
			prev = i
		}
	}
	return true
}
```
