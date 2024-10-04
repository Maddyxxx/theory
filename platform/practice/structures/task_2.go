
package structures

//Задача 2
//1. Добавить код, который выведет тип переменной whoami

func printType(whoami interface{}) {
	// fmt.Printf("%T\n", whoami)
}

func T2() {
  printType(42)
  printType("im string")
  printType(true)
}