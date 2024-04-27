package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

func main() {
	total, max := 10, 3
	var wg sync.WaitGroup
	for i := 0; i < total; i += max {
		limit := max
		if i+max > total {
			limit = total - i
		}

		for j := 0; j < limit; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				conn, err := net.Dial("tcp", ":8080")
				if err != nil {
					log.Fatalf("could not dial: %v", err)
				}
				bs, err := ioutil.ReadAll(conn)
				if err != nil {
					log.Fatalf("could not read from conn: %v", err)
				}
				if string(bs) != "success" {
					log.Fatalf("request error, request: %d", i+1+j)
				}

				fmt.Printf("request %d: success\n", i+1+j)
			}()
		}
		wg.Wait()
	}
}
