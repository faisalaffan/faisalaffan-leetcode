# 0220 — Contains Duplicate Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #220: Contains Duplicate III
// https://leetcode.com/problems/contains-duplicate-iii/
// Difficulty: Hard

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	if valueDiff < 0 || indexDiff <= 0 {
		return false
	}

  // HashMap: O(1) lookup
	buckets := make(map[int]int)

	for i, num := range nums {
		bucketID := num / (valueDiff + 1)
		if num < 0 {
			bucketID--
		}

		if _, exists := buckets[bucketID]; exists {
			return true
		}
		if val, exists := buckets[bucketID-1]; exists && num-val <= valueDiff {
			return true
		}
		if val, exists := buckets[bucketID+1]; exists && val-num <= valueDiff {
			return true
		}

		buckets[bucketID] = num

		if i >= indexDiff {
			oldNum := nums[i-indexDiff]
			oldBucket := oldNum / (valueDiff + 1)
			if oldNum < 0 {
				oldBucket--
			}
			delete(buckets, oldBucket)
		}
	}

	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
}
```
