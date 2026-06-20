# 3109 — Find The Index Of Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func getPermutationIndex(perm []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3109: Find the Index of Permutation
// https://leetcode.com/problems/find-the-index-of-permutation/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n)

import "fmt"

func getPermutationIndex(perm []int) int {
	n := len(perm)
	mod := int64(1000000007)

  // Alokasi slice integer
	fact := make([]int64, n)
	fact[0] = 1
	for i := 1; i < n; i++ {
		fact[i] = fact[i-1] * int64(i) % mod
	}

  // Alokasi slice integer
	bit := make([]int, n+1)

	update := func(idx, val int) {
		for idx <= n {
			bit[idx] += val
			idx += idx & -idx
		}
	}

	query := func(idx int) int {
		sum := 0
		for idx > 0 {
			sum += bit[idx]
			idx -= idx & -idx
		}
		return sum
	}

	for i := 1; i <= n; i++ {
		update(i, 1)
	}

	ans := int64(0)
	for i := 0; i < n; i++ {
		smaller := query(perm[i]) - 1
		ans = (ans + int64(smaller)*fact[n-1-i]) % mod
		update(perm[i], -1)
	}

	return int((ans + 1) % mod)
}

func main() {
	fmt.Println(getPermutationIndex([]int{1, 2, 3})) // Expected: 1
	fmt.Println(getPermutationIndex([]int{3, 2, 1})) // Expected: 6
	fmt.Println(getPermutationIndex([]int{2, 1, 3})) // Expected: 3
}
```
