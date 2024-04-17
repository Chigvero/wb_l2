package main

import (
	"fmt"
)

func or(channels ...<-chan interface{}) <-chan interface{} {

	result := make(chan interface{})

	go func() {
		defer close(result)

		done := make(chan struct{})
		defer close(done)

		for _, ch := range channels {
			go func(ch <-chan interface{}) {
				select {
				case val, ok := <-ch:
					if ok {
						select {
						case result <- val:
						case <-done:
						}
					} else {
						return
					}
				case <-done:
					return
				}
			}(ch)
		}
	}()

	return result
}

func main() {

	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	result := or(ch1, ch2, ch3)

	close(ch2)

	for val := range result {
		fmt.Println("Received:", val)
	}
}
