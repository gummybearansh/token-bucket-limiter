# Token Bucket Rate Limiter

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![Algorithm](https://img.shields.io/badge/Math-Token_Bucket-FFD700?style=for-the-badge)](#)

> *A thread-safe middleware engine for throttling network traffic using the Token Bucket algorithm.*

## 🏗️ Engineering Showcase
- **Lazy Evaluation:** Calculates token math using precision nanosecond deltas only upon request, saving idle CPU cycles.
- **Concurrency Shields:** Utilizes strict `sync.Mutex` locks to guarantee zero race conditions under massive parallel traffic.
- **Asynchronous GC:** Background Goroutines automatically sweep and purge stale IP records to prevent memory leaks.

## ⚙️ Tech Stack
- **Core:** Pure Go
- **Concurrency:** `sync.Mutex` and Goroutines

## 🚀 Steps to Run
1. `go get github.com/gummybearansh/token-bucket-limiter@v1.0.0`
2. Wrap your HTTP handlers with the rate limiter instance.
3. Test physics via `go test -v`.

<br>
<div align="center">
  <img src="https://capsule-render.vercel.app/api?type=rect&color=timeGradient&height=80&text=Built%20by%20Gummybearansh:%20Building%20uncompromising%20backend%20infrastructure&fontSize=18&fontColor=ffffff&fontAlignY=50" width="100%"/>
</div>
