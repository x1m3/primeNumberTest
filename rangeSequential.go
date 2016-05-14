package main

import "fmt"


/** 
 * Accepts a range of integers [min, max] and
 * returns an array with all the prime numbers in this range.
 * 
 * Execution is done sequentially
 */

func primesInRange( min int64, max int64) []int64 {
var primeNumbers []int64
var prev int64

	prev=0
	for i:=min; i<=max; i++ {
		if isPrime(i) {
			primeNumbers= append(primeNumbers, i)
			
			// -- Stupid code that shows the progress to the user. 
			done := 100 * (i-min)/(max-min)
			if prev!=done {
				fmt.Printf("%d %% done.\n",done)
				prev=done
			}
		}
	}
	return primeNumbers
}
