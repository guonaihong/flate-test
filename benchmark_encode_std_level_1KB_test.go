package flatetest

import (
	"bytes"
	"os"
	"testing"
)

func Benchmark_Encode_9_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover(&b, 9)
		w.Write(dat)
		w.Close()
	}
}

func Benchmark_Encode_8_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover(&b, 8)
		w.Write(dat)
		w.Close()
	}
}

func Benchmark_Encode_7_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover(&b, 7)
		w.Write(dat)
		w.Close()
	}
}

func Benchmark_Encode_6_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover(&b, 6)
		w.Write(dat)
		w.Close()
	}
}

func Benchmark_Encode_5_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover(&b, 6)
		w.Write(dat)
		w.Close()
		w.Close()
	}
}

func Benchmark_Encode_4_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover(&b, 4)
		w.Write(dat)
		w.Close()
	}
}

func Benchmark_Encode_3_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover(&b, 3)
		w.Write(dat)
		w.Close()
	}
}

func Benchmark_Encode_1_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover(&b, 1)
		w.Write(dat)
		w.Close()
	}
}
