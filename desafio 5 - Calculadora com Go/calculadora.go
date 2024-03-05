package main

import "fmt"

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
func subtrai(i ...int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}
func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}
func divide(i ...int) int {
	total := 2
	for _, v := range i {
		total = v / total
	}
	return total
}

func main() {
	x := soma(1, 2, 5)
	y := multiplica(10, 3)
	w := subtrai(2, 2)
	z := divide(100)
	fmt.Println(x, y, w, z)
}
