## Problem 

Given an input array of integers, provide implementation of shuffle and reset functionality. 

### Solution
	- O(N) shuffle algorithm 
	- shuffle algorithm (Knuth shuffl)
	- Loop over i from 0 to len-1 of array
		- during each iteration - find a random number between 0 and i (remember its NOT 0 - len)
		- swap values at ith and r (random index 0 - i) 

