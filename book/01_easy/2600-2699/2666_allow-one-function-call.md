# 2666 — Allow One Function Call

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func AllowOneFunctionCall(fn func(int) int) func(int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2666: Allow One Function Call
// https://leetcode.com/problems/allow-one-function-call/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Ensures fn is called at most once.

import "fmt"

func main() {
	fn := func(x int) int { return x * 2 }
	onceFn := AllowOneFunctionCall(fn)
	fmt.Println(onceFn(5))
	fmt.Println(onceFn(10))
}

func AllowOneFunctionCall(fn func(int) int) func(int) int {
	called := false
	return func(x int) int {
		if called {
			return 0
		}
		called = true
		return fn(x)
	}
}
```
