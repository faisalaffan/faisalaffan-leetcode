# 0898 — Bitwise Ors Of Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func BitwiseOrsOfSubarrays(arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * log(max))  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #898: Bitwise ORs of Subarrays
// https://leetcode.com/problems/bitwise-ors-of-subarrays/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(BitwiseOrsOfSubarrays([]int{0}))
	fmt.Println(BitwiseOrsOfSubarrays([]int{1, 1, 2}))
	fmt.Println(BitwiseOrsOfSubarrays([]int{1, 2, 4}))
}

// Time: O(n * log(max)) | Space: O(n)
func BitwiseOrsOfSubarrays(arr []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	set := make(map[int]bool)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(arr); i++ {
		set[arr[i]] = true
		for j := i - 1; j >= 0; j-- {
			if arr[i]|arr[j] == arr[j] {
				break
			}
			arr[j] |= arr[i]
			set[arr[j]] = true
		}
	}

	return len(set)
}
```
