package slices

import (
	"fmt"
)

//Задача 6
//1. Что выведется?
//2. Зная обо всех таких нюансах, которые могут возникнуть, какие есть рекомендации?

// Вариант 1

func mod1(a []int) {
  for i := range a {
    a[i] = 5 // все элементы слайса становятся равными 5
  }
  fmt.Printf("mod1: %d\n", a)
}

func printMod1() {
  sl := []int{1, 2, 3, 5} 
  fmt.Printf("original sl 1: %d\n", sl) // [1 2 3 5]
  mod1(sl)
  fmt.Printf("printMod1: %d\n\n", sl) // [5 5 5 5]
}

// Вариант 2

func mod2(a []int) {
  for i := range a {
    a[i] = 5
  }
  fmt.Printf("mod2: %d\n", a) // [5 5 5 5]
}

func printMod2() {
  sl := make([]int, 4, 8)
  sl[0] = 1
  sl[1] = 2
  sl[2] = 3
  sl[3] = 5
  fmt.Printf("original sl 2: %d\n", sl) // [1 2 3 5]
  mod2(sl)
  fmt.Printf("printMod2: %d\n\n", sl) // [5 5 5 5]
}

// Вариант 3

func mod3(a []int) {
  a = append(a, 125) // append создает новый слайс с копией элементов из оригинально слайса + 125
  for i := range a {
    a[i] = 5
  }
  fmt.Printf("mod3: %d\n", a) // [5 5 5 5 5]
}

func printMod3() {
  sl := make([]int, 4, 8) // создание слайса 
  sl[0] = 1
  sl[1] = 2
  sl[2] = 3
  sl[3] = 5
  fmt.Printf("original sl 3: %d\n", sl) // [1 2 3 5]
  mod3(sl)
  fmt.Printf("printMod3: %d\n\n", sl) // [5 5 5 5]
}

// Вариант 4

func mod4(a []int) {
  a = append(a, 125) // append создает новый слайс с элементами оригинального слайса, поэтому изменения в функции не затрагивают ориг слайс
  for i := range a {
    a[i] = 5 
  }
  fmt.Printf("mod4: %d\n", a) // [5 5 5 5 5 5]
}

func printMod4() {
  sl := []int{1, 2, 3, 4, 5}
  fmt.Printf("original sl 4: %d\n", sl) // [1 2 3 4 5]
  mod4(sl)
  fmt.Printf("printMod4: %d\n\n", sl) // [1 2 3 4 5] - функция mod4() не изменяет слайс
}

func PrintMod(){
	/*
	Рекомендации по работе со слайсами:
	1. не меняйте размер слайса из входных параметров и не ведите работу с ним напрямую если ожидается изменение размера
	2. при изменении размера слайса создается другой слайс с копией элементов оригинального
	3. хотите сделать работу со слайсом - создавайте копию и явно возвращайте как результат функции 
	*/
	printMod1()
	printMod2()
	printMod3()
	printMod4()
}

