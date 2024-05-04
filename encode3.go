package flatetest

// Copyright 2017 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"errors"
	"io"
	"sync"

	flate2 "github.com/klauspost/compress/flate"
)

const (
	minCompressionLevel2     = -2 // flate.HuffmanOnly not defined in Go < 1.6
	maxCompressionLevel2     = flate2.BestCompression
	defaultCompressionLevel2 = 1
)

var (
	flateWriterPools2 [maxCompressionLevel - minCompressionLevel + 1]sync.Pool
)

/*
func isValidCompressionLevel(level int) bool {
	return minCompressionLevel <= level && level <= maxCompressionLevel
}
*/

func compressNoContextTakeover2(w io.WriteCloser, level int) io.WriteCloser {
	p := &flateWriterPools2[level-minCompressionLevel2]
	tw := &truncWriter2{w: w}
	fw, _ := p.Get().(*flate2.Writer)
	if fw == nil {
		fw, _ = flate2.NewWriter(tw, level)
	} else {
		fw.Reset(tw)
	}
	return &flateWriteWrapper2{fw: fw, tw: tw, p: p}
}

// truncWriter is an io.Writer that writes all but the last four bytes of the
// stream to another io.Writer.
type truncWriter2 struct {
	w io.WriteCloser
	n int
	p [4]byte
}

func (w *truncWriter2) Write(p []byte) (int, error) {
	n := 0

	// fill buffer first for simplicity.
	if w.n < len(w.p) {
		n = copy(w.p[w.n:], p)
		p = p[n:]
		w.n += n
		if len(p) == 0 {
			return n, nil
		}
	}

	m := len(p)
	if m > len(w.p) {
		m = len(w.p)
	}

	if nn, err := w.w.Write(w.p[:m]); err != nil {
		return n + nn, err
	}

	copy(w.p[:], w.p[m:])
	copy(w.p[len(w.p)-m:], p[len(p)-m:])
	nn, err := w.w.Write(p[:len(p)-m])
	return n + nn, err
}

type flateWriteWrapper2 struct {
	fw *flate2.Writer
	tw *truncWriter2
	p  *sync.Pool
}

func (w *flateWriteWrapper2) Write(p []byte) (int, error) {
	if w.fw == nil {
		return 0, ErrWriteClosed
	}
	return w.fw.Write(p)
}

func (w *flateWriteWrapper2) Close() error {
	if w.fw == nil {
		return ErrWriteClosed
	}
	err1 := w.fw.Flush()
	w.p.Put(w.fw)
	w.fw = nil
	if w.tw.p != [4]byte{0, 0, 0xff, 0xff} {
		return errors.New("websocket: internal error, unexpected bytes at end of flate stream")
	}
	err2 := w.tw.w.Close()
	if err1 != nil {
		return err1
	}
	return err2
}
