package parallelizm

import (
	"fmt"
	"sync"
   )

//Задача 2
//Что выведет код? Должны выводится все значения

// почему то выводится примерно половина значений
func T2() {
	a := 5000
	for i := 0; i < a; i++ {
		go fmt.Println(i)
	}
}
   
func T2Edit() {
	a := 5000
	var wg sync.WaitGroup
   
	for i := 0; i < a; i++ {
		wg.Add(1) // Увеличиваем счетчик ожидания
		go func(i int) {
			defer wg.Done() // Уменьшаем счетчик после завершения горутины
			fmt.Println(i)
		}(i) // Передаем значение i в горутину
	}
	wg.Wait() // Ожидаем завершения всех горутин
   }