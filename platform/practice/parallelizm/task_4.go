package parallelizm

import (
	"fmt"

)

// Задача 4
// 1. Как это работает, что не так, что поправить?

// В вашем коде есть несколько проблем, связанных с порядком операций с каналом.
// Вы пытаетесь отправить значение в канал до того, как кто-либо его сможет прочитать,
// что приводит к блокировке. 

// Горутина запускается до того, как мы отправляем значение в канал.


func T4() {
	ch := make(chan bool)

	go func() {
		for v := range ch {
			fmt.Println(v) // проитерироваться по каналу чтобы получить значения
		}
		
	}()

	ch <- true 
	close(ch) // закрыть канал после записи вв него всех значений
	
}

