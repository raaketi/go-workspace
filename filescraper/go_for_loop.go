package main

import (
	"fmt"
	"sync"
)

func main() {
	file_name := make(chan int)
	wg := new(sync.WaitGroup)
	workers := 3
	wg.Add(2)
	go func() {
		for v := range file_name {
			fmt.Println(v)
		}
		defer wg.Done()

	}()
	for i := 0; i < workers; i++ {
		go func() {
			for i := 0; i < 1000000; i++ {
				file_name <- i
			}
			defer wg.Done()
			defer close(file_name)
		}()
	}
	go func() {
		for i := 0; i < 1000000; i++ {
			file_name <- i
		}
		defer wg.Done()
		defer close(file_name)
	}()
	// for i := 0; i < 1000000; i++ {
	// 	file_name <- i
	// }
	wg.Wait()

}
