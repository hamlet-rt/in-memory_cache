# in-memory_cache

This is a simple in-memory cache implementation in Go. The cache allows you to store, retrieve, and delete key-value pairs.

## Usage

Here is an example of how to use the in-memory cache:

## Example
```go
package main

import "fmt"

func main() {
	cache := New()

	// Set a value in the cache
	cache.Set("userId", 42)

	// Get the value from the cache
	userId := cache.Get("userId")
	fmt.Println(userId) // Output: 42

	// Delete the value from the cache
	cache.Delete("userId")

	// Try to get the deleted value
	userId = cache.Get("userId")
	fmt.Println(userId) // Output: <nil>
}
```