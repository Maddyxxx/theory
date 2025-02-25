package algo

import (
    "fmt"
    "reflect"
)


/**
 * Нужно написать функцию, которая принимает на вход строку,
 * а на выходе возвращает для каждого уникального символа 
 * максимальное число его беспрерывных повторений.
 * Input: aafbaaaaffaac
 * Output: a:4 b:1 f:2 c:1
 * <p>
  * Тесты
 * Input: aafbaaaaffc
 * Output: a:4 b:1 f:2 c:1
 * <p>
 * Input: bbbbbb
 * Output: b:6
 * <p>
 * Input: abc
 * Output: a:1 b:1 c:1
 * <p>
 * Input: aabbcc
 * Output: a:2 b:2 c:2
 * <p>
 * Input: aaabbbccc
 * Output: a:3 b:3 c:3
 * <p>
 * Input: a
 * Output: a:1
 * <p>
 * Input: ""
 * Output:

 **/

func unicNum(str string) map[string]int {
    mapUnic := make(map[string]int)
    counter := 1

    if str == "" {
        return mapUnic
    }

    for i := 1; i < len(str); i++ {
        if str[i] == str[i-1] {
            counter++
        } else {
            if prevCount, exists := mapUnic[string(str[i-1])]; !exists || counter > prevCount {
                mapUnic[string(str[i-1])] = counter
            }
            counter = 1
        }
    }
    if prevCount, exists := mapUnic[string(str[len(str)-1])]; !exists || counter > prevCount {
        mapUnic[string(str[len(str)-1])] = counter
    }

    return mapUnic
}

func TestUnicNum() {
    testInputs := []string{"aafbaaaaffc", "bbbbbbccbbb", "abcaabbccabc", "aabbcc", "aaabbbccc", "a", ""}
    testOutputs := []map[string]int{
        {"a": 4, "b": 1, "f": 2, "c": 1},
        {"b": 6},
        {"a": 1, "b": 1, "c": 1},
        {"a": 2, "b": 2, "c": 2},
        {"a": 3, "b": 3, "c": 3},
        {"a": 1},
        {},
    }

    for i, testInput := range testInputs {
        currOutput := unicNum(testInput)
        if !reflect.DeepEqual(currOutput, testOutputs[i]) {
            fmt.Println("test ", i+1, "failed, want ", testOutputs[i], " have ", currOutput)
        } else {
            fmt.Println("test ", i+1, "OK!")
        }
    }
}