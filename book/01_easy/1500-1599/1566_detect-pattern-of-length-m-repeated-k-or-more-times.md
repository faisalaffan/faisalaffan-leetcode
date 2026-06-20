# 1566 — Detect Pattern Of Length M Repeated K Or More Times

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func containsPattern(arr []int, m int, k int) bool

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1566: Detect Pattern of Length M Repeated K or More Times
// https://leetcode.com/problems/detect-pattern-of-length-m-repeated-k-or-more-times/
// Difficulty: Easy
//
// LeetCode submission: func containsPattern(arr []int, m int, k int) bool

import "fmt"

func main() {
	fmt.Println(DetectPatternOfLengthMRepeatedKOrMoreTimes([]int{1, 2, 4, 4, 4, 4}, 1, 3)) // true
	fmt.Println(DetectPatternOfLengthMRepeatedKOrMoreTimes([]int{1, 2, 1, 2, 1, 1, 1, 3}, 2, 2)) // true
	fmt.Println(DetectPatternOfLengthMRepeatedKOrMoreTimes([]int{1, 2, 1, 2, 1, 3}, 2, 3)) // false
}

// Time: O(n * m), Space: O(1)
func DetectPatternOfLengthMRepeatedKOrMoreTimes(arr []int, m int, k int) bool {
	n := len(arr)
	if m*k > n {
		return false
	}
	for i := 0; i <= n-m*k; i++ {
		count := 1
		for j := i + m; j <= n-m; j += m {
			match := true
			for t := 0; t < m; t++ {
				if arr[j+t] != arr[i+t] {
					match = false
					break
				}
			}
			if match {
				count++
				if count >= k {
					return true
				}
			} else {
				break
			}
		}
	}
	return false
}
```
