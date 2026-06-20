# 2023 — Number Of Pairs Of Strings With Concatenation Equal To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func numOfPairs(nums []string, target string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n * L)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2023: Number of Pairs of Strings With Concatenation Equal to Target
// https://leetcode.com/problems/number-of-pairs-of-strings-with-concatenation-equal-to-target/
// Difficulty: Medium
// Time: O(n * L) | Space: O(n)

import "fmt"

func numOfPairs(nums []string, target string) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[string]int)
	count := 0

	for _, num := range nums {
		// Check if any prefix/suffix of target matches
		for i := 1; i < len(target); i++ {
			prefix := target[:i]
			suffix := target[i:]
			if num == prefix {
				count += freq[suffix]
			}
			if num == suffix {
				count += freq[prefix]
			}
		}
		freq[num]++
	}

	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numOfPairs([]string{"777", "7", "77", "77"}, "777"))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", numOfPairs([]string{"123", "4", "12", "34"}, "1234"))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", numOfPairs([]string{"1", "1", "1"}, "11"))
	// Expected: 6
}
```
