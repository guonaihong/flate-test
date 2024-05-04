package flatetest

import (
	"bytes"
	"io"
	"os"
	"testing"
)

type writeWrap struct {
	io.Writer
}

func (w *writeWrap) Close() error {
	return nil
}

func Benchmark_Encode_9(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
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

func Benchmark_Encode_8(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
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

func Benchmark_Encode_7(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
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

func Benchmark_Encode_6(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
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

func Benchmark_Encode_5(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
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

func Benchmark_Encode_4(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
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

func Benchmark_Encode_3(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
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

func Benchmark_Encode_1(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1.txt")
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
