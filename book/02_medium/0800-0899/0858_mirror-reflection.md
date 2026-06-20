# 0858 — Mirror Reflection

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func MirrorReflection(p int, q int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log min(p,q))  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #858: Mirror Reflection
// https://leetcode.com/problems/mirror-reflection/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MirrorReflection(2, 1))
	fmt.Println(MirrorReflection(3, 1))
	fmt.Println(MirrorReflection(4, 3))
}

// Time: O(log min(p,q)) | Space: O(1)
func MirrorReflection(p int, q int) int {
	g := gcd(p, q)
	p /= g
	q /= g

	if p%2 == 0 {
		return 2
	}
	if q%2 == 0 {
		return 0
	}
	return 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
