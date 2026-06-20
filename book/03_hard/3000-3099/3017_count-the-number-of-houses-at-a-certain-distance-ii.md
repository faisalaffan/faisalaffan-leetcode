# 3017 — Count The Number Of Houses At A Certain Distance Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countOfPairs(n int, x int, y int) []int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3017: Count the Number of Houses at a Certain Distance II
// https://leetcode.com/problems/count-the-number-of-houses-at-a-certain-distance-ii/
// Difficulty: Hard
//
// There are n houses on a line numbered 1 to n. There is an extra edge
// between house x and house y. For each distance d (1..n-1), count how
// many pairs of houses (i, j) with i < j have shortest path distance d.
//
// Approach: Sweep using difference array
//   Model the distances from each house to all others. Use the fact that
//   the extra edge creates a "shortcut" that splits the line into two
//   segments. Use difference arrays to add contributions efficiently.

import "fmt"

func countOfPairs(n int, x int, y int) []int64 {
	if x > y {
		x, y = y, x
	}
  // Alokasi slice
	diff := make([]int64, n+1)

	add := func(l, r int, val int64) {
		if l > r {
			return
		}
		if r > n {
			r = n
		}
		diff[l] += val
		if r+1 <= n {
			diff[r+1] -= val
		}
	}

	for i := 1; i <= n; i++ {
		if x+1 >= y {
			// No shortcut effect (x and y are adjacent or same)
			add(1, n-i, 2)
			continue
		}

		// The line is split by the edge (x, y) into two segments:
		//   left segment: [1, x]
		//   middle segment: [x+1, y-1]
		//   right segment: [y, n]
		if i <= x {
			// i is in the left segment
			k := (x + y + 1) / 2
			// Direct distances within left portion
			add(1, k-i, 2)
			// Distances using the shortcut, going to right side
			add(x-i+2, x-i+y-k, 2)
			// Distances going to the far right (beyond y)
			add(x-i+1, x-i+1+n-y, 2)
		} else if i < (x+y)/2 {
			// i is in the middle or right segment, but closer to left side
			k := i + (y-x+1)/2
			add(1, k-i, 2)
			add(i-x+2, i-x+y-k, 2)
			add(i-x+1, i-x+1+n-y, 2)
		} else {
			// i is on the right side; contributions are symmetric to left side
			add(1, n-i, 2)
		}
	}

  // Alokasi slice
	ans := make([]int64, n)
	cur := int64(0)
	for i := 1; i <= n; i++ {
		cur += diff[i]
		ans[i-1] = cur
	}
	return ans
}

func main() {
	fmt.Println("Test 1: n=3, x=1, y=3")
	res := countOfPairs(3, 1, 3)
	for d, cnt := range res {
		fmt.Printf("  distance=%d: %d\n", d+1, cnt)
	}

	fmt.Println("Test 2: n=5, x=2, y=4")
	res = countOfPairs(5, 2, 4)
	for d, cnt := range res {
		fmt.Printf("  distance=%d: %d\n", d+1, cnt)
	}

	fmt.Println("Test 3: n=4, x=1, y=1")
	res = countOfPairs(4, 1, 1)
	for d, cnt := range res {
		fmt.Printf("  distance=%d: %d\n", d+1, cnt)
	}

	fmt.Println("Test 4: n=6, x=3, y=5")
	res = countOfPairs(6, 3, 5)
	for d, cnt := range res {
		fmt.Printf("  distance=%d: %d\n", d+1, cnt)
	}
}
```
