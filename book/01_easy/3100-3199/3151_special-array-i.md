# 3151 — Special Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SpecialArrayI(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3151: Special Array I
// https://leetcode.com/problems/special-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isArraySpecial
	fmt.Println(SpecialArrayI([]int{1}))       // true
	fmt.Println(SpecialArrayI([]int{2, 1, 4})) // true
	fmt.Println(SpecialArrayI([]int{4, 3, 1, 6})) // false
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: isArraySpecial
func SpecialArrayI(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if (nums[i]%2 == 0) == (nums[i-1]%2 == 0) {
			return false
		}
	}
	return true
}
```
