package base58_test

import (
	"testing"
	"crypto/rand"
	"log"
	"github.com/akamensky/base58"
)

func getRandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}
	return b
}

func benchmarkEncode(n int, b *testing.B) {
	b.StopTimer()
	b.SetBytes(int64(n))
	data := make([]([]byte), b.N, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = getRandomBytes(n)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		base58.Encode(data[i])
	}
}

func benchmarkDecode(n int, b *testing.B) {
	b.StopTimer()
	b.SetBytes(int64(n))
	data := make([]string, b.N, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = base58.Encode(getRandomBytes(n))
	}
	b.StartTimer()
	
	for i := 0; i < b.N; i++ {
		base58.Decode(data[i])
	}
}

func BenchmarkEncode32(b *testing.B) {
	benchmarkEncode(32, b)
}

func BenchmarkEncode64(b *testing.B) {
	benchmarkEncode(64, b)
}

func BenchmarkEncode128(b *testing.B) {
	benchmarkEncode(128, b)
}

func BenchmarkEncode256(b *testing.B) {
	benchmarkEncode(256, b)
}

func BenchmarkEncode512(b *testing.B) {
	benchmarkEncode(512, b)
}

func BenchmarkDecode32(b *testing.B) {
	benchmarkDecode(32, b)
}

func BenchmarkDecode64(b *testing.B) {
	benchmarkDecode(64, b)
}

func BenchmarkDecode128(b *testing.B) {
	benchmarkDecode(128, b)
}

func BenchmarkDecode256(b *testing.B) {
	benchmarkDecode(256, b)
}

func BenchmarkDecode512(b *testing.B) {
	benchmarkDecode(512, b)
}
