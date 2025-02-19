package main

import "testing"

func TestIsPrimeOK(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	for _, n := range primes {
		if !isPrimeV1(n) {
			t.Errorf("Expected %d to be prime", n)
		}
	}
}

func TestIsPrimeNo(t *testing.T) {
	nonPrimes := []int{1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20}
	for _, n := range nonPrimes {
		if isPrimeV1(n) {
			t.Errorf("Expected %d to be non-prime", n)
		}
	}
}

func TestIsPrimenegative(t *testing.T) {
	if isPrimeV1(-1) {
		t.Error("Expected -1 to be non-prime")
	}
}
