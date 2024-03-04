package main

import "fmt"

func main() {

	i := 1

	for ; i <= 100; i++ {

		switch {

		case i%3 == 0 && i%5 == 0:

			fmt.Println("Pinpan")

		case i%3 == 0 && i%5 != 0:

			fmt.Println("Pin")

		case i%5 == 0 && i%3 != 0:

			fmt.Println("Pan")

		default:

			fmt.Println(i)
		}

	}
}
