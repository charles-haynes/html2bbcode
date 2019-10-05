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
	test{"Line break", "Lorem ipsum<br />", "Lorem ipsum\n", nil},
	test{"Paragraph", "<p>Lorem ipsum</p>", "Lorem ipsum\n", nil},
}

func TestConvert(t *testing.T) {
	for _, d := range tests {
		bbcode, err := Convert(d.html)
		if err != d.err {
			t.Errorf("%s: want err = %v got %s",
				d.name, d.err, err)
		}
		if bbcode != d.bbcode {
			t.Errorf("%s: want bbcode = %s got %s",
				d.name, d.bbcode, bbcode)
		}
	}
}
