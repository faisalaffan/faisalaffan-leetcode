# 0862 — Shortest Subarray With Sum At Least K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestSubarray(nums []int, k int) int
```

> **💡 Hint:** Monotonic increasing deque of prefix sums. For each prefix sum, maintain

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Trie, Prefix Sum, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Trie** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #862: Shortest Subarray with Sum at Least K
// https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/
// Difficulty: Hard
// Approach: Monotonic increasing deque of prefix sums. For each prefix sum, maintain
// a deque where sums are increasing. When a new prefix sum makes earlier entries
// irrelevant (they are larger), pop from back. When the difference from the front
// is >= K, update answer and pop from front.

import "fmt"

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
  // Alokasi slice integer
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	ans := n + 1
	// deque stores indices with increasing prefix sums
  // Alokasi slice integer
	deque := make([]int, 0)

	for i := 0; i <= n; i++ {
		// While the front of deque gives a valid subarray, update answer
		for len(deque) > 0 && prefix[i]-prefix[deque[0]] >= k {
			if i-deque[0] < ans {
				ans = i - deque[0]
			}
			deque = deque[1:] // Pop front
		}

		// Maintain monotonic increasing property: pop back while prefix is larger
		for len(deque) > 0 && prefix[i] <= prefix[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1] // Pop back
		}

		deque = append(deque, i)
	}

	if ans == n+1 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(shortestSubarray([]int{1}, 1))    // Expected: 1
	fmt.Println(shortestSubarray([]int{1, 2}, 4)) // Expected: -1
	fmt.Println(shortestSubarray([]int{2, -1, 2}, 3)) // Expected: 3
}
```
