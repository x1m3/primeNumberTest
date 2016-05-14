package main

import "fmt"
import "os"
import "strconv" 

func main() {
	
	if len(os.Args)!=4 {
		fmt.Printf("Usage %s <min> <max> [parallel|sequential]\n", os.Args[0])
		fmt.Printf("Returns the list of prime numbers between <min> and <max>.\n")
		fmt.Printf("The third parameter changes the algorithm used\n")
	} else {
		min,_ := strconv.ParseInt(os.Args[1],10,64)
		max,_ := strconv.ParseInt(os.Args[2],10,64)
		if os.Args[3]=="parallel"  {
			fmt.Printf("\nLa lista de primos es", primesInRangeParallel(min,max))
		}else if os.Args[3]=="sequential" {
			fmt.Printf("\nLa lista de primos es", primesInRange(min,max))
		} else {
			fmt.Printf("Don't know what does %s means. Valid values are parallel or sequential",os.Args[3]);
		}
	}	
}

