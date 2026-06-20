# 2511 — Maximum Enemy Forts That Can Be Captured

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximumEnemyFortsThatCanBeCaptured(forts []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2511: Maximum Enemy Forts That Can Be Captured
// https://leetcode.com/problems/maximum-enemy-forts-that-can-be-captured/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MaximumEnemyFortsThatCanBeCaptured([]int{1, 0, 0, -1, 0, 0, 0, 0, 1})) // 4
	fmt.Println(MaximumEnemyFortsThatCanBeCaptured([]int{0, 0, 1, -1}))                 // 0
}

func MaximumEnemyFortsThatCanBeCaptured(forts []int) int {
	maxCap := 0
  // Linear scan O(n)
	for i := 0; i < len(forts); i++ {
		if forts[i] == 1 {
			// Move right
			for j := i + 1; j < len(forts); j++ {
				if forts[j] == -1 {
					if j-i-1 > maxCap {
						maxCap = j - i - 1
					}
					break
				} else if forts[j] == 1 {
					break
				}
			}
		} else if forts[i] == -1 {
			// Move right (capturing from -1 to 1)
			for j := i + 1; j < len(forts); j++ {
				if forts[j] == 1 {
					if j-i-1 > maxCap {
						maxCap = j - i - 1
					}
					break
				} else if forts[j] == -1 {
					break
				}
			}
		}
	}
	return maxCap
}
```
