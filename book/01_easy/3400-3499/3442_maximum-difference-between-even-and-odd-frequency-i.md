# 3442 — Maximum Difference Between Even And Odd Frequency I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumDifferenceBetweenEvenAndOddFrequencyI(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3442: Maximum Difference Between Even and Odd Frequency I
// https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumDifferenceBetweenEvenAndOddFrequencyI("aaaaabbc"))
	fmt.Println(MaximumDifferenceBetweenEvenAndOddFrequencyI("abcabcab"))
}

// MaximumDifferenceBetweenEvenAndOddFrequencyI returns the max difference between max even-frequency and min odd-frequency in s.
// Time: O(n). Space: O(1).
func MaximumDifferenceBetweenEvenAndOddFrequencyI(s string) int {
  // Alokasi slice integer
	freq := make([]int, 26)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	maxEven := 0
	minOdd := -1
	for _, f := range freq {
		if f == 0 {
			continue
		}
		if f%2 == 0 && f > maxEven {
			maxEven = f
		} else if f%2 == 1 && (minOdd == -1 || f < minOdd) {
			minOdd = f
		}
	}

	if maxEven == 0 || minOdd == -1 {
		return 0
	}
	return maxEven - minOdd
}
```
