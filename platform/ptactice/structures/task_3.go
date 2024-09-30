package main

import "fmt"

//Задача 3
//Исправить функцию, чтобы она работала. Сигнатуру менять нельзя


func printNumber(ptrToNumber interface{}) {
        if ptrToNumber != nil { //  if ptrToNumber.(*int) != nil
                fmt.Println(*ptrToNumber.(*int))
        } else {
                fmt.Println("nil")
        }
}

func main() {
        v := 10
        printNumber(&v)
        var pv *int
        printNumber(pv)
        pv = &v
        printNumber(pv)
}