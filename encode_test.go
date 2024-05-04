package flatetest

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func Test_Encode(t *testing.T) {

	dat, err := os.ReadFile("./testdata/1.txt")
	if err != nil {
		panic(err.Error())
	}

	var out bytes.Buffer
	var b = writeWrap{&out}
	w := compressNoContextTakeover(&b, 1)
	_, err = w.Write(dat)
	if err != nil {
		panic(err.Error())
	}
	w.Close()

	fmt.Printf("before:%d, after:%d\n", len(dat), out.Len())
}

func Test_Encode2(t *testing.T) {

	dat, err := os.ReadFile("./testdata/1.txt")
	if err != nil {
		panic(err.Error())
	}

	buf := make([]byte, 0, len(dat))
	out := bytes.NewBuffer(buf)
	var b = writeWrap{out}
	w := compressNoContextTakeover(&b, 1)
	_, err = w.Write(dat)
	if err != nil {
		panic(err.Error())
	}
	w.Close()

	fmt.Printf("before:%d, after:%d\n", len(dat), out.Len())
	out.Reset()

	w = compressNoContextTakeover(&b, 1)
	_, err = w.Write(dat)
	if err != nil {
		panic(err.Error())
	}
	w.Close()

	fmt.Printf("before:%d, after:%d\n", len(dat), out.Len())
}
