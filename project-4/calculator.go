package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type calculator struct {
	res atomic.Value
}

func (c *calculator) add(n float64) {
	c.res.Store(c.result() + n)
}

func (c *calculator) subt(n float64) {
	c.res.Store(c.result() - n)
}

func (c *calculator) mult(n float64) {
	c.res.Store(c.result() * n)
}

func (c *calculator) div(n float64) {
	if n == 0 {
		panic("division by zero")
	}
	c.res.Store(c.result() / n)
}

func (c *calculator) result() float64 {
	r, ok := c.res.Load().(float64)
	if !ok {
		panic("operating with wrong type")
	}
	return r
}

func main() {
	c := calculator{}
	c.res.Store(float64(0))
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		defer wg.Done()
		c.add(10)
	}()
	go func() {
		defer wg.Done()
		c.subt(5)
	}()
	go func() {
		defer wg.Done()
		c.div(3)
	}()
	go func() {
		defer wg.Done()
		c.mult(4)
	}()
	go func() {
		defer wg.Done()
		c.add(13)
	}()
	wg.Wait()
	fmt.Println(c.result())
}
