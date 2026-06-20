# 0220 — Contains Duplicate Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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

  // Membuat map (HashMap) — pencarian O(1)
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
