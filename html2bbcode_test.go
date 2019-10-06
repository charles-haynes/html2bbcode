package html2bbcode

import (
	"fmt"
	"testing"
)

type test struct {
	name   string
	html   string
	bbcode string
	err    error
}

var tests = []test{
	test{
		"Plain text",
		"Lorem ipsum",
		"Lorem ipsum",
		nil,
	},
	test{
		"Line break",
		"Lorem ipsum<br />",
		"Lorem ipsum\n",
		nil,
	},
	test{
		"Paragraph",
		"<p>Lorem ipsum</p>",
		"Lorem ipsum\n",
		nil,
	},
	test{
		"Bulleted list",
		"<ul><li>Lorem</li><li>ipsum</li></ul>",
		"[*]Lorem[*]ipsum",
		nil},
	test{
		"Numbered list",
		"<ol><li>Lorem</li><li>ipsum</li></ol>",
		"[#]Lorem[#]ipsum",
		nil,
	},
	test{
		"Nested list",
		"<ol><ul><li>Lorem</li></ul><li>ipsum</li><ol><li>dolor</li></ol></ol>",
		"[*]Lorem[#]ipsum[#]dolor",
		nil,
	},
	test{
		"Naked url",
		`<a href="Lorem ipsum">Lorem ipsum</a>`,
		"Lorem ipsum",
		nil,
	},
	test{
		"img",
		`<img src="Lorem ipsum" />`,
		"[img=Lorem ipsum]",
		nil,
	},
	test{
		"desc image",
		`<img class="scale_image" onclick="lightbox.init(this, $(this).width());" alt="https://lut.im/9wZAp52yrB/0RELtSt1QzgHZIoz.jpg" src="https://redacted.ch/image.php?c=1&amp;i=https%3A%2F%2Flut.im%2F9wZAp52yrB%2F0RELtSt1QzgHZIoz.jpg" />`,
		"[img=https://lut.im/9wZAp52yrB/0RELtSt1QzgHZIoz.jpg]",
		nil},
	test{
		"img with alt",
		`<img alt="https://ptpimg.me/72r077.png" src="https://ptpimg.me/72r077.png" />`,
		`[img=https://ptpimg.me/72r077.png]`,
		nil,
	},
	test{
		"bold",
		"<strong>Lorem ipsum</strong>",
		"[b]Lorem ipsum[/b]",
		nil,
	},
	test{
		"italic",
		`<span style="font-style: italic;">Lorem ipsum</span>`,
		"[i]Lorem ipsum[/i]",
		nil,
	},
	test{
		"align center",
		`<div style="text-align: center;">Lorem ipsum</div>`,
		"[align=center]Lorem ipsum[/align]",
		nil,
	},
	test{
		"align left",
		`<div style="text-align: left;">Lorem ipsum</div>`,
		"[align=left]Lorem ipsum[/align]",
		nil,
	},
	test{
		"align right",
		`<div style="text-align: right;">Lorem ipsum</div>`,
		"[align=right]Lorem ipsum[/align]",
		nil,
	},
	test{
		"smiley heart",
		`<img border="0" src="static/common/smileys/heart.gif" alt="" />`,
		"<3",
		nil,
	},
	test{
		"smiley wave",
		`<img border="0" src="static/common/smileys/wave.gif" alt="" />`,
		":wave:",
		nil,
	},
	test{
		"smiley sad",
		`<img border="0" src="static/common/smileys/sad.gif" alt="" />`,
		":(",
		nil,
	},
	test{
		"complete entry with 'More info'",
		`<strong><span class="size4">Tracklist</span></strong><br />
<strong>01.</strong> Claude Vonstroke &amp; Eddy M - Getting Hot <span style="font-style: italic;">(3:50)</span><br />
<br />
It&#39;s &#39;Getting Hot&#39;, and Claude VonStroke and Eddy M team up to give you what you want! Eddy M hails from Barcelona, where he is a resident of the infamous El Row parties, and him and Claude are a perfect match for this funky, rolling monster.<br />
<br />
<strong>More info:</strong> <span style="display:inline-block; padding: 0px 3px;"><a rel="noreferrer" target="_blank" href="https://listen.tidal.com/album/106594181"><img width="18" class="scale_image" onclick="lightbox.init(this, $(this).width());"
alt="https://ptpimg.me/dhyvs6.png" src="https://ptpimg.me/dhyvs6.png" /> Tidal</a></span> <span style="display:inline-block; padding: 0px 3px;"><a rel="noreferrer" target="_blank" href="https://pro.beatport.com/release/getting-hot/2550383"><img width="18" class="scale_image" onclick="lightbox.init(this, $(this).width());" alt="https://ptpimg.me/26k503.png" src="https://ptpimg.me/26k503.png" /> Beatport</a></span>`,
		`[b][size=4]Tracklist[/size][/b]
[b]01.[/b] Claude Vonstroke & Eddy M - Getting Hot [i](3:50)[/i]

It's 'Getting Hot', and Claude VonStroke and Eddy M team up to give you what you want! Eddy M hails from Barcelona, where he is a resident of the infamous El Row parties, and him and Claude are a perfect match for this funky, rolling monster.

[b]More info:[/b] [pad=0|3][url=https://listen.tidal.com/album/106594181][img=18]https://ptpimg.me/dhyvs6.png[/img] Tidal[/url][/pad] [pad=0|3][url=https://pro.beatport.com/release/getting-hot/2550383][img=18]https://ptpimg.me/26k503.png[/img] Beatport[/url][/pad]`,
		nil,
	},
	test{
		"spoiler without tag",
		`<strong>Hidden text</strong>: <a href="javascript:void(0);" onclick="BBCode.spoiler(this);">Show</a><blockquote class="hidden spoiler">Lorem ipsum</blockquote>`,
		"[hide]Lorem ipsum[/hide]",
		nil,
	},
	test{
		"spoiler with tag",
		`<strong>dolor</strong>: <a href="javascript:void(0);" onclick="BBCode.spoiler(this);">Show</a><blockquote class="hidden spoiler">Lorem ipsum</blockquote>`,
		"[hide=dolor]Lorem ipsum[/hide]",
		nil,
	},
	test{
		"not spoiler",
		`<strong>dolor</strong>: <a href="javascript:void(0);" onclick="BBCode.spoiler(this);">Show</a><blockquote class="dolor">Lorem ipsum</blockquote>`,
		"[b]dolor[/b]: [url=javascript:void(0);]Show[/url][quote]Lorem ipsum[/quote]",
		nil,
	},
	test{
		"artist",
		`<a href="artist.php?artistname=Lorem ipsum">Lorem ipsum</a>`,
		"[artist]Lorem ipsum[/artist]",
		nil,
	},
	test{
		"artist error",
		`<a href="artist.php?artistname=Lorem ipsum">dolor</a>`,
		"",
		fmt.Errorf("artist tag doesn't match text, Lorem ipsum != dolor"),
	},
	test{
		"link",
		`<a rel="noreferrer" target="_blank" href="Lorem ipsum">Lorem ipsum</a>`,
		"Lorem ipsum",
		nil,
	},
	test{
		"link text",
		`<a rel="noreferrer" target="_blank" href="Lorem ipsum">dolor</a>`,
		"[url=Lorem ipsum]dolor[/url]",
		nil,
	},
	test{
		"taglist",
		`<a href="https://torrents.php?taglist=Lorem ipsum">Lorem ipsum</a>`,
		"Lorem ipsum",
		nil,
	},
	test{
		"taglist error",
		`<a href="https://torrents.php?taglist=Lorem ipsum">dolor</a>`,
		"",
		fmt.Errorf("taglist tag doesn't match text, Lorem ipsum != dolor"),
	},
	test{
		"recordlabel",
		`<a href="https://torrents.php?recordlabel=Lorem ipsum">Lorem ipsum</a>`,
		"Lorem ipsum",
		nil,
	},
	test{
		"recordlabel error",
		`<a href="https://torrents.php?recordlabel=Lorem ipsum">dolor</a>`,
		"",
		fmt.Errorf("recordlabel tag doesn't match text, Lorem ipsum != dolor"),
	},
	test{
		"color",
		`<span style="color: Lorem ipsum;">dolor</span>`,
		"[color=Lorem ipsum]dolor[/color]",
		nil,
	},
	test{
		"pre",
		`<pre>Lorem ipsum</pre>`,
		"[pre]Lorem ipsum[/pre]",
		nil,
	},
	test{
		"quote",
		`<blockquote>Lorem ipsum</blockquote>`,
		"[quote]Lorem ipsum[/quote]",
		nil,
	},
	test{
		"attributed quote",
		`<strong class="quoteheader">dolor</strong> wrote: <blockquote>Lorem ipsum</blockquote>`,
		"[quote=dolor]Lorem ipsum[/quote]",
		nil,
	},
	test{
		"linked quote",
		`<a href="#" onclick="QuoteJump(event, 'ipsum'); return false;"><strong class="quoteheader">Lorem</strong> wrote: </a><blockquote>dolor</blockquote>`,
		"[quote=Lorem]dolor[/quote]",
		nil,
	},
	test{
		"real example with padding",
		`<strong><span class="size4">Tracklist</span></strong><br />
<strong>01.</strong> Claude Vonstroke &amp; Eddy M - Getting Hot <span style="font-style: italic;">(3:50)</span><br />
<br />
It&#39;s &#39;Getting Hot&#39;, and Claude VonStroke and Eddy M team up to give you what you want! Eddy M hails from Barcelona, where he is a resident of the infamous El Row parties, and him and Claude are a perfect match for this funky, rolling monster.<br />
<br />
<strong>More info:</strong> <span style="display:inline-block; padding: 0px 3px;"><a rel="noreferrer" target="_blank" href="https://listen.tidal.com/album/106594181"><img width="18" class="scale_image" onclick="lightbox.init(this, $(this).width());" alt="https://ptpimg.me/dhyvs6.png" src="https://ptpimg.me/dhyvs6.png" /> Tidal</a></span> <span style="display:inline-block; padding: 0px 3px;"><a rel="noreferrer" target="_blank" href="https://pro.beatport.com/release/getting-hot/2550383"><img width="18" class="scale_image" onclick="lightbox.init(this, $(this).width());" alt="https://ptpimg.me/26k503.png" src="https://ptpimg.me/26k503.png" /> Beatport</a></span></div>`,
		`[b][size=4]Tracklist[/size][/b]
[b]01.[/b] Claude Vonstroke & Eddy M - Getting Hot [i](3:50)[/i]

It's 'Getting Hot', and Claude VonStroke and Eddy M team up to give you what you want! Eddy M hails from Barcelona, where he is a resident of the infamous El Row parties, and him and Claude are a perfect match for this funky, rolling monster.

[b]More info:[/b] [pad=0|3][url=https://listen.tidal.com/album/106594181][img=18]https://ptpimg.me/dhyvs6.png[/img] Tidal[/url][/pad] [pad=0|3][url=https://pro.beatport.com/release/getting-hot/2550383][img=18]https://ptpimg.me/26k503.png[/img] Beatport[/url][/pad]`,
		nil,
	},
	test{
		"padding 0 3",
		`<span style="display:inline-block; padding: 0px 3px;">Lorem ipsum</span>`,
		"[pad=0|3]Lorem ipsum[/pad]",
		nil,
	},
	test{
		"padding 5 0 0 0",
		`<span style="display:inline-block; padding: 5px 0px 0px 0px;">Lorem ipsum</span>`,
		"[pad=5|0|0|0]Lorem ipsum[/pad]",
		nil,
	},
	test{
		"forum link",
		`<a href="/forums.php?action=viewthread&threadid=dolor">Lorem ipsum</a>`,
		"Lorem ipsum",
		nil,
	},
	test{
		"important",
		`<strong class="important_text">Lorem ipsum</strong>`,
		"[important]Lorem ipsum[/important]",
		nil,
	},
}

func EqualErrors(a, b error) bool {
	if a == b {
		return true
	}
	if a != nil && b != nil && a.Error() == b.Error() {
		return true
	}
	return false
}

func Context(want, got string) (string, string) {
	if len(want)+len(got) < 60 {
		return want, got
	}
	var i int
	// find where they differ
	for i = 0; i < len(want) && i < len(got) && want[i] == got[i]; i++ {
	}
	start := i - 10
	s := "..."
	if start < 0 {
		start = 0
		s = ""
	}
	end := i + 10
	if end > len(want) {
		want = s + want[start:]
	} else {
		want = s + want[start:end] + "..."
	}
	if end > len(got) {
		got = s + got[start:]
	} else {
		got = s + got[start:end] + "..."
	}
	return want, got
}

func TestConvert(t *testing.T) {
	for _, d := range tests {
		bbcode, err := Convert(d.html)
		if !EqualErrors(err, d.err) {
			t.Errorf(`%s: want err = %v got "%s"`,
				d.name, d.err, err)
		}
		if bbcode != d.bbcode {
			want, got := Context(d.bbcode, bbcode)
			t.Errorf(`%s: want bbcode = "%s" got "%s"`,
				d.name, want, got)
		}
	}
}
