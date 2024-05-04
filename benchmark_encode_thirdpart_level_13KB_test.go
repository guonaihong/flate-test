package flatetest

import (
	"bytes"
	"os"
	"testing"
)

func Benchmark_Encode_thirdpart_9_13KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
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

func Benchmark_Encode_thirdpart_8_13KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
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

func Benchmark_Encode_thirdpart_7_13KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
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

func Benchmark_Encode_thirdpart_6_13KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
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

func Benchmark_Encode_thirdpart_5_13KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
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

func Benchmark_Encode_thirdpart_4_13KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
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

func Benchmark_Encode_thirdpart_3_13KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
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

func Benchmark_Encode_thirdpart_1_13KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
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
