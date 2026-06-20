# 1533 — Find The Index Of The Large Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func sumRange(arr []int, l, r int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(log N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1533: Find the Index of the Large Integer
// https://leetcode.com/problems/find-the-index-of-the-large-integer/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	// Problem: There is an array where one element is larger than the rest
	// (which are all equal). Find the index of the larger element.
	// We have an ArrayReader API that returns:
	//   -1 if sum(arr[l..r]) < sum(arr[subL..subR])
	//   0 if equal
	//   1 if sum(arr[l..r]) > sum(arr[subL..subR])

	arr := []int{1, 1, 1, 1, 1, 10, 1, 1, 1}
	reader := &arrayReader{arr: arr}
	idx := getIndex(reader, len(arr))
	fmt.Println("Index of larger element:", idx, "value:", arr[idx])

	// All equal
	arr2 := []int{2, 2, 2, 2, 2}
	reader2 := &arrayReader{arr: arr2}
	idx2 := getIndex(reader2, len(arr2))
	fmt.Println("Index (all equal):", idx2)
}

type arrayReader struct {
	arr []int
}

func (ar *arrayReader) compareSub(l, r, subL, subR int) int {
	s1 := sumRange(ar.arr, l, r)
	s2 := sumRange(ar.arr, subL, subR)
	if s1 < s2 {
		return -1
	} else if s1 > s2 {
		return 1
	}
	return 0
}

func sumRange(arr []int, l, r int) int {
	sum := 0
	for i := l; i <= r; i++ {
		sum += arr[i]
	}
	return sum
}

// getIndex uses the compareSub API to find the index of the largest element.
func getIndex(reader interface{ compareSub(int, int, int, int) int }, n int) int {
	// Time: O(log N), Space: O(1)
	left, right := 0, n-1

  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := left + (right-left)/2
		len1 := mid - left + 1
		len2 := right - mid

		if len1 == len2 {
			result := reader.compareSub(left, mid, mid+1, right)
			if result == 1 {
				right = mid
			} else if result == -1 {
				left = mid + 1
			} else {
				return -1 // all equal
			}
		} else {
			// len1 > len2 (since mid is floor)
			result := reader.compareSub(left, mid-1, mid+1, right)
			if result == 0 {
				return mid // the extra element is the larger one
			} else if result == 1 {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return left
}
```
