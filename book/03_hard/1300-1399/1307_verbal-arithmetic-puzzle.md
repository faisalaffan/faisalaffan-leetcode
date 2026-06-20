# 1307 — Verbal Arithmetic Puzzle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isSolvable(words []string, result string) bool
```

> **💡 Hint:** Backtracking with digit assignment.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Backtracking

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1307: Verbal Arithmetic Puzzle
// https://leetcode.com/problems/verbal-arithmetic-puzzle/
// Difficulty: Hard
//
// Approach: Backtracking with digit assignment.
// Collect all unique letters (max 10). Try each digit 0-9 for each letter,
// respecting the constraint that leading letters cannot be 0.
// When all letters are assigned, verify the equation: word[0] + ... + word[k] == result.
// Prune: leading-letter-zero check, no digit reuse.

import "fmt"

func isSolvable(words []string, result string) bool {
  // Membuat map (HashMap) — pencarian O(1)
	letterSet := make(map[byte]bool)
	addLetters := func(s string) {
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(s); i++ {
			letterSet[s[i]] = true
		}
	}
	for _, w := range words {
		addLetters(w)
	}
	addLetters(result)

	letters := make([]byte, 0, len(letterSet))
	for c := range letterSet {
		letters = append(letters, c)
	}
	if len(letters) > 10 {
		return false
	}

  // Membuat map (HashMap) — pencarian O(1)
	nonZero := make(map[byte]bool)
	for _, w := range words {
		if len(w) > 1 {
			nonZero[w[0]] = true
		}
	}
	if len(result) > 1 {
		nonZero[result[0]] = true
	}

  // Membuat map (HashMap) — pencarian O(1)
	mapping := make(map[byte]int)
	used := make([]bool, 10)

	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == len(letters) {
			sum := 0
			for _, w := range words {
				val := 0
  // Loop linear O(n): iterasi setiap elemen
				for i := 0; i < len(w); i++ {
					val = val*10 + mapping[w[i]]
				}
				sum += val
			}
			res := 0
  // Loop linear O(n): iterasi setiap elemen
			for i := 0; i < len(result); i++ {
				res = res*10 + mapping[result[i]]
			}
			return sum == res
		}

		c := letters[idx]
		for d := 0; d <= 9; d++ {
			if used[d] {
				continue
			}
			if d == 0 && nonZero[c] {
				continue
			}
			used[d] = true
			mapping[c] = d
			if dfs(idx + 1) {
				return true
			}
			used[d] = false
		}
		return false
	}

	return dfs(0)
}

func main() {
	fmt.Println(isSolvable([]string{"SEND", "MORE"}, "MONEY"))                  // true
	fmt.Println(isSolvable([]string{"SIX", "SEVEN", "SEVEN"}, "TWENTY"))        // true
	fmt.Println(isSolvable([]string{"LEET", "CODE"}, "POINT"))                  // false
	fmt.Println(isSolvable([]string{"A", "B"}, "A"))                            // true (B=0)
}
```
