# 1348 — Tweet Counts Per Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() TweetCounts
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n + q) where n = tweets count, q = query interval count  
**Kompleksitas Ruang:** O(n) for storing tweets

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1348: Tweet Counts Per Frequency
// https://leetcode.com/problems/tweet-counts-per-frequency/
// Difficulty: Medium

import "fmt"
import "sort"

type TweetCounts struct {
	tweets map[string][]int
}

func main() {
	tc := Constructor()
	tc.RecordTweet("tweet3", 0)
	tc.RecordTweet("tweet3", 60)
	tc.RecordTweet("tweet3", 10)
	fmt.Println(tc.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59)) // [2]
	fmt.Println(tc.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60)) // [2,1]

	tc2 := Constructor()
	tc2.RecordTweet("tweet1", 0)
	tc2.RecordTweet("tweet1", 10)
	tc2.RecordTweet("tweet2", 5)
	fmt.Println(tc2.GetTweetCountsPerFrequency("minute", "tweet2", 0, 100)) // [1]

	// Test case 3 - hour frequency
	tc3 := Constructor()
	tc3.RecordTweet("tweet", 0)
	tc3.RecordTweet("tweet", 3599)
	tc3.RecordTweet("tweet", 3600)
	fmt.Println(tc3.GetTweetCountsPerFrequency("hour", "tweet", 0, 7200)) // [2,1]
}

func Constructor() TweetCounts {
	return TweetCounts{tweets: make(map[string][]int)}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	this.tweets[tweetName] = append(this.tweets[tweetName], time)
}

// Time: O(n log n + q) where n = tweets count, q = query interval count
// Space: O(n) for storing tweets
func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	f := 60 // default minute
	switch freq {
	case "hour":
		f = 3600
	case "day":
		f = 86400
	}

	times := this.tweets[tweetName]
  // Urutkan secara ascending — O(n log n)
	sort.Ints(times)

	size := (endTime-startTime)/f + 1
  // Alokasi slice integer
	result := make([]int, size)

	for _, t := range times {
		if t < startTime || t > endTime {
			continue
		}
		idx := (t - startTime) / f
		result[idx]++
	}

	return result
}
```
