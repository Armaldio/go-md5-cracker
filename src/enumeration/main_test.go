package main

import "testing"

func BenchmarkHack1(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		hack("e2fc714c4727ee9395f324cd2e7f331f", 1)
	}
}

func BenchmarkHack2(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		hack("e2fc714c4727ee9395f324cd2e7f331f", 2)
	}
}

func BenchmarkHack3(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		hack("e2fc714c4727ee9395f324cd2e7f331f", 3)
	}
}

func BenchmarkHack(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		hack("e2fc714c4727ee9395f324cd2e7f331f", 4)
	}
}