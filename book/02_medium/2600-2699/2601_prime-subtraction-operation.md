# 2601 — Prime Subtraction Operation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func primeSubOperation(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * maxVal)  
**Kompleksitas Ruang:** O(maxVal)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2601: Prime Subtraction Operation
// https://leetcode.com/problems/prime-subtraction-operation/
// Difficulty: Medium
// Time: O(n * maxVal) | Space: O(maxVal)

import "fmt"

func primeSubOperation(nums []int) bool {
	// Sieve: find all primes up to 1000
	maxN := 1000
	isPrime := make([]bool, maxN+1)
	for i := 2; i <= maxN; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= maxN; i++ {
		if isPrime[i] {
			for j := i * i; j <= maxN; j += i {
				isPrime[j] = false
			}
		}
	}
	primes := []int{}
	for i := 2; i <= maxN; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	prev := 0
	for _, v := range nums {
		// Find largest prime p such that v-p > prev
		chosen := v
		for _, p := range primes {
			if p >= v {
				break
			}
			if v-p > prev {
				chosen = v - p
			} else {
				break
			}
		}
		if chosen <= prev {
			return false
		}
		prev = chosen
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", primeSubOperation([]int{4, 9, 6, 10}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", primeSubOperation([]int{6, 8, 11, 12}))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", primeSubOperation([]int{5, 8, 3}))
	// Expected: false
}
```
