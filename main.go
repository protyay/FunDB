package main

import (
	"fmt"
	"sync"
)

// MemTable represents an in-memory database.
type MemTable struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewMemTable creates a new empty MemTable.
func NewMemTable() *MemTable {
	return &MemTable{
		data: make(map[string]string),
	}
}

// Put inserts or updates a key-value pair in the MemTable.
func (mt *MemTable) Put(key, value string) {
	mt.mu.Lock()
	defer mt.mu.Unlock()
	mt.data[key] = value
}

// Get retrieves the value for a given key from the MemTable.
// It returns an empty string if the key does not exist.
func (mt *MemTable) Get(key string) string {
	mt.mu.RLock()
	defer mt.mu.RUnlock()
	return mt.data[key]
}

// Delete removes a key-value pair from the MemTable by key.
func (mt *MemTable) Delete(key string) {
	mt.mu.Lock()
	defer mt.mu.Unlock()
	delete(mt.data, key)
}

func main() {
	mt := NewMemTable()

	// Example usage
	mt.Put("name", "Replit")
	fmt.Println("Get 'name':", mt.Get("name")) // Output: Get 'name': Replit

	mt.Delete("name")
	fmt.Println("Get 'name' after delete:", mt.Get("name")) // Output: Get 'name' after delete:
}
