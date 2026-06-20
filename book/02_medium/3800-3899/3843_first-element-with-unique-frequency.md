# 3843 — First Element With Unique Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FirstElementWithUniqueFrequency(nums []int) int
```

> **💡 Hint:** Count frequencies, then count frequency-of-frequency, find first with freq=1.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3843: First Element with Unique Frequency
// https://leetcode.com/problems/first-element-with-unique-frequency/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Count frequencies, then count frequency-of-frequency, find first with freq=1.

import "fmt"

func FirstElementWithUniqueFrequency(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}

  // Membuat map (HashMap) — pencarian O(1)
	freqCnt := make(map[int]int)
	for _, c := range cnt {
		freqCnt[c]++
	}

	for _, v := range nums {
		if freqCnt[cnt[v]] == 1 {
			return v
		}
	}

	return -1
}

func main() {
	// Example 1
	fmt.Println(FirstElementWithUniqueFrequency([]int{20, 10, 30, 30})) // Expected: 30

	// Example 2
	fmt.Println(FirstElementWithUniqueFrequency([]int{20, 20, 10, 30, 30, 30})) // Expected: 20

	// Example 3
	fmt.Println(FirstElementWithUniqueFrequency([]int{10, 10, 20, 20})) // Expected: -1
}
```
