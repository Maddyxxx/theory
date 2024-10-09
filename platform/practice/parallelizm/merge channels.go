package parallelizm

import (
	//"fmt"
	"sync"
)
//Задача 10
//1. Merge n channels
//2. Если один из входных каналов закрывается, то нужно закрыть все остальные каналы

func makeChannels() []chan int {
	var chs []chan int
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)

	ch2 := make(chan int, 3)
	ch2 <- 3
	ch2 <- 4
	ch2 <- 5
	close(ch2)

	return append(chs, ch1, ch2)
}

func case3(channels ...chan int) chan int {
	//var a atomic.AddInt64()
	var wg sync.WaitGroup
	out := make(chan int)
	   
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}
	   
	for _, ch := range channels {
		// ввести атомик
		// если канал закрыт, уменьшить атомик
		wg.Add(1)
		go output(ch)
	}
	   
	go func() {
		wg.Wait()
		close(out)
	}()

	return out  
}

//func T10() {
//	for v := range case3(makeChannels()...) {
//		fmt.Printf("merged channel: %d\n", v)
//	}
//}