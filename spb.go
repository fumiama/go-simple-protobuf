package spb

import "io"

type SimplePB struct {
	StructLen, RealLen uint32
	Target             [][]byte
}

func NewSimplePB(r io.Reader) (s SimplePB, err error) {
	cnt := uint32(0)
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
	n := uint32(0)
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
		x := 0
		x, err = r.Read(t)
		cnt += uint32(x)
		if err != nil {
			break
		}
		s.Target = append(s.Target, t)
	}
	s.RealLen = uint32(cnt)
	return
}
