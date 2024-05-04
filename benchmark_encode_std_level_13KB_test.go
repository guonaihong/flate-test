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

func TestBenchmark_Encode_9_13KB(t *testing.T) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
	if err != nil {
		t.Fatal(err)
	}

	var b = writeWrap{&bytes.Buffer{}}
	w := compressNoContextTakeover(&b, 9)
	w.Write(dat)
	w.Close()
}

func TestBenchmark_Encode_8_13KB(t *testing.T) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
	if err != nil {
		t.Fatal(err)
	}

	var b = writeWrap{&bytes.Buffer{}}
	w := compressNoContextTakeover(&b, 8)
	w.Write(dat)
	w.Close()
}

func TestBenchmark_Encode_7_13KB(t *testing.T) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
	if err != nil {
		t.Fatal(err)
	}

	var b = writeWrap{&bytes.Buffer{}}
	w := compressNoContextTakeover(&b, 7)
	w.Write(dat)
	w.Close()
}

func TestBenchmark_Encode_6_13KB(t *testing.T) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
	if err != nil {
		t.Fatal(err)
	}

	var b = writeWrap{&bytes.Buffer{}}
	w := compressNoContextTakeover(&b, 6)
	w.Write(dat)
	w.Close()
}

func TestBenchmark_Encode_5_13KB(t *testing.T) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
	if err != nil {
		t.Fatal(err)
	}

	var b = writeWrap{&bytes.Buffer{}}
	w := compressNoContextTakeover(&b, 6)
	w.Write(dat)
	w.Close()
	w.Close()
}

func TestBenchmark_Encode_4_13KB(t *testing.T) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
	if err != nil {
		t.Fatal(err)
	}

	var b = writeWrap{&bytes.Buffer{}}
	w := compressNoContextTakeover(&b, 4)
	w.Write(dat)
	w.Close()
}

func TestBenchmark_Encode_3_13KB(t *testing.T) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
	if err != nil {
		t.Fatal(err)
	}

	var b = writeWrap{&bytes.Buffer{}}
	w := compressNoContextTakeover(&b, 3)
	w.Write(dat)
	w.Close()
}

func TestBenchmark_Encode_1_13KB(t *testing.T) {
	dat, err := os.ReadFile("./testdata/13KB.txt")
	if err != nil {
		t.Fatal(err)
	}

	var b = writeWrap{&bytes.Buffer{}}
	w := compressNoContextTakeover(&b, 3)
	w.Write(dat)
	w.Close()
}
