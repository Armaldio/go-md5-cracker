package main

import "testing"

var hash = "060453b490e5d87744c3703195df2f1a" // not found, 8 char long

func BenchmarkHack1(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		hack(hash, 1)
	}
}

func BenchmarkHack2(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		hack(hash, 2)
	}
}

func BenchmarkHack3(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		hack(hash, 3)
	}
}

func BenchmarkHack4(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		hack(hash, 4)
	}
}