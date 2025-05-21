package WB

import "fmt"

/*
Дан диапазон портов min и max. Также есть список уже занятых портов busy.
Нужно написать функцию, возвращающую все диапазоны портов внутри min и max

### Пример
min = 30000
max = 32000
busy = [30100, 30200]

expected = [[30000, 30099], [30101, 30199], [30201, 32000]]
*/

func rangeFreePorts(min int, max int, busy []int) [][]int {
	var rangePorts = make([][]int, 0)
	for _, port := range busy {
		var ports []int
		if port == min {
			min++
		} else if port < min {
			continue
		} else if port < max {
			ports = append(ports, min, port-1)
			rangePorts = append(rangePorts, ports)
			min = port + 1
		}
	}
	if min <= max {
		if max == busy[len(busy)-1] {
			max = min
		}
		ports := []int{min, min}

		rangePorts = append(rangePorts, ports)
	}
	return rangePorts
}

func T1() {
	min_ := 20001
	max_ := 20003
	busy := []int{20001, 20002, 20003, 20005}
	ports := rangeFreePorts(min_, max_, busy)
	fmt.Println(ports)

}
