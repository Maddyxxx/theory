package slices

import (
	"fmt"
)

//Задача 4
//1. Что выведет код

// выведет ошибку, тк на 18 строчке итерация по нулевой мапе. нужно исправить создание мапы на 13 строке

func T4() {
  m := make(map[string]int) // было var m map[string]int
  for _, word := range []string{"hello", "world", "from", "the",
    "best", "language", "in", "the", "world"} {
    m[word]++
  }
  for k, v := range m {
    fmt.Println(k, v) // вывод неупорядоченный тк структура мап не сортирована
  }
}
