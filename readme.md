# Token Bucket Rate Limiter 🚦

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![Concurrency](https://img.shields.io/badge/Concurrency-Goroutines-336699?style=for-the-badge)](#)
[![Data Structures](https://img.shields.io/badge/Math-Token_Bucket-FFD700?style=for-the-badge)](#)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

A high-performance, zero-dependency middleware engine for throttling network traffic in Go. Built from scratch to understand the low-level memory and concurrency physics of network proxies.

This library implements the classic **Token Bucket** algorithm using lazy evaluation and strict mutex locking, capable of handling thousands of concurrent goroutines without data races or memory leaks.

## 🚀 The Tech Real: How it Works

At the core of this system is a mathematically precise throttling engine:

- **Lazy Evaluation:** Instead of wasting CPU cycles constantly refilling tokens in an infinite background loop, token math is calculated dynamically at the exact microsecond a request arrives using precise `time.Duration` nanosecond deltas.
- **Concurrency Shield:** A global `sync.Mutex` acts as a cooperative traffic light, guaranteeing zero race conditions when thousands of sockets attempt to update the memory map simultaneously.
- **Asynchronous Garbage Collection:** A detached background daemon (Goroutine) wakes up every 2 minutes to perform an $O(N)$ sweep, automatically deleting stale IP addresses to prevent memory leaks over long-running server uptimes.

## ⚡ Tech Stack

- **Backend System:** Pure `Go`.
- **Synchronization Engine:** `sync.Mutex` ensures memory-safe concurrency.
- **Time/Math Engine:** `time` package provides nanosecond precision delta evaluations.

## 🛠️ Quick Start

This engine is designed to sit directly inside your TCP/HTTP `Accept()` loop.

```bash
# Install the library
go get github.com/gummybearansh/token-bucket-limiter@v1.0.0
```

## 🎮 Interacting with the System

Initialize the global brain and enforce it on a standard route:

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/gummybearansh/token-bucket-limiter"
)

func main() {
	// Initialize the global brain:
	// Rate: 1.0 token per second | Capacity: Burst limit of 3 tokens
	engine := limiter.NewLimiter(1.0, 3.0)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if !engine.Allow(ip) { // Ask the engine for the verdict
			http.Error(w, "429 Too Many Requests", http.StatusTooManyRequests)
			return
		}
		fmt.Fprintln(w, "Request Allowed: Welcome to the server.")
	})

	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
```

**Testing the Protocol:**
The test suite rigorously verifies the burst capacity, the exact cut-off boundary, and the time-delta refill mathematics. Run the proving grounds using:
```bash
go test -v
```

<br>
<div align="center">
  <img src="https://capsule-render.vercel.app/api?type=waving&color=timeGradient&height=120&section=footer&text=Built%20by%20Gummybearansh:%20Building%20uncompromising%20backend%20infrastructure&fontSize=24&fontAlignY=38&animation=fadeIn" width="100%"/>
</div>
