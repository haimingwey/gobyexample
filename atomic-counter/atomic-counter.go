package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var num uint64

	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)

		go func() {
			for i := 0; i < 1000; i++ {
				// output 500000
				atomic.AddUint64(&num, 1)
				// output 451121,
				// 可以用go run -race atomic-counter.go 获取数据竞争失败的详情
				//num++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("num:", num)
}
