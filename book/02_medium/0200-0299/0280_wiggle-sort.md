# 0280 — Wiggle Sort

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func wiggleSort(nums []int) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #280: Wiggle Sort
// https://leetcode.com/problems/wiggle-sort/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func wiggleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if i%2 == 1 {
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		} else {
			if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}
}

func main() {
	nums1 := []int{3, 5, 2, 1, 6, 4}
	wiggleSort(nums1)
	fmt.Println(nums1)

	nums2 := []int{1, 2, 3, 4}
	wiggleSort(nums2)
	fmt.Println(nums2)

	nums3 := []int{1}
	wiggleSort(nums3)
	fmt.Println(nums3)
}
```
