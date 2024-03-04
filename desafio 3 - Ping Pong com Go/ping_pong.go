package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"

	}
}

func pong(c chan string) {
	for {
		c <- "pong"

	}
}

func imprime(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(1 * time.Second)

	}
}

func main() {
	var c chan string = make(chan string)

	go ping(c)
	go pong(c)
	go imprime(c)
	
	var entra string
	fmt.Scanln(&entra)

}
