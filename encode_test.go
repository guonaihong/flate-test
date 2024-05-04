package flatetest

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	flate2 "github.com/klauspost/compress/flate"
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

func Test_Encode3(t *testing.T) {
	// 要压缩的数据
	dat, err := os.ReadFile("./testdata/1.txt")
	if err != nil {
		panic(err.Error())
	}
	// 创建一个新的缓冲区来存放压缩后的数据
	var b bytes.Buffer

	// 创建一个新的压缩器,使用默认配置
	w, err := flate2.NewWriter(&b, flate2.DefaultCompression)
	if err != nil {
		panic(err.Error())
	}

	// 写入数据到压缩器
	_, err = w.Write(dat)
	if err != nil {
		panic(err.Error())
	}

	// 确保调用Close来终止流
	err = w.Close()
	if err != nil {
		panic(err.Error())
	}

	// b.Bytes()现在包含压缩后的数据
	fmt.Printf("before:%d, after:%d\n", len(dat), b.Len())
}
