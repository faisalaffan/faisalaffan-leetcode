# 1224 — Maximum Equal Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxEqualFreq(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1224: Maximum Equal Frequency
// https://leetcode.com/problems/maximum-equal-frequency/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("1224. Maximum Equal Frequency")
	fmt.Println("[2,2,1,1,5,3,3,5]:", maxEqualFreq([]int{2, 2, 1, 1, 5, 3, 3, 5}), "(expected 7)")
	fmt.Println("[1,1,1,2,2,2]:", maxEqualFreq([]int{1, 1, 1, 2, 2, 2}), "(expected 5)")
	fmt.Println("[1,2]:", maxEqualFreq([]int{1, 2}), "(expected 2)")
}

func maxEqualFreq(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int) // num -> freq
  // Membuat map (HashMap) — pencarian O(1)
	cnt := make(map[int]int)  // freq -> count of numbers with that freq
	result := 0

	for i, num := range nums {
		oldF := freq[num]
		newF := oldF + 1
		freq[num] = newF

		if oldF > 0 {
			cnt[oldF]--
			if cnt[oldF] == 0 {
				delete(cnt, oldF)
			}
		}
		cnt[newF]++

		if len(cnt) == 1 {
			// All numbers have the same frequency.
			// Valid if frequency is 1 (remove any one element) or
			// only one number exists (count of that freq == 1).
			for f, c := range cnt {
				if f == 1 || c == 1 {
					result = i + 1
				}
			}
		} else if len(cnt) == 2 {
			// Two distinct frequencies. Valid cases:
			// 1. One freq = other+1, and count of higher freq is 1
			// 2. One freq is 1, and count of freq 1 is 1
			var f1, c1, f2, c2 int
			first := true
			for f, c := range cnt {
				if first {
					f1, c1 = f, c
					first = false
				} else {
					f2, c2 = f, c
				}
			}
			if f1 > f2 {
				f1, f2 = f2, f1
				c1, c2 = c2, c1
			}
			// f1 < f2
			if (f2 == f1+1 && c2 == 1) || (f1 == 1 && c1 == 1) {
				result = i + 1
			}
		}
	}

	return result
}
```
