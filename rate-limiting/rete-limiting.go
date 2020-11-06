package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(500 * time.Microsecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 2)

	for i := 0; i < 2; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(500 * time.Microsecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	/**
	TODO 和教程的不一致，为什么？
	request 1 2020-11-06 17:32:40.8644824 +0800 CST m=+0.008000401
	request 2 2020-11-06 17:32:40.9104824 +0800 CST m=+0.054003101
	request 3 2020-11-06 17:32:40.9114824 +0800 CST m=+0.055003101
	request 4 2020-11-06 17:32:40.9124824 +0800 CST m=+0.056003201
	request 5 2020-11-06 17:32:40.9134824 +0800 CST m=+0.057003201
	request 1 2020-11-06 17:32:40.9134824 +0800 CST m=+0.057003201
	request 2 2020-11-06 17:32:40.9134824 +0800 CST m=+0.057003201
	request 3 2020-11-06 17:32:40.9144824 +0800 CST m=+0.058003301
	request 4 2020-11-06 17:32:40.9154824 +0800 CST m=+0.059003401
	request 5 2020-11-06 17:32:40.9164824 +0800 CST m=+0.060003401
	*/
}
