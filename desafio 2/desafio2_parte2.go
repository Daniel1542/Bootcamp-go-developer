package main

import "fmt"

func main() {
	
	i := 1
	
	for ; i <= 100; i++	 {

		if  i%3 == 0 && i%5 == 0{	

			fmt.Println("Pinpan")
			fmt.Println()
			
		}else if i%3 == 0 && i%5 != 0{
		
			fmt.Println("Pin")	
			fmt.Println()
			
		}else if i%5 == 0 && i%3 != 0{

			fmt.Println("Pan")	
			fmt.Println()

		}else{	
			
			fmt.Println(i)
			fmt.Println()		
			
		}		

	}
}
