package spb

import "io"

type Iterator struct {
	r                  io.Reader
	structlen, reallen uint32
	i                  uint32
	v                  []byte
}

// NewReader 迭代器形式读取而非一次解析完
func NewReader(r io.Reader) (it Iterator, err error) {
	it.structlen, it.reallen, err = ReadNum(r)
	if err != nil {
		return
	}
	if it.structlen <= 1 || it.structlen >= uint32(1)<<20 {
		err = ErrInvalidStructLen
		return
	}
	it.r = r
	return
}

// Next 是否有下一个
func (it *Iterator) Next() bool {
	if it.i >= it.structlen {
		return false
	}
	var offset, datalen, n uint32
	var err error
	offset, n, err = ReadNum(it.r)
	it.reallen += n
	if err != nil {
		return false
	}
	datalen, n, err = ReadNum(it.r)
	it.reallen += n
	if err != nil {
		return false
	}
	if datalen == 0 {
		return false
	}
	if datalen > offset {
		return false
	}
	it.v = append(it.v[:0], make([]byte, datalen)...)
	x, err := it.r.Read(it.v)
	it.reallen += uint32(x)
	if err != nil {
		return false
	}
	it.i += offset
	return true
}

// Bytes 本次迭代的原始值
func (it *Iterator) Bytes() []byte {
	return it.v
}

// String 将本次迭代的值解释为 string
func (it *Iterator) String() string {
	return string(it.v)
}
