package main

// LeetCode #2186: Minimum Number of Steps to Make Two Strings Anagram II
// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram-ii/
// Difficulty: Medium
//
// LOGIC:
// 1. Versi upgrade dari #1347 — beda: |s| tidak harus == |t|, kita boleh delete karakter
//    (satu langkah = satu delete/append karakter di SALAH SATU string)
// 2. Hitung selisih frekuensi: karakter di s → +1, karakter di t → -1
// 3. Nilai positif = karakter di s yang TIDAK ada di t → harus dihapus dari s
// 4. Nilai negatif = karakter di t yang TIDAK ada di s → harus dihapus dari t
// 5. Total langkah = |∑positif| + |∑negatif| = semua karakter yang tidak punya pasangan
// 6. Intuisi: kita bikin dua string menjadi anagram dengan membuang semua karakter yang tidak cocok
//
// Time: O(n + m) | Space: O(1)

import "fmt"

func minSteps(s string, t string) int {
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}
	for _, ch := range t {
		count[ch-'a']--
	}

	steps := 0
	for _, c := range count {
		if c > 0 {
			steps += c
		} else {
			steps -= c
		}
	}
	return steps
}

func main() {
	// Test case 1
	fmt.Println(minSteps("leetcode", "coats"))
	// Expected: 7

	// Test case 2
	fmt.Println(minSteps("night", "thing"))
	// Expected: 0

	// Test case 3
	fmt.Println(minSteps("aba", "bab"))
	// Expected: 2
}
