package main

import "fmt"

// We will use this struct to communicate results via a channel
type PrimeResult struct {
	number int64 // A number
	prime bool   // Is prime or not
}


/**
 * A function to return a prime calculation over a channel. This way
 * we don't need to have 2 versions of isPrime function, one for 
 * sequential calculations and another for paralel 
 */

func isPrimeAsync(number int64, channel chan PrimeResult)  {

	result:= new (PrimeResult)
	result.number= number
	result.prime= isPrime(number)
	channel <- *result
}



/**
 * Accepts a range of integers [min, max] and a channel and it executes
 * in PARALEL the processes that check if a number is prime or not.
 * 
 * This function does nothing with the result. In another point, somebody
 * will have to read the channel and process the results
 */
func firePrimeCalculations( min int64, max int64, channel chan PrimeResult) {
var i int64
	for i=min; i<=max; i++ {
		go isPrimeAsync(i, channel)
	}
}


/** 
 * Accepts a range of integers [min, max] and
 * returns an array with all the prime numbers in this range.
 * 
 * Execution is done in paralel. First it fires all the 
 * processes that check for a prime number. These processes
 * will write the result in a channel.
 * 
 * We will receive the results over this channel creating the 
 * list of prime numbers and returning it
 * 
 */

func primesInRangeParallel( min int64, max int64) []int64 {
var  primeNumbers []int64
var res PrimeResult
var prev int64

	channel := make(chan PrimeResult)
	defer close(channel)

	go firePrimeCalculations(min, max, channel)

	prev=0
	for i:=min; i<=max; i++ {
		res = <- channel
		if res.prime {
			primeNumbers= append( primeNumbers, res.number)
			
			done := 100 * (i-min)/(max-min)
			if prev!=done {
				fmt.Printf("%d %% done.\n",done)
				prev=done
			}
		}
	}
	return primeNumbers
}


