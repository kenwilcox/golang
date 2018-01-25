package stringutil

import "testing"

func TestReverse(t *testing.T) {
  cases := []struct {
    in, want string
  }{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
  }
  for _, c := range cases {
    got := Reverse(c.in)
    if got != c.want {
      t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
    }
  }
}

func BenchmarkReverse(b *testing.B) {
	str := "Hello World"
	for i := 0; i < b.N; i++ {
		str = Reverse(str)
		//str += "!"
	}
}