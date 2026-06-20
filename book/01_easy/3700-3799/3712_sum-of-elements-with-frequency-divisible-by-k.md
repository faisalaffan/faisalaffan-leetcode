# 3712 — Sum Of Elements With Frequency Divisible By K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SumOfElementsWithFrequencyDivisibleByK(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3712: Sum of Elements With Frequency Divisible by K
// https://leetcode.com/problems/sum-of-elements-with-frequency-divisible-by-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SumOfElementsWithFrequencyDivisibleByK([]int{1, 2, 2, 3, 3, 3, 3, 4}, 2))
	fmt.Println(SumOfElementsWithFrequencyDivisibleByK([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(SumOfElementsWithFrequencyDivisibleByK([]int{4, 4, 4, 1, 2, 3}, 3))
}

// Time: O(n)
// Space: O(1)
func SumOfElementsWithFrequencyDivisibleByK(nums []int, k int) int {
	cnt := [101]int{}
	for _, v := range nums {
		cnt[v]++
	}

	sum := 0
	for v := 1; v <= 100; v++ {
		if cnt[v]%k == 0 {
			sum += v * cnt[v]
		}
	}
	return sum
}
```
