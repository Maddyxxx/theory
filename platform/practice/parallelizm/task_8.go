package parallelizm

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Задача 8
// Что выведется и как исправить?

//выведется число, близкое 1000

var wg sync.WaitGroup

func T8() {
	var counter int64
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		//запускается 1000 горутин, каждая из которых увеличивает счетчик на 1
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1) // атомарное увеличение счетчика
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
