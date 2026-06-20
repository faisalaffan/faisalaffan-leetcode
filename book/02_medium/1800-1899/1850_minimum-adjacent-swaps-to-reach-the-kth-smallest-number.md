# 1850 — Minimum Adjacent Swaps To Reach The Kth Smallest Number

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func GetMinSwaps(num string, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n*k + n^2), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1850: Minimum Adjacent Swaps to Reach the Kth Smallest Number
// https://leetcode.com/problems/minimum-adjacent-swaps-to-reach-the-kth-smallest-number/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetMinSwaps("5489355142", 4))
	fmt.Println(GetMinSwaps("11112", 4))
	fmt.Println(GetMinSwaps("00123", 1))
}

// Time: O(n*k + n^2), Space: O(n)
func GetMinSwaps(num string, k int) int {
	nums := []byte(num)
	n := len(nums)

	// Generate k-th next permutation
	var nextPerm func()
	nextPerm = func() {
		// Find longest non-increasing suffix
		i := n - 2
		for i >= 0 && nums[i] >= nums[i+1] {
			i--
		}
		if i < 0 {
			return
		}
		// Find rightmost element > nums[i]
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		// Reverse suffix
		left, right := i+1, n-1
  // Two-pointer: gerakkan kiri atau kanan
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	for i := 0; i < k; i++ {
		nextPerm()
	}
	target := string(nums)

	// Count minimum adjacent swaps from original to target
	original := []byte(num)
	swaps := 0
	for i := 0; i < n; i++ {
		if original[i] != target[i] {
			j := i
			for j < n && original[j] != target[i] {
				j++
			}
			for j > i {
				original[j], original[j-1] = original[j-1], original[j]
				swaps++
				j--
			}
		}
	}
	return swaps
}
```
