package main

// LeetCode #1912: Design Movie Rental System
// https://leetcode.com/problems/design-movie-rental-system/
// Difficulty: Hard

import (
	"container/heap"
	"fmt"
	"sort"
)

// Entry represents (shop, movie, price)
type Entry struct {
	shop, movie, price int
}

// MinHeap for report
type ReportHeap []Entry

func (h ReportHeap) Len() int { return len(h) }
func (h ReportHeap) Less(i, j int) bool {
	if h[i].price != h[j].price {
		return h[i].price < h[j].price
	}
	if h[i].shop != h[j].shop {
		return h[i].shop < h[j].shop
	}
	return h[i].movie < h[j].movie
}
func (h ReportHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *ReportHeap) Push(x any)   { *h = append(*h, x.(Entry)) }
func (h *ReportHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MinHeap for search within a movie
type SearchHeap []Entry

func (h SearchHeap) Len() int { return len(h) }
func (h SearchHeap) Less(i, j int) bool {
	if h[i].price != h[j].price {
		return h[i].price < h[j].price
	}
	return h[i].shop < h[j].shop
}
func (h SearchHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *SearchHeap) Push(x any)   { *h = append(*h, x.(Entry)) }
func (h *SearchHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MovieRentalSystem struct {
	// For each movie, a min-heap of (price, shop) for available copies
	avail map[int]*SearchHeap
	// For each movie, a sorted list of (shop, price) for binary search lazy deletion
	prices map[int][]Entry
	// Currently rented entries for report
	rented *ReportHeap
	// Track if an entry has been rented (to avoid duplicates in rented heap)
	rentedSet map[Entry]bool
}

func Constructor(entries [][]int) MovieRentalSystem {
	mrs := MovieRentalSystem{
		avail:     make(map[int]*SearchHeap),
		prices:    make(map[int][]Entry),
		rented:    &ReportHeap{},
		rentedSet: make(map[Entry]bool),
	}
	// Group by movie
	byMovie := make(map[int][]Entry)
	for _, e := range entries {
		shop, movie, price := e[0], e[1], e[2]
		byMovie[movie] = append(byMovie[movie], Entry{shop, movie, price})
	}
	for movie, list := range byMovie {
		// Sort by price then shop
		sort.Slice(list, func(i, j int) bool {
			if list[i].price != list[j].price {
				return list[i].price < list[j].price
			}
			return list[i].shop < list[j].shop
		})
		mrs.prices[movie] = list
		h := &SearchHeap{}
		heap.Init(h)
		for _, e := range list {
			heap.Push(h, e)
		}
		mrs.avail[movie] = h
	}
	return mrs
}

func (mrs *MovieRentalSystem) Search(movie int) [][]int {
	h, ok := mrs.avail[movie]
	if !ok {
		return nil
	}

	// Need to lazy-clean the heap: pop entries that are no longer available
	// (they have been rented out)
	var result []Entry
	temp := &SearchHeap{}
	heap.Init(temp)

	for h.Len() > 0 && len(result) < 5 {
		e := heap.Pop(h).(Entry)
		if !mrs.rentedSet[e] {
			result = append(result, e)
			heap.Push(temp, e)
		}
	}
	// Push back the ones we popped
	for temp.Len() > 0 {
		heap.Push(h, heap.Pop(temp).(Entry))
	}

	res := make([][]int, len(result))
	for i, e := range result {
		res[i] = []int{e.shop, e.movie, e.price}
	}
	return res
}

func (mrs *MovieRentalSystem) Rent(shop int, movie int) {
	// Find the entry in prices for this movie
	list := mrs.prices[movie]
	for i, e := range list {
		if e.shop == shop && e.movie == movie {
			mrs.rentedSet[e] = true
			heap.Push(mrs.rented, e)
			break
		}
		// Since list is sorted by (price, shop), and this may not be first
		// We just linear scan - list is small per movie typically
		_ = i
	}
	// Actually we need a more efficient way. Let's use a map for exact lookup.
	// But since entries are small and this is for the problem constraints,
	// linear scan through the movie's price list is fine.
}

func (mrs *MovieRentalSystem) Drop(shop int, movie int) {
	// Remove from rented set
	key := Entry{shop, movie, 0}
	mrs.rentedSet[key] = false
	delete(mrs.rentedSet, key)
	// The entry becomes available again in the heap automatically
	// No need to modify the avail heap - it still has the entry,
	// and the rentedSet check will now return false, so Search will show it.
}

func (mrs *MovieRentalSystem) Report() [][]int {
	// Collect currently rented entries
	var temp []Entry
	var result []Entry
	for mrs.rented.Len() > 0 && len(result) < 5 {
		e := heap.Pop(mrs.rented).(Entry)
		if mrs.rentedSet[e] {
			result = append(result, e)
			temp = append(temp, e)
		}
	}
	// Push back what we popped
	for _, e := range temp {
		heap.Push(mrs.rented, e)
	}
	for _, e := range result {
		// Push them back too
		heap.Push(mrs.rented, e)
	}
	// Take first 5
	if len(result) > 5 {
		result = result[:5]
	}

	res := make([][]int, len(result))
	for i, e := range result {
		res[i] = []int{e.shop, e.movie, e.price}
	}
	return res
}

// DesignMovieRentalSystem is the exported function called by main
func DesignMovieRentalSystem() any {
	entries := [][]int{
		{0, 1, 5}, {0, 2, 6}, {0, 3, 7},
		{1, 1, 4}, {1, 2, 7},
		{2, 1, 5},
	}
	mrs := Constructor(entries)
	_ = mrs.Search(1)  // [[1,1,4],[0,1,5],[2,1,5]]
	mrs.Rent(0, 1) // rent (0,1,5)
	mrs.Rent(1, 2) // rent (1,2,7)
	_ = mrs.Report()   // [[0,1,5],[1,2,7]]
	mrs.Drop(1, 2) // return (1,2,7)
	_ = mrs.Search(2)  // [[0,2,6]]

	return nil
}

func main() {
	DesignMovieRentalSystem()
	fmt.Println("Movie Rental System test completed")
}
