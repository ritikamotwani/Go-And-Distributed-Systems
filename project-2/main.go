package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweets chan *Tweet, quit chan int) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			quit <- 0
		}
		if err != nil {
			fmt.Printf("Error reading stream: %v\n", err)
			quit <- 0
		}

		tweets <- tweet
	}
}

func consumer(tweets chan *Tweet, quit chan int) {
	for {
		select {
		case tweet := <-tweets:
			if tweet.IsTalkingAboutGo() {
				fmt.Println(tweet.Username, "\ttweets talking about golang")
			} else {
				fmt.Println(tweet.Username, "\t tweets not about golang")
			}
		case <-quit:
			fmt.Println("Quitting")
			return
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	tweets := make(chan *Tweet)
	quit := make(chan int)

	// Producer
	go producer(stream, tweets, quit)

	// Consumer
	consumer(tweets, quit)

	fmt.Printf("Process took %s\n", time.Since(start))
}
