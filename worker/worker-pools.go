package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job ", j)
		results <- j * 2
	}
}

func main() {

	start := time.Now()
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

	elapseTime := time.Since(start)
	fmt.Println("elapsed:", elapseTime)
}
