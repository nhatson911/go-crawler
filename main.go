package main

import (
	"fmt"
	"sync"
	"time"
)

func startWorker(queue chan int, name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range queue {
		fmt.Printf("Worker %s is crawling URL %d.\n", name, i)
		time.Sleep(time.Second / 10)
	}

	fmt.Printf("Worker %s has finished.\n", name)
}

func main() {
	fmt.Println("Hello, playground")

	numberOfURL := 100
	numberOfWorker := 5

	wg := new(sync.WaitGroup)
	wg.Add(numberOfWorker)

	queue := make(chan int, numberOfURL)

	go func() {
		for i := 1; i <= numberOfURL; i++ {
			queue <- i
		}
		close(queue)
	}()

	for i := 1; i <= numberOfWorker; i++ {
		go func(s string) {
			startWorker(queue, s, wg)
		}(fmt.Sprintf("%d", i))
	}

	wg.Wait()

	fmt.Println("All workers are done!!!")
}
