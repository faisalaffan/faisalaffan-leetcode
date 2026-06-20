# 2404 — Most Frequent Even Element

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MostFrequentEvenElement(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2404: Most Frequent Even Element
// https://leetcode.com/problems/most-frequent-even-element/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(MostFrequentEvenElement([]int{0, 1, 2, 2, 4, 4, 1})) // 2
	fmt.Println(MostFrequentEvenElement([]int{4, 4, 4, 9, 2, 4}))    // 4
	fmt.Println(MostFrequentEvenElement([]int{1, 3, 5, 7}))           // -1
}

func MostFrequentEvenElement(nums []int) int {
	freq := map[int]int{}
	for _, n := range nums {
		if n%2 == 0 {
			freq[n]++
		}
	}
	if len(freq) == 0 {
		return -1
	}
	bestNum := -1
	bestCount := 0
	for n, c := range freq {
		if c > bestCount || (c == bestCount && n < bestNum) {
			bestNum = n
			bestCount = c
		}
	}
	return bestNum
}
```
