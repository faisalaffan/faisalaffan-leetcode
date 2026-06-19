package main

// LeetCode #1500: Design a File Sharing System
// https://leetcode.com/problems/design-a-file-sharing-system/
// Difficulty: Medium [Paid]

import "fmt"

// FileSharing simulates a file sharing system.
type FileSharing struct {
	chunks   map[int][]int   // user -> list of chunk IDs
	nextID   int             // next user ID
	released []int           // reusable user IDs
}

func Constructor(m int) FileSharing {
	return FileSharing{
		chunks:   make(map[int][]int),
		nextID:   1,
		released: make([]int, 0),
	}
}

func (fs *FileSharing) Join(ownedChunks []int) int {
	var userID int
	if len(fs.released) > 0 {
		// Reuse the smallest released ID
		userID = fs.released[0]
		fs.released = fs.released[1:]
	} else {
		userID = fs.nextID
		fs.nextID++
	}

	// Store a copy of chunks
	chunks := make([]int, len(ownedChunks))
	copy(chunks, ownedChunks)
	fs.chunks[userID] = chunks
	return userID
}

func (fs *FileSharing) Leave(userID int) {
	delete(fs.chunks, userID)
	// Insert sorted
	idx := 0
	for idx < len(fs.released) && fs.released[idx] < userID {
		idx++
	}
	fs.released = append(fs.released, 0)
	copy(fs.released[idx+1:], fs.released[idx:])
	fs.released[idx] = userID
}

func (fs *FileSharing) Request(userID int, chunkID int) []int {
	owners := make([]int, 0)
	for uid, chunks := range fs.chunks {
		if uid != userID {
			for _, c := range chunks {
				if c == chunkID {
					owners = append(owners, uid)
					break
				}
			}
		}
	}

	// Sort owners
	sortInts(owners)

	// Requesting user gets the chunk
	fs.chunks[userID] = append(fs.chunks[userID], chunkID)

	return owners
}

func sortInts(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	fs := Constructor(4)
	user1 := fs.Join([]int{1, 2})
	fmt.Println("User1 ID:", user1)
	user2 := fs.Join([]int{2, 3})
	fmt.Println("User2 ID:", user2)
	user3 := fs.Join([]int{4})
	fmt.Println("User3 ID:", user3)

	owners := fs.Request(user2, 1)
	fmt.Println("Request chunk 1 owners:", owners)
	owners = fs.Request(user2, 4)
	fmt.Println("Request chunk 4 owners:", owners)

	fs.Leave(user1)
	user4 := fs.Join([]int{1})
	fmt.Println("User4 ID (reused):", user4)
}
