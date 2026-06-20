# 3005 — Count Elements With Maximum Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountElementsWithMaximumFrequency(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3005: Count Elements With Maximum Frequency
// https://leetcode.com/problems/count-elements-with-maximum-frequency/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maxFrequencyElements
	fmt.Println(CountElementsWithMaximumFrequency([]int{1, 2, 2, 3, 1, 4})) // 4
	fmt.Println(CountElementsWithMaximumFrequency([]int{1, 2, 3, 4, 5}))    // 5
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: maxFrequencyElements
func CountElementsWithMaximumFrequency(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	maxFreq := 0
	for _, v := range nums {
		freq[v]++
		if freq[v] > maxFreq {
			maxFreq = freq[v]
		}
	}
	total := 0
	for _, f := range freq {
		if f == maxFreq {
			total += f
		}
	}
	return total
}
```
