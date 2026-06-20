# 1242 — Web Crawler Multithreaded

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func crawlParallel(startUrl string, parser HtmlParser) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(V + E)  |  **Ruang:** O(V)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sync"
)

// LeetCode #1242: Web Crawler Multithreaded
// https://leetcode.com/problems/web-crawler-multithreaded/
// Difficulty: Medium [Paid]

// Multithreaded web crawler using goroutines and channels.
// Same hostname constraint.

// Time: O(V + E)
// Space: O(V)

type HtmlParser interface {
	GetUrls(url string) []string
}

type mockParser struct {
	urls map[string][]string
}

func (m *mockParser) GetUrls(url string) []string {
	return m.urls[url]
}

func crawlParallel(startUrl string, parser HtmlParser) []string {
	getHost := func(url string) string {
		host := ""
  // Linear scan O(n)
		for i := 0; i < len(url)-7; i++ {
			if url[i:i+7] == "http://" {
				url = url[7:]
				break
			}
		}
		for _, c := range url {
			if c == '/' || c == ':' {
				break
			}
			host += string(c)
		}
		return host
	}

	hostname := getHost(startUrl)
  // HashMap: O(1) lookup
	visited := make(map[string]bool)
	var mu sync.Mutex
	var wg sync.WaitGroup

	var crawl func(url string)
	crawl = func(url string) {
		defer wg.Done()
		urls := parser.GetUrls(url)
		for _, next := range urls {
			mu.Lock()
			if visited[next] || getHost(next) != hostname {
				mu.Unlock()
				continue
			}
			visited[next] = true
			mu.Unlock()

			wg.Add(1)
			go crawl(next)
		}
	}

	visited[startUrl] = true
	wg.Add(1)
	go crawl(startUrl)
	wg.Wait()

	result := make([]string, 0, len(visited))
	for url := range visited {
		result = append(result, url)
	}
	return result
}

func main() {
	parser := &mockParser{
		urls: map[string][]string{
			"http://example.com":       {"http://example.com/about", "http://other.com"},
			"http://example.com/about": {"http://example.com/contact"},
			"http://example.com/contact": {},
		},
	}
	result := crawlParallel("http://example.com", parser)
	fmt.Printf("%v\n", result)
}
```
