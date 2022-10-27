package spb

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestSPB(t *testing.T) {
	f, err := os.Open("dict.sp")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	ft, err := os.Open("dict.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer ft.Close()
	sc := bufio.NewScanner(ft)
	var s SimplePB
	i := 1
	for sc.Scan() {
		s, err = NewSimplePB(f)
		if err != nil {
			break
		}
		if len(s.Target) != 2 {
			t.Fatal("invalid target")
		}
		if fmt.Sprint(string(s.Target[0]), "\t", string(s.Target[1])) != sc.Text() {
			t.Fatal("invalid text @ line", i)
		}
		i++
	}
}
