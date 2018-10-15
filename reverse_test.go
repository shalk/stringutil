package stringutil

import "testing"


func TestReverse(t *testing.T) {
    cases := []struct {
	    in, want string
	}{
		{"",""},
		{"hello","olleh"},
		{"haha","ahah"},
		{"你好世界","界世好你"},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
