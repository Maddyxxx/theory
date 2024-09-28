// задача на знание внутренней структуры интерфейсов

// интерфейс - структура
// 	указатель на объект
//	таблица методов

package main

import "fmt"

type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil

}
func B() I {
	var ret *impl
	return ret
}

func main() {
	a := A()
	b := B()
	// interface != interface
	fmt.Println(a == b)
}
