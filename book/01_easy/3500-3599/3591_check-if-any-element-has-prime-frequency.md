# 3591 — Check If Any Element Has Prime Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfAnyElementHasPrimeFrequency(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + sqrt(m)) where m is max frequency  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3591: Check if Any Element Has Prime Frequency
// https://leetcode.com/problems/check-if-any-element-has-prime-frequency/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfAnyElementHasPrimeFrequency([]int{1, 2, 3, 4, 5, 4}))
	fmt.Println(CheckIfAnyElementHasPrimeFrequency([]int{1, 2, 3, 4, 5}))
	fmt.Println(CheckIfAnyElementHasPrimeFrequency([]int{2, 2, 2, 4, 4}))
}

// Time: O(n + sqrt(m)) where m is max frequency
// Space: O(1)
func CheckIfAnyElementHasPrimeFrequency(nums []int) bool {
	freq := [101]int{}
	for _, v := range nums {
		freq[v]++
	}

	for _, f := range freq {
		if f > 1 && isPrime(f) {
			return true
		}
	}
	return false
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
```
