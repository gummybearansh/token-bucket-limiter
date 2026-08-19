# Token Bucket Rate Limiter

A high-performance, zero-dependency middleware engine for throttling network traffic in Go.

Built from scratch to understand the low-level memory and concurrency physics of network proxies, this library implements the **Token Bucket** algorithm using lazy evaluation and strict mutex locking, capable of handling thousands of concurrent goroutines without data races or memory leaks.

## Core Architecture

- **Lazy Evaluation:** Instead of wasting CPU cycles constantly refilling tokens in a background loop, token math is calculated dynamically at the exact microsecond a request arrives using precise `time.Duration` nanosecond deltas.
- **Concurrency Shield:** A global `sync.Mutex` acts as a cooperative traffic light, guaranteeing zero race conditions when thousands of sockets attempt to update the memory map simultaneously.
- **Asynchronous Garbage Collection:** A detached background daemon (Goroutine) wakes up every 2 minutes to perform an $O(N)$ sweep, automatically deleting stale IP addresses to prevent memory leaks over long-running server uptimes.

## Installation

```bash
go get [github.com/gummybearansh/token-bucket-limiter@v1.0.0](https://github.com/gummybearansh/token-bucket-limiter@v1.0.0)
```

## Quick Start

This engine is designed to sit directly inside your TCP/HTTP `Accept()` loop.

```go
package main

import (
	"fmt"
	"net/http"
	"[github.com/gummybearansh/token-bucket-limiter](https://github.com/gummybearansh/token-bucket-limiter)"
)

func main() {
	// Initialize the global brain:
	// Rate: 1.0 token per second
	// Capacity: Burst limit of 3 tokens
	engine := limiter.NewLimiter(1.0, 3.0)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Extract user IP (In production, strip the port or check X-Forwarded-For)
		ip := r.RemoteAddr

		// Ask the engine for the verdict
		if !engine.Allow(ip) {
			http.Error(w, "429 Too Many Requests", http.StatusTooManyRequests)
			return
		}

		fmt.Fprintln(w, "Request Allowed: Welcome to the server.")
	})

	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
```

## Testing

The test suite rigorously verifies the burst capacity, the exact cut-off boundary, and the time-delta refill mathematics. To run the proving grounds:

```bash
go test -v
```

## Curriculum Context

This project is Phase 03 of the **Sovereign Architect Curriculum**—a masterclass in building high-performance distributed systems, network infrastructure, and advanced data physics from absolute scratch.
