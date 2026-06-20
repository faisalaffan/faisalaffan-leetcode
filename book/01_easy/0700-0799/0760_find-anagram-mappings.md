# 0760 — Find Anagram Mappings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func anagramMappings(nums1 []int, nums2 []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #760: Find Anagram Mappings
// https://leetcode.com/problems/find-anagram-mappings/
// Difficulty: Easy [Paid]
// Note: This is a premium problem. Implementation based on public description.

import "fmt"

func main() {
	fmt.Println(anagramMappings([]int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28})) // [1,4,3,2,0]
	fmt.Println(anagramMappings([]int{1, 2}, []int{2, 1}))                              // [1,0]
}

// anagramMappings returns a mapping array P where P[i] is the index of A[i] in B.
// Time: O(n). Space: O(n).
func anagramMappings(nums1 []int, nums2 []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[int]int)
	for i, v := range nums2 {
		pos[v] = i
	}
  // Alokasi slice integer
	result := make([]int, len(nums1))
	for i, v := range nums1 {
		result[i] = pos[v]
	}
	return result
}
```
