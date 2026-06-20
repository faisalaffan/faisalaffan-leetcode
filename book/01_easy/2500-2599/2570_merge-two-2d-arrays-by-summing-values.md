# 2570 — Merge Two 2D Arrays By Summing Values

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MergeTwoTwoDArraysBySummingValues(nums1 [][]int, nums2 [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
