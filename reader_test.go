package spb

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestReader(t *testing.T) {
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
	var s Iterator
	i := 1
	for sc.Scan() {
		s, err = NewReader(f)
		if err != nil {
			break
		}
		if !s.Next() {
			t.Fatal("unexpected no next")
		}
		t0 := s.String()
		if !s.Next() {
			t.Fatal("unexpected no next")
		}
		t1 := s.String()
		if fmt.Sprint(t0, "\t", t1) != sc.Text() {
			t.Fatal("invalid text @ line", i)
		}
		i++
	}
}
