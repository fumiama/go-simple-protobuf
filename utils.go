package spb

import "io"

func ReadNum(r io.Reader) (n uint32, cnt uint32, err error) {
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

func WriteNum(w io.Writer, n uint32) (cnt uint32, err error) {
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
