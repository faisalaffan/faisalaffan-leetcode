# 3265 — Count Almost Equal Pairs I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countPairs(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2 * d) Space: O(d) where d = number of digits  
**Kompleksitas Ruang:** O(d) where d = number of digits

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3265: Count Almost Equal Pairs I
// https://leetcode.com/problems/count-almost-equal-pairs-i/
// Difficulty: Medium
// Time: O(n^2 * d) Space: O(d) where d = number of digits

import "fmt"

func main() {
	fmt.Println(countPairs([]int{1, 1, 1}))                                           // 3
	fmt.Println(countPairs([]int{3, 12, 30, 17, 21}))                                 // 2
	fmt.Println(countPairs([]int{1023, 3012, 1230, 2301, 0, 0}))                     // 4
}

func countPairs(nums []int) int {
	ans := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if isAlmostEqual(nums[i], nums[j]) {
				ans++
			}
		}
	}
	return ans
}

func isAlmostEqual(a, b int) bool {
	sa, sb := fmt.Sprintf("%07d", a), fmt.Sprintf("%07d", b)
	diff := 0
  // Alokasi slice integer
	ca, cb := make([]int, 10), make([]int, 10)
	for k := 0; k < 7; k++ {
		if sa[k] != sb[k] {
			diff++
			if diff > 2 {
				return false
			}
		}
		ca[sa[k]-'0']++
		cb[sb[k]-'0']++
	}
	for k := 0; k < 10; k++ {
		if ca[k] != cb[k] {
			return false
		}
	}
	return diff <= 2
}
```
