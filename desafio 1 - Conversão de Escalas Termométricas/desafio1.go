package main

import "fmt"

const ebulicaoK = 373

func main() {
	
	tempK := ebulicaoK
	
	temperaturaC := tempK - 273

	fmt.Printf("A temperatura de ebulição da água de %v K em Celsius é %v °C", tempK,temperaturaC)

}