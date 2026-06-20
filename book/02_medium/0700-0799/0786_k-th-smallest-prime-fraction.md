# 0786 — K Th Smallest Prime Fraction

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func kthSmallestPrimeFraction(arr []int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n log max)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #786: K-th Smallest Prime Fraction
// https://leetcode.com/problems/k-th-smallest-prime-fraction/
// Difficulty: Medium
// Time: O(n log max)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
	fmt.Println(kthSmallestPrimeFraction([]int{1, 7}, 1))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	left, right := 0.0, 1.0

  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := (left + right) / 2.0
		count := 0
		maxFraction := 0.0
		p, q := 0, 1

		j := 1
		for i := 0; i < n; i++ {
			for j < n && float64(arr[i])/float64(arr[j]) > mid {
				j++
			}
			if j == n {
				break
			}
			count += n - j

			fraction := float64(arr[i]) / float64(arr[j])
			if fraction > maxFraction {
				maxFraction = fraction
				p, q = arr[i], arr[j]
			}
		}

		if count == k {
			return []int{p, q}
		} else if count < k {
			left = mid
		} else {
			right = mid
		}
	}

	return nil
}
```
