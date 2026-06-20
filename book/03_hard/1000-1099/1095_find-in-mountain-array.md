# 1095 — Find In Mountain Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findInMountainArray(target int, mountainArr *MountainArray) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1095: Find in Mountain Array
// https://leetcode.com/problems/find-in-mountain-array/
// Difficulty: Hard
//
// Binary search for the peak, then binary search each side (ascending then
// descending). Return the minimum index where target is found.

import "fmt"

// MountainArray provides read-only access to a mountain array.
// The get(i) and length() methods have at most 100 calls total.

type MountainArray struct {
	arr []int
}

func (ma *MountainArray) get(index int) int { return ma.arr[index] }
func (ma *MountainArray) length() int       { return len(ma.arr) }

func main() {
	ma := &MountainArray{arr: []int{1, 2, 3, 4, 5, 3, 1}}
	fmt.Println(findInMountainArray(3, ma))
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()

	// Find the peak index
	lo, hi := 0, n-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	peak := lo

	// Search left ascending side (inclusive of peak)
	if idx := binarySearchAsc(mountainArr, target, 0, peak); idx != -1 {
		return idx
	}

	// Search right descending side
	return binarySearchDesc(mountainArr, target, peak+1, n-1)
}

func binarySearchAsc(ma *MountainArray, target, lo, hi int) int {
	for lo <= hi {
		mid := lo + (hi-lo)/2
		val := ma.get(mid)
		if val == target {
			return mid
		} else if val < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

func binarySearchDesc(ma *MountainArray, target, lo, hi int) int {
	for lo <= hi {
		mid := lo + (hi-lo)/2
		val := ma.get(mid)
		if val == target {
			return mid
		} else if val > target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
```
