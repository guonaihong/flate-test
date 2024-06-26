package flatetest

import (
	"bytes"
	"os"
	"testing"
)

func Benchmark_Encode_WithPool_9_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	out := make([]byte, 0, len(dat))
	buf := bytes.NewBuffer(out)
	for i := 0; i < b.N; i++ {
		var b = writeWrap{buf}
		w := compressNoContextTakeover(&b, 9)
		w.Write(dat)
		w.Close()
		buf.Reset()
	}
}

func Benchmark_Encode_WithPool_8_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	out := make([]byte, 0, len(dat))
	buf := bytes.NewBuffer(out)
	for i := 0; i < b.N; i++ {

		var b = writeWrap{buf}
		w := compressNoContextTakeover(&b, 8)
		w.Write(dat)
		w.Close()
		buf.Reset()
	}
}

func Benchmark_Encode_WithPool_7_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	out := make([]byte, 0, len(dat))
	buf := bytes.NewBuffer(out)
	for i := 0; i < b.N; i++ {
		var b = writeWrap{buf}
		w := compressNoContextTakeover(&b, 7)
		w.Write(dat)
		w.Close()
		buf.Reset()
	}
}

func Benchmark_Encode_WithPool_6_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	out := make([]byte, 0, len(dat))
	buf := bytes.NewBuffer(out)
	for i := 0; i < b.N; i++ {
		var b = writeWrap{buf}
		w := compressNoContextTakeover(&b, 6)
		w.Write(dat)
		w.Close()
		buf.Reset()
	}
}

func Benchmark_Encode_WithPool_5_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	out := make([]byte, 0, len(dat))
	buf := bytes.NewBuffer(out)
	for i := 0; i < b.N; i++ {
		var b = writeWrap{buf}
		w := compressNoContextTakeover(&b, 6)
		w.Write(dat)
		w.Close()
		w.Close()
		buf.Reset()
	}
}

func Benchmark_Encode_WithPool_4_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	out := make([]byte, 0, len(dat))
	buf := bytes.NewBuffer(out)
	for i := 0; i < b.N; i++ {
		var b = writeWrap{buf}
		w := compressNoContextTakeover(&b, 4)
		w.Write(dat)
		w.Close()
		buf.Reset()
	}
}

func Benchmark_Encode_WithPool_3_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	out := make([]byte, 0, len(dat))
	buf := bytes.NewBuffer(out)
	for i := 0; i < b.N; i++ {
		var b = writeWrap{buf}
		w := compressNoContextTakeover(&b, 3)
		w.Write(dat)
		w.Close()
		buf.Reset()
	}
}

func Benchmark_Encode_WithPool_1_1KB(b *testing.B) {
	dat, err := os.ReadFile("./testdata/1KB.txt")
	if err != nil {
		panic(err.Error())
	}
	out := make([]byte, 0, len(dat))
	buf := bytes.NewBuffer(out)
	for i := 0; i < b.N; i++ {
		var b = writeWrap{buf}
		w := compressNoContextTakeover(&b, 1)
		w.Write(dat)
		w.Close()
		buf.Reset()
	}
}
