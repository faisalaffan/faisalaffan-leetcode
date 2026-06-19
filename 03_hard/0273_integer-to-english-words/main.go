package main

// LeetCode #273: Integer to English Words
// https://leetcode.com/problems/integer-to-english-words/
// Difficulty: Hard
//
// Approach: Groups of Three (Chunking).
//  1. Handle zero separately.
//  2. Process number in groups of 3 digits: billions, millions, thousands, ones.
//  3. For each group, convert the 3-digit number to words using helper functions.
//  4. Append the scale word (Billion, Million, Thousand) for non-zero groups.

import "fmt"

func main() {
	// Example 1: 123 -> "One Hundred Twenty Three"
	fmt.Println("123 ->", numberToWords(123))

	// Example 2: 12345 -> "Twelve Thousand Three Hundred Forty Five"
	fmt.Println("12345 ->", numberToWords(12345))

	// Example 3: 1234567 -> "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
	fmt.Println("1234567 ->", numberToWords(1234567))

	// Edge cases
	fmt.Println("0 ->", numberToWords(0))
	fmt.Println("1000 ->", numberToWords(1000))
	fmt.Println("1000000 ->", numberToWords(1000000))
	fmt.Println("1000000000 ->", numberToWords(1000000000))
	fmt.Println("2147483647 ->", numberToWords(2147483647))
}

var (
	belowTwenty = []string{
		"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
		"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen",
		"Seventeen", "Eighteen", "Nineteen",
	}
	tens = []string{
		"", "", "Twenty", "Thirty", "Forty", "Fifty",
		"Sixty", "Seventy", "Eighty", "Ninety",
	}
	thousands = []string{"", "Thousand", "Million", "Billion"}
)

// numberToWords converts a non-negative integer to English words.
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	result := ""
	unitIndex := 0

	for num > 0 {
		chunk := num % 1000
		if chunk > 0 {
			chunkWords := convertChunk(chunk)
			if result == "" {
				result = chunkWords + " " + thousands[unitIndex]
			} else {
				result = chunkWords + " " + thousands[unitIndex] + " " + result
			}
		}
		num /= 1000
		unitIndex++
	}

	// Clean up extra spaces.
	// We trim spaces at the end instead of complex logic.
	return trimSpace(result)
}

// convertChunk converts a 3-digit number (0-999) to English words.
// Note: chunk is guaranteed > 0 when called.
func convertChunk(num int) string {
	var result string

	hundreds := num / 100
	remainder := num % 100

	if hundreds > 0 {
		result = belowTwenty[hundreds] + " Hundred"
	}

	if remainder > 0 {
		if result != "" {
			result += " "
		}
		if remainder < 20 {
			result += belowTwenty[remainder]
		} else {
			ten := remainder / 10
			one := remainder % 10
			result += tens[ten]
			if one > 0 {
				result += " " + belowTwenty[one]
			}
		}
	}

	return result
}

// trimSpace removes trailing spaces from a string.
func trimSpace(s string) string {
	if len(s) == 0 {
		return s
	}
	// Find the last non-space character.
	end := len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	return s[:end+1]
}

// Stub compatibility.
func IntegerToEnglishWords() any {
	return numberToWords(123)
}
