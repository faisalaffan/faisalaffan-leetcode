# 3874 — Valid Subarrays With Exactly One Peak

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func ValidSubarraysWithExactlyOnePeak(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(N)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3874: Valid Subarrays With Exactly One Peak
// https://leetcode.com/problems/valid-subarrays-with-exactly-one-peak/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(N)
// Approach: Find all peaks. For each peak, count valid subarrays that contain
// exactly this peak, bounded by adjacent peaks and distance k from the peak.

import "fmt"

func ValidSubarraysWithExactlyOnePeak(nums []int, k int) int {
	n := len(nums)

	// Find peak indices
	peaks := []int{}
	for i := 1; i < n-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			peaks = append(peaks, i)
		}
	}

	if len(peaks) == 0 {
		return 0
	}

	ans := 0
	for idx, p := range peaks {
		// Left bound: can't go beyond k steps from peak, and can't include previous peak
		leftBound := p - k
		if idx > 0 && peaks[idx-1] >= leftBound {
			leftBound = peaks[idx-1] + 1
		}
		if leftBound < 0 {
			leftBound = 0
		}

		// Right bound: can't go beyond k steps from peak, and can't include next peak
		rightBound := p + k
		if idx < len(peaks)-1 && peaks[idx+1] <= rightBound {
			rightBound = peaks[idx+1] - 1
		}
		if rightBound >= n {
			rightBound = n - 1
		}

		leftOptions := p - leftBound + 1
		rightOptions := rightBound - p + 1
		ans += leftOptions * rightOptions
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(ValidSubarraysWithExactlyOnePeak([]int{1, 3, 2}, 1)) // Expected: 4

	// Example 2
	fmt.Println(ValidSubarraysWithExactlyOnePeak([]int{7, 8, 9}, 2)) // Expected: 0

	// Example 3
	fmt.Println(ValidSubarraysWithExactlyOnePeak([]int{4, 3, 5, 1}, 2)) // Expected: 6
}
```
