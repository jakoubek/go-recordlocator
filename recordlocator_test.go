package recordlocator

import (
	"testing"
)

func TestEncode(t *testing.T) {
	var cases = []struct {
		nr   int64
		want string
	}{
		{10, "D"},
		{94, "4Y"},
		{5325, "78G"},
		{10000, "CRJ"},
		{3512345, "5E82T"},
		{33554431, "ZZZZZ"},
	}
	recloc := NewRecordLocator()
	for _, c := range cases {
		t.Run(string(c.nr), func(t *testing.T) {
			have, _ := recloc.Encode(c.nr)
			if have != c.want {
				t.Errorf("got %s, want %s", have, c.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	var cases = []struct {
		str  string
		want int64
	}{
		{"D", 10},
		{"4Y", 94},
		{"78G", 5325},
		{"CRJ", 10000},
		{"5E82T", 3512345},
		{"ZZZZZ", 33554431},
	}
	recloc := NewRecordLocator()
	for _, c := range cases {
		t.Run(c.str, func(t *testing.T) {
			have, _ := recloc.Decode(c.str)
			if have != c.want {
				t.Errorf("got %d, want %d", have, c.want)
			}
		})
	}
}
