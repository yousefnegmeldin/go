package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "function 1 done"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "function 2 done"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "function 3 done"
	}()

	// func() {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("function 1 done")

	// }()

	// func() {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("function 2 done")
	// }()

	// func() {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("function 3 done")
	// }()

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("main function took %s", elapsed)

}
