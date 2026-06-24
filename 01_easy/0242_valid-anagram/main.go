package main

// LeetCode #242: Valid Anagram
// https://leetcode.com/problems/valid-anagram/
// Difficulty: Easy
//
// LOGIC:
// 1. Jika panjang s != t, langsung return false
// 2. Buat dua array [26]int: hitung frekuensi karakter s dan t
// 3. Iterasi s dan t: countS[s[i]-'a']++, countT[t[i]-'a']++
// 4. Bandingkan kedua array: jika identik, s dan t adalah anagram

import "fmt"

// Time: O(n) | Space: O(1) (fixed 26+26 chars)
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var countS, countT [26]int
	for i := 0; i < len(s); i++ {
		countS[s[i]-'a']++
		countT[t[i]-'a']++
	}
	fmt.Println(countS)
	fmt.Println(countT)
	return countS == countT
}

func main() {
	fmt.Println(IsAnagram("anagram", "nagaram"))
	// fmt.Println(IsAnagram("anagram", "manager"))
	// fmt.Println(IsAnagram("rat", "car"))
}
