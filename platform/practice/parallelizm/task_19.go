package parallelizm

import (
	"fmt"
	"runtime"
)
	
//Задача 19 - пропуск
//Что выведет код и почему?

//код выведет finish

func T19() {
	runtime.GOMAXPROCS(1)
	//количество потоков
	// если поток выполняется больше 10 мс, то он ставится на паузу, чтобы дать другим потокам возможность отработать
	ch := 0
	go func() {
		ch = 1
	}()

	for ch == 0 {}

	fmt.Println("finish")
}