package main

// LeetCode #2514: Count Anagrams
// https://leetcode.com/problems/count-anagrams/
// Difficulty: Hard
//
// LOGIC:
// 1. Berapa banyak anagram DISTINCT yang bisa dibentuk dari string s?
// 2. Untuk setiap kata terpisah (dipisah spasi):
//    - Jumlah permutasi jika semua huruf unik = n!
//    - Tapi huruf yang berulang menyebabkan duplikasi → bagi dengan cnt[c]!
//    - Rumus: distinct anagrams = n! / (cnt['a']! * cnt['b']! * ... * cnt['z']!)
// 3. Kalikan hasil semua kata → total anagram distinct dari string s
// 4. Karena angka bisa sangat besar, semua operasi MOD 1e9+7
// 5. Pembagian di modulo pakai Fermat's Little Theorem:
//    a / b (mod M) = a * b^(M-2) (mod M), karena M prima
// 6. Precompute factorial n! untuk setiap n yang dibutuhkan
// 7. Gunakan modular inverse (modPow) untuk menghitung 1 / cnt[c]!
//
// Contoh: "too" → 3! / 2! = 6/2 = 3. Tapi "too" hanya punya 3 anagram distinct:
// too, oto, oot (benar: 3! / 2! = 6/2 = 3) — wait, seharusnya bukan 3...
// Rumus tepat: 3! / (cnt[t]! * cnt[o]!) = 6 / (1! * 2!) = 6/2 = 3 ✓

import (
	"fmt"
	"strings"
)

const MOD = 1000000007

func main() {
	// Example 1: "too" => 2
	fmt.Println(countAnagrams("too"))
	// Example 2: "aa aa" => 1
	fmt.Println(countAnagrams("aa aa"))
	// Edge: single char
	fmt.Println(countAnagrams("a"))
	// Edge: all same letters
	fmt.Println(countAnagrams("aaa"))
	// Edge: multiple words
	fmt.Println(countAnagrams("abc def ghi"))
}

func countAnagrams(s string) int {
	words := strings.Fields(s)
	result := int64(1)

	for _, word := range words {
		n := len(word)
		// Count character frequencies
		cnt := make(map[rune]int)
		for _, ch := range word {
			cnt[ch]++
		}

		// result *= n! / product(cnt[c]!)
		// Compute n! * inverse(product(cnt[c]!))
		res := factorial(n)
		for _, c := range cnt {
			inv := modInv(factorial(c))
			res = (res * inv) % MOD
		}
		result = (result * res) % MOD
	}

	return int(result)
}

func factorial(n int) int64 {
	res := int64(1)
	for i := 2; i <= n; i++ {
		res = (res * int64(i)) % MOD
	}
	return res
}

func modInv(a int64) int64 {
	return modPow(a, MOD-2)
}

func modPow(a int64, b int) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return res
}
