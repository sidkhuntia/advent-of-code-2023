package main

import (
	"testing"
)

func BenchmarkCode(b *testing.B) {

	for i := 0; i < b.N; i++ {
		solve()
	}
}
