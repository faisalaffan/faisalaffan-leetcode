package main

// LeetCode #1347: Minimum Number of Steps to Make Two Strings Anagram
// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/
// Difficulty: Medium
//
// LOGIC:
// 1. Soal: berapa langkah minimum mengubah karakter di t agar s dan t jadi anagram?
//    (satu langkah = satu perubahan karakter di t)
// 2. Hitung selisih frekuensi: untuk setiap karakter, s[i] → +1, t[i] → -1
// 3. Nilai positif di freq[c] = kelebihan huruf c di s yang tidak punya pasangan di t
// 4. Jumlahkan semua nilai positif = total langkah yang dibutuhkan
// 5. Intuisi: kita hanya perlu mengganti karakter di t yang "kurang" match dengan s
//    Karena |s| = |t|, jumlah kelebihan = jumlah kekurangan

import "fmt"

func main() {
	// Test case 1
	fmt.Println(minSteps("bab", "aba")) // 1

	// Test case 2
	fmt.Println(minSteps("leetcode", "practice")) // 5

	// Test case 3
	fmt.Println(minSteps("anagram", "mangaar")) // 0
}

// Time: O(n) where n = length of strings
// Space: O(1) - fixed size array of 26
func minSteps(s string, t string) int {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	steps := 0
	for _, f := range freq {
		if f > 0 {
			steps += f
		}
	}
	return steps
}
