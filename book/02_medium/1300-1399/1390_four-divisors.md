# 1390 — Four Divisors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func sumFourDivisors(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * sqrt(m)) where n = len(nums), m = max value in nums  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1390: Four Divisors
// https://leetcode.com/problems/four-divisors/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(sumFourDivisors([]int{21, 4, 7})) // 32

	// Test case 2
	fmt.Println(sumFourDivisors([]int{21, 21})) // 64

	// Test case 3
	fmt.Println(sumFourDivisors([]int{1, 2, 3, 4, 5})) // 0
}

// Time: O(n * sqrt(m)) where n = len(nums), m = max value in nums
// Space: O(1)
func sumFourDivisors(nums []int) int {
	total := 0

	for _, num := range nums {
		divCount := 0
		divSum := 0

		for i := 1; i*i <= num; i++ {
			if num%i == 0 {
				divCount++
				divSum += i

				if i*i != num {
					divCount++
					divSum += num / i
				}
			}
			if divCount > 4 {
				break
			}
		}

		if divCount == 4 {
			total += divSum
		}
	}

	return total
}
```
