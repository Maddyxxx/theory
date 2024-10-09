package parallelizm

import (

	"sync/atomic"
)
//Задача 10
//1. Merge n channels
//2. Если один из входных каналов закрывается, то нужно закрыть все остальные каналы

func T10(channels ...chan int) chan int {
  out := make(chan int)
  var i int32
  atomic.StoreInt32(&i, int32(len(channels)))
  for _, c := range channels {
    go func(c <-chan int) {
      for v := range c {
        out <- v
      }
      if atomic.AddInt32(&i, -1) == 0 {
        close(out)
      }
    }(c)
  }
  return out
}