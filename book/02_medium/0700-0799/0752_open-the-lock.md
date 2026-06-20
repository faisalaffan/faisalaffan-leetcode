# 0752 — Open The Lock

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func openLock(deadends []string, target string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, BFS

**Waktu:** O(10^4 * 8) ~ O(1)  |  **Ruang:** O(10^4)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #752: Open the Lock
// https://leetcode.com/problems/open-the-lock/
// Difficulty: Medium
// Time: O(10^4 * 8) ~ O(1)
// Space: O(10^4)

import "fmt"

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"8888"}, "0009"))
}

func openLock(deadends []string, target string) int {
  // HashMap: O(1) lookup
	dead := make(map[string]bool)
	for _, d := range deadends {
		dead[d] = true
	}

	if dead["0000"] {
		return -1
	}

  // HashMap: O(1) lookup
	visited := make(map[string]bool)
	queue := []string{"0000"}
	visited["0000"] = true
	steps := 0

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			curr := queue[i]
			if curr == target {
				return steps
			}

			for j := 0; j < 4; j++ {
				for _, d := range []int{-1, 1} {
					next := []byte(curr)
					next[j] = byte('0' + (int(next[j]-'0')+d+10)%10)
					s := string(next)
					if !visited[s] && !dead[s] {
						visited[s] = true
						queue = append(queue, s)
					}
				}
			}
		}
		queue = queue[n:]
		steps++
	}

	return -1
}
```
