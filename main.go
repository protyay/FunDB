package main

import (
	"errors"
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
// It returns an error if the key is empty.
func (mt *MemTable) Put(key , value string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	mt.mu.Lock()
	defer mt.mu.Unlock()
	mt.data[key] = value
	return nil
}

// Get retrieves the value for a given key from the MemTable.
// It returns an empty string if the key does not exist or is empty.
func (mt *MemTable) Get(key string) string {
	if key == "" {
		return ""
	}
	mt.mu.RLock()
	defer mt.mu.RUnlock()
	return mt.data[key]
}

// Delete removes a key-value pair from the MemTable by key.
// If the key is empty, it does nothing.
func (mt *MemTable) Delete(key string) {
	if key == "" {
		return
	}
	mt.mu.Lock()
	defer mt.mu.Unlock()
	delete(mt.data, key)
}

func main() {
	mt := NewMemTable()

	// Example usage
	if err := mt.Put("name", "Replit"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Put 'name' successfully")
	}

	fmt.Println("Get 'name':", mt.Get("name")) // Output: Get 'name': Replit

	if err := mt.Put("", "EmptyKey"); err != nil {
		fmt.Println("Error:", err) // Output: Error: key cannot be empty
	}

	mt.Delete("name")
	fmt.Println("Get 'name' after delete:", mt.Get("name")) // Output: Get 'name' after delete:
}