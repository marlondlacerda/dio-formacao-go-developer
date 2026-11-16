package main

import (
	"fmt"
	"sync"
)

func main() {
	ping := make(chan struct{})
	pong := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			<-ping
			fmt.Println("ping")
			pong <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			<-pong
			fmt.Println("pong")

			if i < 4 {
				ping <- struct{}{}
			}
		}
	}()

	ping <- struct{}{}

	wg.Wait()
}
