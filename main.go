package main

import (
	"fmt"
	"time"
)

func Craw(routineId int, url int) {
	fmt.Println("Routine ID: ", routineId, " craw url: ", url)
}

func main() {
	fmt.Printf("Crawler \n")
	numURL := 100
	numChan := 5
	numRoutine := 10

	ch := make(chan int, numChan)

	go func() {
		for i := 1; i <= numURL; i++ {
			fmt.Println("Send URL: ", i)
			ch <- i
		}
		fmt.Println("Sent")
	}()

	for i := 1; i <= numRoutine; i++ {
		go func(i int, ch chan int) {
			for {
				x := <-ch
				Craw(i, x)
			}
		}(i, ch)
	}

	time.Sleep(time.Second * 10)
}
