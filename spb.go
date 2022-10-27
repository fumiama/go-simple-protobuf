package spb

import "io"

func ReadNum(r io.Reader) (n uint32, cnt int, err error) {
	var buf [1]byte
	for cnt < 5 {
		_, err = r.Read(buf[:])
		if err != nil {
			return
		}
		n |= uint32(buf[0]&0x7f) << (7 * cnt)
		cnt++
		if buf[0]&0x80 == 0 {
			break
		}
	}
	return
}

/*
func WriteNum(w io.Writer, n uint32) (cnt int, err error) {
	var buf [1]byte
	for n > 0 {
		buf[0] = uint8(n & 0x7f)
		if n>>7 > 0 {
			buf[0] |= 0x80
		}
		_, err = w.Write(buf[:])
		if err != nil {
			return
		}
		n >>= 7
		cnt++
	}
	return
}
*/

type SimplePB struct {
	StructLen, RealLen uint32
	Target             [][]byte
}

func NewSimplePB(r io.Reader) (s SimplePB, err error) {
	cnt := 0
	s.StructLen, cnt, err = ReadNum(r)
	if err != nil {
		s.RealLen = uint32(cnt)
		return
	}
	if s.StructLen <= 1 || s.StructLen >= uint32(1)<<20 {
		err = ErrInvalidStructLen
		s.RealLen = uint32(cnt)
		return
	}
	var offset, datalen uint32
	n := 0
	for i := uint32(0); i < s.StructLen; i += offset {
		offset, n, err = ReadNum(r)
		cnt += n
		if err != nil {
			break
		}
		datalen, n, err = ReadNum(r)
		cnt += n
		if err != nil {
			break
		}
		if datalen == 0 {
			break
		}
		if datalen > offset {
			err = ErrInvalidDataLen
			break
		}
		t := make([]byte, datalen)
		n, err = r.Read(t)
		cnt += n
		if err != nil {
			break
		}
		s.Target = append(s.Target, t)
	}
	s.RealLen = uint32(cnt)
	return
}
