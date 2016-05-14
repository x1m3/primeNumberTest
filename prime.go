package main

import "math"

/**
 * A non very optimal function that checks if a number is prime
 * 
 * A prime number is one that only is divisible by 1 or itself
 * 
 * The algorithm is pure brute force, iterating by numbers smaller 
 * than the candidate and calculating the modulus
 * 
 * We only use 3 optimizations. 
 *  1) If the number is different than 2, but divisible by 2, it is not a prime. 
 *  2) Then, we only test for odd divisors (3,5,7,9, etc..)
 *  3) We only check for numbers smaller or equal the root square
 *     of the candidate. It is imposible to have divisors greater
 *     than this value
 */
func isPrime(candidate int64) (bool) {
var i,limit int64

	if candidate==2 {
		return true
	}

	if math.Mod(float64(candidate), 2)==0 {
		return false;
	}


	limit = int64( math.Ceil( math.Sqrt(float64(candidate) )))
	for i=3; i<=limit; i+=2 { //Only odd numbers
		if math.Mod(float64(candidate), float64(i))==0 { 
			return false 
		}
	}
	return true
}
