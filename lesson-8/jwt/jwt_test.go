package jwt

import (
	"testing"
)

var data = Payload{
	ID:     1,
	Name:   "Bob",
	IsPaid: true,
}

var dataFaster = PayloadFaster{
	ID:     1,
	Name:   "Bob",
	IsPaid: true,
}

func BenchmarkJWT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Make(data)
	}
}

func BenchmarkJWTFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeFaster(dataFaster)
	}
}
