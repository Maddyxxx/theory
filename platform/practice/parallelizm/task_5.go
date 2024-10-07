package parallelizm

import (
	"fmt"
	"time"

)

// Задача 5
// 1. Как будет работать код?
// 2. Как сделать так, чтобы выводился только первый ch?

// код будет печатать в консоль значение из того канала, который прочитает первым
// случайным образом




func T5() {
    ch := make(chan bool)
    ch2 := make(chan bool)
    ch3 := make(chan bool)

    go func() {
        ch <- true
    }()

    go func() {
		time.Sleep(2 * time.Second) // ничего лучше не придумал как поставить тайм слип)
    	ch2 <- true
    }()

    go func() {
		time.Sleep(2 * time.Second)
        ch3 <- true
    }()

    select {
    case <-ch:
        fmt.Printf("val from ch")
    case <-ch2:
        fmt.Printf("val from ch2")
    case <-ch3:
        fmt.Printf("val from ch3")
    }
}
