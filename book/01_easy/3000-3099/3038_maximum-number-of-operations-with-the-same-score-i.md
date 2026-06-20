# 3038 — Maximum Number Of Operations With The Same Score I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumNumberOfOperationsWithTheSameScoreI(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3038: Maximum Number of Operations With the Same Score I
// https://leetcode.com/problems/maximum-number-of-operations-with-the-same-score-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maxOperations
	fmt.Println(MaximumNumberOfOperationsWithTheSameScoreI([]int{3, 2, 1, 4, 5})) // 2
	fmt.Println(MaximumNumberOfOperationsWithTheSameScoreI([]int{3, 2, 6, 1, 4})) // 1
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: maxOperations
func MaximumNumberOfOperationsWithTheSameScoreI(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	score := nums[0] + nums[1]
	count := 1
	for i := 2; i+1 < len(nums); i += 2 {
		if nums[i]+nums[i+1] == score {
			count++
		} else {
			break
		}
	}
	return count
}
```
