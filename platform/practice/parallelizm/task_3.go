package parallelizm

import (
	"fmt"
	//"sync"
   )

//Задача 3
//Будет ошибка что все горутины заблокированы. Какие горутины будут заблокированы? И почему?

func T3() {
	ch := make(chan int)
	ch <- 1
  	go func() {
    	fmt.Println(<-ch)
	}()
}
