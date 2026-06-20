# 3170 — Lexicographically Minimum String After Removing Stars

## Deskripsi

**Soal:** [3170. Lexicographically Minimum String After Removing Stars](https://leetcode.com/problems/lexicographically-minimum-string-after-removing-stars/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * 26)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Queue (antrian FIFO)

**Fungsi Solusi:** `func clearStars(s string) string`

## Solusi Go

```go
package main

// LeetCode #3170: Lexicographically Minimum String After Removing Stars
// https://leetcode.com/problems/lexicographically-minimum-string-after-removing-stars/
// Difficulty: Medium
// Time: O(n * 26) | Space: O(n)

import "fmt"

func clearStars(s string) string {
	n := len(s)
	bytes := []byte(s)
  // Membuat slice 2D untuk DP/tabel
	queues := make([][]int, 26)
  // Iterasi seluruh elemen
	for i := range queues {
		queues[i] = make([]int, 0)
	}

	for i := 0; i < n; i++ {
		if s[i] == '*' {
			for j := 0; j < 26; j++ {
				if len(queues[j]) > 0 {
					idx := queues[j][len(queues[j])-1]
					queues[j] = queues[j][:len(queues[j])-1]
					bytes[idx] = '*'
					break
				}
			}
			bytes[i] = '*'
		} else {
			queues[s[i]-'a'] = append(queues[s[i]-'a'], i)
		}
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]byte, 0, n)
	for _, ch := range bytes {
		if ch != '*' {
			ans = append(ans, ch)
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(clearStars("aaba*"))       // Expected: "aab"
	fmt.Println(clearStars("abc"))          // Expected: "abc"
	fmt.Println(clearStars("a*b*c*"))       // Expected: ""
}
```
