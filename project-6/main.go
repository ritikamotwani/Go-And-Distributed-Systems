package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64
}

const seconds_per_user = 10

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if u.IsPremium {
		process()
		return true
	}

	if atomic.LoadInt64(&u.TimeUsed) > seconds_per_user {
		return false
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		process()
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("done")
			return true
		case <-ticker.C:
			fmt.Println("tick")
			atomic.AddInt64(&u.TimeUsed, 1)
			if atomic.LoadInt64(&u.TimeUsed) > seconds_per_user {
				return false
			}
		}
	}
}

func main() {
	RunMockServer()
}
