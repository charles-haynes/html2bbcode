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
	test{"Bulleted list", "<ul><li>Lorem</li><li>ipsum</li></ul>", "[*]Lorem[*]ipsum", nil},
	test{"Numbered list", "<ol><li>Lorem</li><li>ipsum</li></ol>", "[#]Lorem[#]ipsum", nil},
	test{"Nested list", "<ol><ul><li>Lorem</li></ul><li>ipsum</li><ol><li>dolor</li></ol></ol>", "[*]Lorem[#]ipsum[#]dolor", nil},
	test{"Naked url", `<a href="Lorem ipsum">Lorem ipsum</a>`, "Lorem ipsum", nil},
	test{"img", `<img src="Lorem ipsum" />`, "[img=Lorem ipsum]", nil},
	test{
		"desc image",
		`<img class="scale_image" onclick="lightbox.init(this, $(this).width());" alt="https://lut.im/9wZAp52yrB/0RELtSt1QzgHZIoz.jpg" src="https://redacted.ch/image.php?c=1&amp;i=https%3A%2F%2Flut.im%2F9wZAp52yrB%2F0RELtSt1QzgHZIoz.jpg" />`,
		"[img=https://lut.im/9wZAp52yrB/0RELtSt1QzgHZIoz.jpg]",
		nil},
}

func TestConvert(t *testing.T) {
	for _, d := range tests {
		bbcode, err := Convert(d.html)
		if err != d.err {
			t.Errorf(`%s: want err = %v got "%s"`,
				d.name, d.err, err)
		}
		if bbcode != d.bbcode {
			t.Errorf(`%s: want bbcode = "%s" got "%s"`,
				d.name, d.bbcode, bbcode)
		}
	}
}
