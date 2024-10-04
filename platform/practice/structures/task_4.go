//Задача 4
//Что выведет код и почему?

package structures

import "fmt"

type MyError struct{}

func (MyError) Error() string {
	return "MyError!"
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func T4() {
	var err *MyError
	errorHandler(err) // Error: <nil>. - Переменная err не инициализирована, то есть равна nil, 
	err = &MyError{}
	errorHandler(err) // Error: MyError! - здесь произошел высов функции Error(). Здесь err != nil, поэтому происходит вызов метода Error()
}
