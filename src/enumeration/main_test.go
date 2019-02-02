package main

import "testing"

func BenchmarkHack(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		hack("e2fc714c4727ee9395f324cd2e7f331f", 4)
	}
}