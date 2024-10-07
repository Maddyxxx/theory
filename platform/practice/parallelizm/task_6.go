package parallelizm

import (
	"fmt"
	"sync"
)

//Задача 6
// Что выведет код и как исправить?

/*
var globalMap = map[string][]int{"test": make([]int, 0), "test2": make([]int, 0), "test3": make([]int, 0)}
var a = 0
 
func main() {
    wg := sync.WaitGroup{}
    wg.Add(3)
    go func() {
        wg.Done()
        a=10
        globalMap["test"] = append(globalMap["test"], a)
         
    }()
    go func() {
        wg.Done()
        a=11
        globalMap["test2"] = append(globalMap["test2"], a)
    }()
    go func() {
        wg.Done()
        a=12
        globalMap["test3"] = append(globalMap["test3"], a)
    }()
    wg.Wait()
    fmt.Printf("%v", globalMap)
    fmt.Printf("%d", a)
}
*/

// данный код выведет мапу globalMap с ключами test, test2 и test3, содержащими рандомные значения а (10, 11, 12)
// в зависимости от порядка выполнения горутин
// исправить можно



var globalMap = map[string][]int{"test": make([]int, 0), "test2": make([]int, 0), "test3": make([]int, 0)}
var a = make(chan int, 3)
 
func T6() {
    wg := sync.WaitGroup{}
    wg.Add(3)
    go func() {
        defer wg.Done()
        a <- 10
        globalMap["test"] = append(globalMap["test"], <- a)
    }()
    go func() {
        defer wg.Done()
        a <- 11
        globalMap["test2"] = append(globalMap["test2"], <- a)
    }()
    go func() {
        defer wg.Done()
        a <- 12
        globalMap["test3"] = append(globalMap["test3"], <- a)
    }()

    wg.Wait()
	close(a)

    fmt.Printf("%v", globalMap)
    
}

