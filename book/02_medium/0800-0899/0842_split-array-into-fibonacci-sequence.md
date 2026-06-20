# 0842 — Split Array Into Fibonacci Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SplitArrayIntoFibonacciSequence(num string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #842: Split Array into Fibonacci Sequence
// https://leetcode.com/problems/split-array-into-fibonacci-sequence/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SplitArrayIntoFibonacciSequence("123456579"))
	fmt.Println(SplitArrayIntoFibonacciSequence("11235813"))
	fmt.Println(SplitArrayIntoFibonacciSequence("112358130"))
}

// Time: O(n^2) | Space: O(n)
func SplitArrayIntoFibonacciSequence(num string) []int {
	n := len(num)
	var ans []int

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == n {
			return len(ans) > 2
		}

		var x int
		for i := pos; i < n; i++ {
			if i > pos && num[pos] == '0' {
				break
			}
			x = x*10 + int(num[i]-'0')
			if x > math.MaxInt32 {
				break
			}
			if len(ans) > 1 && x > ans[len(ans)-1]+ans[len(ans)-2] {
				break
			}
			if len(ans) < 2 || x == ans[len(ans)-1]+ans[len(ans)-2] {
				ans = append(ans, x)
				if dfs(i + 1) {
					return true
				}
				ans = ans[:len(ans)-1]
			}
		}
		return false
	}

	dfs(0)
	return ans
}
```
