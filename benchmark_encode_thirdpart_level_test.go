package flatetest

import (
	"bytes"
	"os"
	"testing"
)

func Benchmark_Encode_3_9(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover2(&b, 9)
		w.Write(dat)
		w.Close()
	}
}

func Benchmark_Encode_3_8(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover2(&b, 8)
		w.Write(dat)
		w.Close()
	}
}

func Benchmark_Encode_3_7(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover2(&b, 7)
		w.Write(dat)
		w.Close()
	}
}

func Benchmark_Encode_3_6(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover2(&b, 6)
		w.Write(dat)
		w.Close()
	}
}

func Benchmark_Encode_3_5(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover2(&b, 6)
		w.Write(dat)
		w.Close()
		w.Close()
	}
}

func Benchmark_Encode_3_4(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover2(&b, 4)
		w.Write(dat)
		w.Close()
	}
}

func Benchmark_Encode_3_3(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover2(&b, 3)
		w.Write(dat)
		w.Close()
	}
}

func Benchmark_Encode_3_1(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var b = writeWrap{&bytes.Buffer{}}
		w := compressNoContextTakeover2(&b, 1)
		w.Write(dat)
		w.Close()
	}
}
