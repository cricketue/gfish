package rpc

import (
	"github.com/golang/snappy"
	"io"
	"io/util"
	"sync"
)

var snappyWriterPool sync.pool
var snappyReaderPool sync.pool

type snappyCompressor struct {
}


func (snappyCompressor) Do(w io.Writer,p []btye) error{
  z,ok := io.
}


