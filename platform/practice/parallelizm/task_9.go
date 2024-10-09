package parallelizm

import (
	"fmt"
)
//Задача 9
//1. Что выведется и как исправить?
//2. Что поправить, чтобы сохранить порядок?

// ничего не выведется
// в 27 строке "go fmt.Println(<-ch)" убрать go

func T9() {
  m := make(chan string, 3)
  cnt := 5
  for i := 0; i < cnt; i++ {
    go func() {
      m <- fmt.Sprintf("Goroutine %d", i)
    }()
  }

  for i := 0; i < cnt; i++ {
    ReceiveFromCh(m)
  }
}

func ReceiveFromCh(ch chan string) {
  fmt.Println(<-ch)
}