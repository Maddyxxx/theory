package main

import (
	"fmt"
)

//Задача 3
//Что выведется?

func main() {
    c := []string{"A", "B", "D", "E"}
	b := c[1:2] // b присвается значение слайса с, b = [B]
	fmt.Println(c) // [A B D E]
    fmt.Println(b) // [B]

    b = append(b, "TT") // к b добавляется ТТ, b = [B, TT]
	fmt.Println(c) // [A B TT E] тк с содержит слайс b = [B, TT]
    fmt.Println(b) // [B TT]

	b = append(b, "HH") // к b добавляется HH, b = [B, TT, HH]
	fmt.Println(c) // [A B TT HH]
    fmt.Println(b) // [B TT HH]

	b = append(b, "PP") // к b добавляется PP, b = [B, TT, HH, PP]
	fmt.Println(c) // [A B TT HH], капа с остается 4
    fmt.Println(b) // [B TT HH PP]

}
