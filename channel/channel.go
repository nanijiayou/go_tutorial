package channel

import (
	"fmt"
	"time"
)

func fib(c1, c2 chan int) {
	x, y := 1, 1
	for {
		// select可以监控多个channle状态
		select {
		case c1 <- x:
			x = y
			y = x + y
		case <-c2:
			fmt.Println("close")
			return
		}
	}
}

func TestChan() {
	fmt.Println("============= channle ===============")

	// 创建：无缓冲
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine end")
		fmt.Println("goroutine running")
		for i := 0; i < 3; i++ {
			c <- i
		}

		// 关闭
		close(c)
	}()

	// num := <-c
	// fmt.Println("main goroutine get num = ", num)

	// for {
	// 	if data, ok := <-c; ok {
	// 		fmt.Println("get data = ", data)
	// 	} else {
	// 		break
	// 	}
	// }

	// 使用range来不断迭代channel操作, 和上面for效果一致
	for data := range c {
		fmt.Println("get data = ", data)
	}

	// 有缓冲
	c2 := make(chan int, 3)
	fmt.Println("len(c2) = ", len(c2), " cap(c2) = ", cap(c2))

	go func() {
		defer fmt.Println("sub goroutine end")
		for i := 0; i < 3; i++ {
			c2 <- i
			fmt.Println("send ele = ", i, "len(c2) = ", len(c2), " cap(c2) = ", cap(c2))
		}
	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 3; i++ {
		num2 := <-c2
		fmt.Println("num2 = ", num2)
	}

	c3 := make(chan int)
	c4 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c3)
		}
		c4 <- 0
	}()

	// select可以监控多个channle状态
	fib(c3, c4)

	fmt.Println("main goroutine end")
}
