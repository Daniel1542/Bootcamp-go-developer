package main

import "fmt"

func main() {
	
	i := 1
	
	for ; i <= 100; i++	 {

		if  i%3 == 0 && i%5 == 0{	

			fmt.Printf("Pinpan %v", i)
			fmt.Println()
			
		}else if i%3 == 0 && i%5 != 0{
		
			fmt.Printf("Pin %v", i)	
			fmt.Println()
			
		}else if i%5 == 0 && i%3 != 0{

			fmt.Printf("Pan %v", i)	
			fmt.Println()

		}else{	

			fmt.Println(i)
			fmt.Println()
			continue		
			
		}		

	}
}
