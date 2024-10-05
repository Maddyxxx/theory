package parallelizm

import (
	"fmt"
	"sync"

)

//Задача 1
//Что выведет код? Исправить все проблемы

// выведет ошибку, deadlock

func T1() {
	ch := make(chan int, 3) // канал дб буферезированный
	wg := sync.WaitGroup{}
	
	for i := 0; i < 3; i++ {
		wg.Add(1) // в цикле 
		go func(v int) {
			defer wg.Done()	
			ch <- v * v
		}(i)
	}

	wg.Wait()
	close(ch) // закрытие канала после завершения всех горутин

	var sum int
	for v := range ch {
		sum += v
	}
	fmt.Printf("result: %d\n", sum)
}