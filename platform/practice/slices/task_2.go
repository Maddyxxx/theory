package slices

import "fmt"
//Задача 2
//1. Что выведет код?


func t2() {
	var foo []int
	var bar []int

	foo = append(foo, 1)
	foo = append(foo, 2)
	foo = append(foo, 3)
	bar = append(foo, 4) // здесь присваивается слайсу bar значения слайса foo и к bar добавляется 4, foo = [1,2,3]
	foo = append(foo, 5) // к foo добавляется значение 5. foo = [1,2,3,5], 
	// и тк bar ссылается на foo, то bar так же получает значения [1,2,3,5]
	
	fmt.Println(foo, bar) // [1,2,3,5]  [1,2,3,5]
}

