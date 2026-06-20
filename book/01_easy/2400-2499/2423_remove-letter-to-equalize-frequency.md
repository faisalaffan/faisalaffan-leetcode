# 2423 — Remove Letter To Equalize Frequency

## Deskripsi

**Soal:** [2423. Remove Letter To Equalize Frequency](https://leetcode.com/problems/remove-letter-to-equalize-frequency/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2423: Remove Letter To Equalize Frequency
// https://leetcode.com/problems/remove-letter-to-equalize-frequency/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(RemoveLetterToEqualizeFrequency("abcc")) // true
	fmt.Println(RemoveLetterToEqualizeFrequency("aazz")) // false
	fmt.Println(RemoveLetterToEqualizeFrequency("bac"))  // true
}

func RemoveLetterToEqualizeFrequency(word string) bool {
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(word); i++ {
		freq[word[i]-'a']++
	}

	// Try removing one occurrence of each letter
	for i := 0; i < 26; i++ {
		if freq[i] == 0 {
			continue
		}
		freq[i]--
		if allSameFreq(freq) {
			return true
		}
		freq[i]++
	}
	return false
}

func allSameFreq(freq []int) bool {
	target := 0
	for _, f := range freq {
		if f == 0 {
			continue
		}
		if target == 0 {
			target = f
		} else if f != target {
			return false
		}
	}
	return true
}
```
