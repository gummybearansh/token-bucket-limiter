package limiter;  // tells go this is a library not a executable binary (don't search for main)

import (
	"time"
	"sync"
);

type Client struct {
	Tokens float64
	LastSeen time.Time // for lazy evaluation of tokens for this ip
};


type RateLimiter struct {
	Mutex sync.Mutex
	Clients map[string]*Client
	Rate float64
	Capacity float64
}


// constructor for the RL
func NewLimiter(rate float64, capacity float64) *RateLimiter {
	rl := &RateLimiter {
		// mutex does not need to be explicitly initlaised
		Clients: make(map[string]*Client),
		Rate: rate,
		Capacity: capacity,
	}

	// detach it from here 
	// and run the sweeper 
	go rl.CleanupLoop()

	return rl
}


// called on every single incoming connection - returns bool why? 
func (rl *RateLimiter) Allow (ip string) bool {
	// first block the Map so other go routines cannot access it - AKA - lock the mutex 
	rl.Mutex.Lock()
	// if you forget to unlock - results in a deadlock 
	// use defer to unlock as soon as this functions returns no matter what happens 
	defer rl.Mutex.Unlock()

	client, exists := rl.Clients[ip]
	// this is the first time this ip has connected - create a new client
	if !exists { 
		// not a new variable update the existing one 
		client = &Client {
			Tokens: rl.Capacity, // give them the max tokens directly
			LastSeen: time.Now(),  // start their last seen time
		}

		rl.Clients[ip] = client
	}

	now := time.Now();
	// subtract current time (nanoseconds) to get Duration - convert to Seconds
	elapsed_seconds := now.Sub(client.LastSeen).Seconds()
	// rate * time many tokens accumulated 
	client.Tokens += rl.Rate * elapsed_seconds
	// clamp to ensure max is capacity 
	client.Tokens = min(client.Tokens, rl.Capacity)

	// state update 
	client.LastSeen = now;

	// now verdict - if we allow them to pass or not 
	// if they have atleast one token (to fullfill this request) we let the pass 
	if client.Tokens >= 1.0 {
		client.Tokens -= 1.0
		return true
	} else {
		return false
	}
}

func (rl *RateLimiter) CleanupLoop() {
	// infinite loop
	for {
		// run every 2 mins 
		time.Sleep(2 * time.Minute)

		// lock the mutex 
		rl.Mutex.Lock()

		for ip, client := range rl.Clients {
			// if seen over 2 mins ago - delete it from the map 
			if time.Since(client.LastSeen) > 2 * time.Minute {
				delete(rl.Clients, ip)
			}
		}

		// need to manually unlock the mutex - defer would not work 
		// it's an infinite loop will never return 
		rl.Mutex.Unlock()
	}
}

