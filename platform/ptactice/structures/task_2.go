

import "fmt"
//Задача 2
//1. Добавить код, который выведет тип переменной whoami

func printType(whoami interface{}) {
	fmt.Printf("%T\n", whoami)
}

func main() {
  printType(42)
  printType("im string")
  printType(true)
}