package cracking_the_coding_interview

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// 在2个goroutine之间顺序的传递数字

func Test_sendNumber(t *testing.T) {
	n := 4
	ch := make(chan int, n)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()

	for i := 0; i < n; i++ {
		fmt.Println(<-ch)
	}
}

func Test_sendNumber2(t *testing.T) {
	n := 4
	var wg sync.WaitGroup
	ch := make(chan int, 1)

	for i := 0; i <= n; i++ {
		wg.Add(1)

		go func(i int) {
			ch <- i
		}(i)

		num := <-ch
		fmt.Println(num)
		wg.Done()
	}

	wg.Wait()
}

// 同1，添加了超时
func Test_sendNumber3(t *testing.T) {
	n := 4
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan int, 1)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case num := <-ch:
			fmt.Println(num)
		}
	}
}
