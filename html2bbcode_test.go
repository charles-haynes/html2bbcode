package html2bbcode

import (
	"testing"
)

type test struct {
	name   string
	html   string
	bbcode string
	err    error
}

var tests = []test{
	test{"Plain text", "Lorem ipsum", "Lorem ipsum", nil},
}

func TestConvert(t *testing.T) {
	for _, d := range tests {
		bbcode, err := Convert(d.html)
		if err != d.err {
			t.Errorf("%s: want err = %s got %s",
				d.name, d.err, err)
		}
		if bbcode != d.bbcode {
			t.Errorf("%s: want bbcode = %s got %s",
				d.name, d.bbcode, bbcode)
		}
	}
}
