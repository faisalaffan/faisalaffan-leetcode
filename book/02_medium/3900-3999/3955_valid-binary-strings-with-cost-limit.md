# 3955 — Valid Binary Strings With Cost Limit

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func ValidBinaryStringsWithCostLimit(n int, k int) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS, Backtracking

**Waktu:** O(2^N)  |  **Ruang:** O(N * 2^N)

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3955: Valid Binary Strings With Cost Limit
// https://leetcode.com/problems/valid-binary-strings-with-cost-limit/
// Difficulty: Medium
// Time: O(2^N) | Space: O(N * 2^N)
// Approach: Backtracking. Generate all strings without consecutive 1s,
// filter by cost (sum of indices where s[i]=='1') <= k.

import "fmt"

func ValidBinaryStringsWithCostLimit(n int, k int) []string {
	var ans []string
	var dfs func(pos int, prev byte, sum int, buf []byte)
	dfs = func(pos int, prev byte, sum int, buf []byte) {
		if sum > k {
			return
		}
		if pos == n {
			ans = append(ans, string(buf))
			return
		}

		// Place '0'
		buf[pos] = '0'
		dfs(pos+1, '0', sum, buf)

		// Place '1' only if no consecutive 1s
		if prev != '1' {
			buf[pos] = '1'
			dfs(pos+1, '1', sum+pos, buf)
		}
	}
	buf := make([]byte, n)
	dfs(0, '0', 0, buf)
	return ans
}

func main() {
	// Example 1
	fmt.Println(ValidBinaryStringsWithCostLimit(3, 1)) // Expected: ["000","010","100"]

	// Example 2
	fmt.Println(ValidBinaryStringsWithCostLimit(1, 0)) // Expected: ["0"]
}
```
