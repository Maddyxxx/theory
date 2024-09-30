//Задача 1
//Что выведет код?
package main

import (
	"fmt"
	"sort"
)

func main() {
	v := []int{3, 4, 1, 2, 5}
	ap(v) // передается копия среза, функция добавляет значение к копии среза, поэтому
	// изменения, сделанные внутри нее, не касаются оригинального среза

	newAp(&v) // указатель ссылается на ориг срез и может его изменять

	sr(v)
	fmt.Println(v) // [1 2 3 4 5]
}

func ap(arr []int) {
	arr = append(arr, 10)
}

func newAp(arr *[]int) {
	*arr = append(*arr, 10)
}

func sr(arr []int) {
	sort.Ints(arr)
}