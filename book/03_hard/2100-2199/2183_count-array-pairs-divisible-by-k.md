# 2183 — Count Array Pairs Divisible By K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countPairs(nums []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, GCD / Matematika

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2183: Count Array Pairs Divisible by K
// https://leetcode.com/problems/count-array-pairs-divisible-by-k/
// Difficulty: Hard
//
// For each element, compute g = gcd(num, k). Two numbers pair to make a
// product divisible by k iff (g1 * g2) % k == 0. Count frequency of each
// gcd value, then iterate over all pairs of gcd values.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countPairs([]int{1, 2, 3, 4, 5}, 2))

	// Example 2
	fmt.Println(countPairs([]int{1, 2, 3, 4}, 5))

	// Example 3
	fmt.Println(countPairs([]int{1, 2, 3, 4, 5}, 6))

	// All divisible
	fmt.Println(countPairs([]int{10, 20, 30}, 5))
}

func countPairs(nums []int, k int) int64 {
	// Count frequency of each gcd value
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, num := range nums {
		g := gcd(num, k)
		freq[g]++
	}

	var ans int64

	// Iterate over unique gcd values
  // Alokasi slice integer
	gcdVals := make([]int, 0, len(freq))
	for g := range freq {
		gcdVals = append(gcdVals, g)
	}

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(gcdVals); i++ {
		for j := i; j < len(gcdVals); j++ {
			g1, g2 := gcdVals[i], gcdVals[j]
			if (int64(g1)*int64(g2))%int64(k) == 0 {
				if i == j {
					// Same gcd: count C(freq, 2)
					f := int64(freq[g1])
					ans += f * (f - 1) / 2
				} else {
					ans += int64(freq[g1]) * int64(freq[g2])
				}
			}
		}
	}

	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
