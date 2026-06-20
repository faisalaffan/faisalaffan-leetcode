# 1242 — Web Crawler Multithreaded

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func crawlParallel(startUrl string, parser HtmlParser) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(V + E)  
**Kompleksitas Ruang:** O(V)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Loop linear O(n): iterasi setiap elemen
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
  // Membuat map (HashMap) — pencarian O(1)
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
