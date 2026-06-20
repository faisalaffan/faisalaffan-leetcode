# 2950 — Number Of Divisible Substrings

## Deskripsi

**Soal:** [2950. Number Of Divisible Substrings](https://leetcode.com/problems/number-of-divisible-substrings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func numberOfDivisibleSubstrings(s string) int`

## Solusi Go

```go
package main

// LeetCode #2950: Number of Divisible Substrings
// https://leetcode.com/problems/number-of-divisible-substrings/
// Difficulty: Medium

import "fmt"

func numberOfDivisibleSubstrings(s string) int {
	n := len(s)

	// Build mapping: old phone keypad, each char maps to a digit
	// a,b,c->1; d,e,f->2; g,h,i->3; j,k,l->4; m,n,o->5;
	// p,q,r->6; s,t,u->7; v,w,x->8; y,z->9
  // Membuat slice untuk menyimpan hasil
	mapping := make([]int, 26)
	for i := 0; i < 26; i++ {
		if i <= 2 { // a,b,c -> 1
			mapping[i] = 1
		} else if i <= 5 { // d,e,f -> 2
			mapping[i] = 2
		} else if i <= 8 { // g,h,i -> 3
			mapping[i] = 3
		} else if i <= 11 { // j,k,l -> 4
			mapping[i] = 4
		} else if i <= 14 { // m,n,o -> 5
			mapping[i] = 5
		} else if i <= 17 { // p,q,r -> 6
			mapping[i] = 6
		} else if i <= 20 { // s,t,u -> 7
			mapping[i] = 7
		} else if i <= 23 { // v,w,x -> 8
			mapping[i] = 8
		} else { // y,z -> 9
			mapping[i] = 9
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += mapping[s[j]-'a']
			length := j - i + 1
			if sum%length == 0 {
				count++
			}
		}
	}

	return count
}

func main() {
	// Test case 1: "abc"
	// a=1, b=1, c=2
	// Substrings: "a"(1%1=0), "b"(1%1=0), "c"(2%1=0),
	// "ab"(2%2=0), "bc"(3%2!=0), "abc"(4%3!=0) -> 4
	fmt.Println(numberOfDivisibleSubstrings("abc")) // 4

	// Test case 2: single character
	fmt.Println(numberOfDivisibleSubstrings("a")) // 1

	// Test case 3: "abcd"
	// a=1,b=1,c=2,d=2
	// "a"(1), "b"(1), "c"(2), "d"(2),
	// "ab"(2%2=0), "bc"(3%2!=0), "cd"(4%2=0),
	// "abc"(4%3!=0), "bcd"(5%3!=0),
	// "abcd"(6%4!=0)
	fmt.Println(numberOfDivisibleSubstrings("abcd")) // 6

	// Test case 4: empty-ish (single char repeated)
	fmt.Println(numberOfDivisibleSubstrings("aaa")) // 6
	// "a"(1), "a"(1), "a"(1), "aa"(2%2=0), "aa"(2%2=0), "aaa"(3%3=0) -> 6
}

// Time: O(n^2) | Space: O(1)
```
