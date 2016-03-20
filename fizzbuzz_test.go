package gofizzbuzz

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "1"},
		{2, "2"},
		{3, "fizz"},
		{4, "4"},
		{5, "buzz"},
		{6, "fizz"},
		{10, "buzz"},
		{12, "fizz"},
		{15, "fizzbuzz"},
		{16, "16"},
	}
	for _, c := range cases {
		got := GoFizzBuzz(c.in)
		if got != c.want {
			message := "GoFizzBuzz(%q) == %q, want %q"
			t.Errorf(message, fmt.Sprintf("%d", c.in), got, c.want)
		}
	}
}
