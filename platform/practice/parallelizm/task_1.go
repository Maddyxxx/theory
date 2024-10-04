package parallelizm

import (
	"fmt"
	"sync"

)

//Задача 1
//Что выведет код? Исправить все проблемы


func T1() {
	ch := make(chan int, 3)
	wg := sync.WaitGroup{}
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()	
			ch <- v * v
		}(i)
	}

	wg.Wait()
	close(ch)

	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Printf("result: %d\n", sum)
}