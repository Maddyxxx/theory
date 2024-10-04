package slices

import (
	"fmt"
)
//Задача 7
//1. Что будет содержать s после инициализации?
//2. Что произойдет в println для слайса и для мапы?


func a(s []int) {
    s = append(s, 37)
}

func b(m map[int]int) {
    m[3] = 33
}

func T7() {
    s := make([]int, 3, 8)
    m := make(map[int]int, 8)
	fmt.Println(s) // [0 0 0]
	
    // add to slice
    a(s) // s не изменится
    fmt.Println(s[3]) //ошибка тк 3го элемента не существует, append создает копию аргумента и не затрагивает оригинал

    // add to map
    b(m)
    fmt.Println(m[3]) // 33
}
