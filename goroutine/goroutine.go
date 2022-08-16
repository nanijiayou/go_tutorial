package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func newWork() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine : i =%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func TestGoroutine() {
	fmt.Println("============= goroutine ===============")
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			fmt.Println("B called")
		}()
		fmt.Println("A called")
	}()

	// go newWork()

	// i := 0
	// for {
	// 	i++
	// 	fmt.Printf("main goroutine i = %d\n", i)
	// 	time.Sleep(1 * time.Second)
	// }
}
