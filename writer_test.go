package spb

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestWriter(t *testing.T) {
	f := bytes.NewBuffer(make([]byte, 0, 65536))
	ft, err := os.Open("dict.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ft.Close()
	sc := bufio.NewScanner(ft)
	for sc.Scan() {
		c := Constructor{}
		for _, s := range strings.Split(sc.Text(), "\t") {
			err = c.WriteString(s, 127)
			if err != nil {
				t.Fatal(err)
			}
		}
		c.WriteTo(f)
	}
	real, err := os.ReadFile("dict.sp")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(real, f.Bytes()) {
		t.Fail()
	}
}
