# 2865 — Beautiful Towers I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func BeautifulTowersI(maxHeights []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n^2)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2865: Beautiful Towers I
// https://leetcode.com/problems/beautiful-towers-i/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func BeautifulTowersI(maxHeights []int) int64 {
	n := len(maxHeights)
	var best int64

	for peak := 0; peak < n; peak++ {
		var total int64 = int64(maxHeights[peak])
		prev := maxHeights[peak]

		// Left side
		for i := peak - 1; i >= 0; i-- {
			h := maxHeights[i]
			if h > prev {
				h = prev
			}
			total += int64(h)
			prev = h
		}

		prev = maxHeights[peak]
		// Right side
		for i := peak + 1; i < n; i++ {
			h := maxHeights[i]
			if h > prev {
				h = prev
			}
			total += int64(h)
			prev = h
		}

		if total > best {
			best = total
		}
	}

	return best
}

func main() {
	fmt.Println(BeautifulTowersI([]int{5, 3, 4, 1, 1}))
	fmt.Println(BeautifulTowersI([]int{6, 5, 3, 9, 2, 7}))
}
```
