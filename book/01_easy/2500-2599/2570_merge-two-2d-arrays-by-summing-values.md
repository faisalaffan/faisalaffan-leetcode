# 2570 — Merge Two 2D Arrays By Summing Values

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func MergeTwoTwoDArraysBySummingValues(nums1 [][]int, nums2 [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2570: Merge Two 2D Arrays by Summing Values
// https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/
// Difficulty: Easy
// Time O(n + m) | Space O(n + m)

import "fmt"

func main() {
	fmt.Println(MergeTwoTwoDArraysBySummingValues([][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 4}, {3, 2}, {4, 1}})) // [[1,6],[2,3],[3,2],[4,6]]
	fmt.Println(MergeTwoTwoDArraysBySummingValues([][]int{{2, 4}, {3, 6}, {5, 5}}, [][]int{{1, 3}, {4, 3}}))          // [[1,3],[2,4],[3,6],[4,3],[5,5]]
}

func MergeTwoTwoDArraysBySummingValues(nums1 [][]int, nums2 [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i][0] == nums2[j][0] {
			res = append(res, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		} else if nums1[i][0] < nums2[j][0] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	for i < len(nums1) {
		res = append(res, nums1[i])
		i++
	}
	for j < len(nums2) {
		res = append(res, nums2[j])
		j++
	}
	return res
}
```
