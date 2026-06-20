# 0496 — Next Greater Element I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NextGreaterElementI(nums1, nums2 []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Stack

**Kompleksitas Waktu:** O(n+m), Space: O(m)  
**Kompleksitas Ruang:** O(m)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #496: Next Greater Element I
// https://leetcode.com/problems/next-greater-element-i/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(m)
func NextGreaterElementI(nums1, nums2 []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	nextGreater := make(map[int]int)
	var stack []int
	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			nextGreater[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
  // Alokasi slice integer
	result := make([]int, len(nums1))
	for i, v := range nums1 {
		if val, ok := nextGreater[v]; ok {
			result[i] = val
		} else {
			result[i] = -1
		}
	}
	return result
}

func main() {
	fmt.Println(NextGreaterElementI([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(NextGreaterElementI([]int{2, 4}, []int{1, 2, 3, 4}))
}
```
