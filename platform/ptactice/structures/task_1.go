

//Задача 1
//Что выведет код?

// интерфейс - структура
// 	указатель на объект
//	таблица методов

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
	a := A() // nil
	b := B() // type impl struct{}
	var c any

	/*
	Переменные в Go всегда инициализируются значением. 
	Это относится и к интерфейсам. интерфейс реализован в виде двух элементов: тип(T) и значение(V).  
	Значение интерфейса будет nil только в том случае, если V и T оба будут nil.
	*/
	fmt.Println(a == b) //false
	fmt.Println(a == c) //true
}
