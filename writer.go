package spb

import (
	"bytes"
	"io"
)

type Constructor struct {
	structlen, reallen uint32
	body               bytes.Buffer
}

// NewWriter 新的写入构造器
func NewWriter(buf []byte) (c Constructor) {
	c.body = *bytes.NewBuffer(buf)
	return
}

// WriteString 按字符串写入一项
func (c *Constructor) WriteString(s string, cap uint32) error {
	if int(cap) < len(s) {
		return ErrInvalidDataLen
	}
	cnt, err := WriteNum(&c.body, cap)
	if err != nil {
		return err
	}
	c.structlen += cap
	c.reallen += cnt
	cnt, err = WriteNum(&c.body, uint32(len(s)))
	if err != nil {
		return err
	}
	c.reallen += cnt
	_, err = c.body.WriteString(s)
	if err != nil {
		return err
	}
	c.reallen += uint32(len(s))
	return nil
}

// Cap structlen 目前写入的总长, 带 padding
func (c *Constructor) Cap() uint32 {
	return c.structlen
}

// Len reallen 目前写入的实值总长, 不带 padding
func (c *Constructor) Len() uint32 {
	return c.reallen
}

// WriteTo 将当前结果写出, 返回写入的实际长度
func (c *Constructor) WriteTo(w io.Writer) (int64, error) {
	cnt, err := WriteNum(w, c.structlen)
	if err != nil {
		return int64(cnt), err
	}
	n, err := w.Write(c.body.Bytes())
	cnt += uint32(n)
	return int64(cnt), err
}
