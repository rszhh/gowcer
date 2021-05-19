package reader

import (
	"bytes"
	"fmt"
	"io"
)

// MultipleReader 代表多重读取器的接口。
type MultipleReader interface {
	// Reader 用于获取一个可关闭读取器的实例。
	// 后者会持有本多重读取器中的数据。
	Reader() io.ReadCloser
}

// myMultipleReader 代表多重读取器的实现类型。
type myMultipleReader struct {
	data []byte
}

// NewMultipleReader 用于新建并返回一个多重读取器的实例。
func NewMultipleReader(reader io.Reader) (MultipleReader, error) {
	var data []byte
	var err error
	if reader != nil {
		data, err = io.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("multiple reader: couldn't create a new one: %s", err)
		}
	} else {
		data = []byte{}
	}
	return &myMultipleReader{
		data: data,
	}, nil
}

// Reader 总是返回一个新的可关闭的读取器。
// 书上应该写错了，返回的是无需关闭的
func (rr *myMultipleReader) Reader() io.ReadCloser {
	// ioutil.NopCloser函数的结果值的Close方法永远只会返回nil
	// 所以这个函数常被用于包装无需关闭的读取器
	// As of Go 1.16, this function simply calls io.NopCloser.
	return io.NopCloser(bytes.NewReader(rr.data))
}
