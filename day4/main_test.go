package main

import (
	"testing"
)

func bankai(b *testing.B) {

	for i := 0; i < b.N; i++ {
		solve()
	}
}
