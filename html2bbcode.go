package html2bbcode

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

/*
		// When removing elements, you have to iterate over the list backwards
		for ($i = $Elements->length - 1; $i >= 0; $i--) {
			$Element = $Elements->item($i);
			if (strpos($Element->getAttribute('class'), 'size') !== false) {
				$NewElement = $Document->createElement('size', $Element->nodeValue);
				$NewElement->setAttribute('size', str_replace('size', '', $Element->getAttribute('class')));
				$Element->parentNode->replaceChild($NewElement, $Element);
			}
			elseif (strpos($Element->getAttribute('style'), 'font-style: italic') !== false) {
				$NewElement = $Document->createElement('italic', $Element->nodeValue);
				$Element->parentNode->replaceChild($NewElement, $Element);
			}
			elseif (strpos($Element->getAttribute('style'), 'text-decoration: underline') !== false) {
				$NewElement = $Document->createElement('underline', $Element->nodeValue);
				$Element->parentNode->replaceChild($NewElement, $Element);
			}
			elseif (strpos($Element->getAttribute('style'), 'color: ') !== false) {
				$NewElement = $Document->createElement('color', $Element->nodeValue);
				$NewElement->setAttribute('color', str_replace(array('color: ', ';'), '', $Element->getAttribute('style')));
				$Element->parentNode->replaceChild($NewElement, $Element);
			}
		}

		$Elements = $Document->getElementsByTagName('ul');
		for ($i = 0; $i < $Elements->length; $i++) {
			$InnerElements = $Elements->item($i)->getElementsByTagName('li');
			for ($j = $InnerElements->length - 1; $j >= 0; $j--) {
				$Element = $InnerElements->item($j);
				$NewElement = $Document->createElement('bullet', $Element->nodeValue);
				$Element->parentNode->replaceChild($NewElement, $Element);
			}
		}

		$Elements = $Document->getElementsByTagName('ol');
		for ($i = 0; $i < $Elements->length; $i++) {
			$InnerElements = $Elements->item($i)->getElementsByTagName('li');
			for ($j = $InnerElements->length - 1; $j >= 0; $j--) {
				$Element = $InnerElements->item($j);
				$NewElement = $Document->createElement('number', $Element->nodeValue);
				$Element->parentNode->replaceChild($NewElement, $Element);
			}
		}

		$Elements = $Document->getElementsByTagName('strong');
		for ($i = $Elements->length - 1; $i >= 0; $i--) {
			$Element = $Elements->item($i);
			if ($Element->hasAttribute('class') === 'important_text') {
				$NewElement = $Document->createElement('important', $Element->nodeValue);
				$Element->parentNode->replaceChild($NewElement, $Element);
			}
		}

		$Elements = $Document->getElementsByTagName('a');
		for ($i = $Elements->length - 1; $i >= 0; $i--) {
			$Element = $Elements->item($i);
			if ($Element->hasAttribute('href')) {
				$Element->removeAttribute('rel');
				$Element->removeAttribute('target');
				if ($Element->getAttribute('href') === $Element->nodeValue) {
					$Element->removeAttribute('href');
				}
				elseif ($Element->getAttribute('href') === 'javascript:void(0);'
					&& $Element->getAttribute('onclick') === 'BBCode.spoiler(this);') {
					$Spoilers = $Document->getElementsByTagName('blockquote');
					for ($j = $Spoilers->length - 1; $j >= 0; $j--) {
						$Spoiler = $Spoilers->item($j);
						if ($Spoiler->hasAttribute('class') && $Spoiler->getAttribute('class') === 'hidden spoiler') {
							$NewElement = $Document->createElement('spoiler', $Spoiler->nodeValue);
							$Element->parentNode->replaceChild($NewElement, $Element);
							$Spoiler->parentNode->removeChild($Spoiler);
							break;
						}
					}
				}
				elseif (substr($Element->getAttribute('href'), 0, 22) === 'artist.php?artistname=') {
					$NewElement = $Document->createElement('artist', $Element->nodeValue);
					$Element->parentNode->replaceChild($NewElement, $Element);
				}
				elseif (substr($Element->getAttribute('href'), 0, 30) === 'user.php?action=search&search=') {
					$NewElement = $Document->createElement('user', $Element->nodeValue);
					$Element->parentNode->replaceChild($NewElement, $Element);
				}
			}
		}

		$Str = str_replace(array("<body>\n", "\n</body>", "<body>", "</body>"), "", $Document->saveHTML($Document->getElementsByTagName('body')->item(0)));
		$Str = str_replace(array("\r\n", "\n"), "", $Str);
		$Str = preg_replace("/\<strong\>([a-zA-Z0-9 ]+)\<\/strong\>\: \<spoiler\>/", "[spoiler=\\1]", $Str);
		$Str = str_replace("</spoiler>", "[/spoiler]", $Str);
		$Str = preg_replace("/\<strong class=\"quoteheader\"\>(.*)\<\/strong\>(.*)wrote\:(.*)\<blockquote\>/","[quote=\\1]", $Str);
		$Str = preg_replace("/\<(\/*)blockquote\>/", "[\\1quote]", $Str);
		$Str = preg_replace("/\<(\/*)strong\>/", "[\\1b]", $Str);
		$Str = preg_replace("/\<(\/*)italic\>/", "[\\1i]", $Str);
		$Str = preg_replace("/\<(\/*)underline\>/", "[\\1u]", $Str);
		$Str = preg_replace("/\<(\/*)important\>/", "[\\1important]", $Str);
		$Str = preg_replace("/\<color color=\"(.*)\"\>/", "[color=\\1]", $Str);
		$Str = str_replace("</color>", "[/color]", $Str);
		$Str = str_replace(array('<number>', '<bullet>'), array('[#]', '[*]'), $Str);
		$Str = str_replace(array('</number>', '</bullet>'), '<br />', $Str);
		$Str = str_replace(array('<ul class="postlist">', '<ol class="postlist">', '</ul>', '</ol>'), '', $Str);
		$Str = preg_replace("/\<size size=\"([0-9]+)\"\>/", "[size=\\1]", $Str);
		$Str = str_replace("</size>", "[/size]", $Str);
		//$Str = preg_replace("/\<a href=\"rules.php\?(.*)#(.*)\"\>(.*)\<\/a\>/", "[rule]\\3[/rule]", $Str);
		//$Str = preg_replace("/\<a href=\"wiki.php\?action=article&name=(.*)\"\>(.*)\<\/a>/", "[[\\1]]", $Str);
		$Str = preg_replace('#/torrents.php\?recordlabel="?(?:[^"]*)#', 'https://'.SITE_URL.'\\0', $Str);
		$Str = preg_replace('#/torrents.php\?taglist="?(?:[^"]*)#', 'https://'.SITE_URL.'\\0', $Str);
		$Str = preg_replace("/\<(\/*)artist\>/", "[\\1artist]", $Str);
		$Str = preg_replace("/\((\/*)user\>/", "[\\1user]", $Str);
		$Str = preg_replace("/\<a href=\"([^\"]*)\">/", "[url=\\1]", $Str);
		$Str = preg_replace("/\<(\/*)a\>/", "[\\1url]", $Str);
		$Str = preg_replace("/\<img(.*)src=\"(.*)\"(.*)\>/", '[img]\\2[/img]', $Str);
		$Str = str_replace('<p>', '', $Str);
		$Str = str_replace('</p>', '<br />', $Str);
		//return $Str;
		return str_replace(array("<br />", "<br>"), "\n", $Str);
	}
}*/

func GetAttr(n *html.Node, key string) (string, error) {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val, nil
		}
	}
	return "", fmt.Errorf("attr %s not found", key)
}

type BBCode struct {
	strings.Builder
	lists []string // stack of nested list types
}

func (bc *BBCode) Node(n *html.Node, tag string) error {
	bc.WriteString("[")
	bc.WriteString(tag)
	bc.WriteString("]")
	if err := bc.convertChildren(n); err != nil {
		return err
	}
	bc.WriteString("[/")
	bc.WriteString(tag)
	bc.WriteString("]")
	return nil
}

func (bc *BBCode) NodeData(n *html.Node, tag string) error {
	if n.FirstChild != nil {
		return fmt.Errorf("expected node %s not to have children", tag)
	}
	bc.WriteString("[")
	bc.WriteString(tag)
	bc.WriteString("]")
	bc.WriteString(n.Data)
	bc.WriteString("[/")
	bc.WriteString(tag)
	bc.WriteString("]")
	return nil
}

func (bc *BBCode) NodeLiteral(n *html.Node, tag string, l string) error {
	if n.FirstChild != nil {
		return fmt.Errorf("expected node %s not to have children", tag)
	}
	bc.WriteString("[")
	bc.WriteString(tag)
	bc.WriteString("]")
	bc.WriteString(l)
	bc.WriteString("[/")
	bc.WriteString(tag)
	bc.WriteString("]")
	return nil
}

func (bc *BBCode) NodeVal(n *html.Node, tag, v string) error {
	bc.WriteString("[")
	bc.WriteString(tag)
	bc.WriteString("=")
	bc.WriteString(v)
	bc.WriteString("]")
	if err := bc.convertChildren(n); err != nil {
		return err
	}
	bc.WriteString("[/")
	bc.WriteString(tag)
	bc.WriteString("]")
	return nil
}

func (bc *BBCode) NodeValData(n *html.Node, tag, v string) error {
	if n.FirstChild != nil {
		return fmt.Errorf("expected node %s not to have children", tag)
	}
	bc.WriteString("[")
	bc.WriteString(tag)
	bc.WriteString("=")
	bc.WriteString(v)
	bc.WriteString("]")
	bc.WriteString(n.Data)
	bc.WriteString("[/")
	bc.WriteString(tag)
	bc.WriteString("]")
	return nil
}

var smileyMap = map[string]string{
	"angry.gif":   `:angry:`,
	"biggrin.gif": `:-D`,
	// "biggrin.gif":   `:D`,
	"blank.gif": `:|`,
	// "blank.gif":     `:-|`,
	"blush.gif":  `:blush:`,
	"cool.gif":   `:cool:`,
	"crying.gif": `:'(`,
	// "crying.gif":    `:crying:`,
	"eyesright.gif": `>>`,
	"frown.gif":     `:frown:`,
	"heart.gif":     `<3`,
	"hmm.gif":       `:unsure:`,
	// "hmm.gif":       `:\`,
	"ilu.gif":      `:whatlove:`,
	"laughing.gif": `:lol:`,
	"loveflac.gif": `:loveflac:`,
	// "loveflac.gif":  `:flaclove:`,
	"ninja.gif":  `:ninja:`,
	"no.gif":     `:no:`,
	"nod.gif":    `:nod:`,
	"ohnoes.gif": `:ohno:`,
	// "ohnoes.gif":    `:ohnoes:`,
	"omg.gif":    `:omg:`,
	"ohshit.gif": `:o`,
	// "ohshit.gif":    `:O`,
	"paddle.gif": `:paddle:`,
	"sad.gif":    `:(`,
	// "sad.gif":       `:-(`,
	"shifty.gif": `:shifty:`,
	"sick.gif":   `:sick:`,
	"smile.gif":  `:)`,
	// "smile.gif":     `:-)`,
	"sorry.gif":  `:sorry:`,
	"thanks.gif": `:thanks:`,
	"tongue.gif": `:P`,
	// "tongue.gif":    `:p`,
	// "tongue.gif":    `:-P`,
	// "tongue.gif": `:-p`,
	"wave.gif": `:wave:`,
	"wink.gif": `;-)`,
	// "wink.gif":      `:wink:`,
	"creepy.gif":  `:creepy:`,
	"worried.gif": `:worried:`,
	"wtf.gif":     `:wtf:`,
	"wub.gif":     `:wub:`,
}

func (bc *BBCode) Img(n *html.Node) error {
	src, err := GetAttr(n, "src")
	if err != nil {
		return err
	}
	alt, _ := GetAttr(n, "alt")
	if border, _ := GetAttr(n, "border"); border == "0" &&
		strings.HasPrefix(src, "static/common/smileys/") {
		smiley := strings.TrimPrefix(src, "static/common/smileys/")
		txt, ok := smileyMap[smiley]
		if !ok {
			return fmt.Errorf("unknown smiley %s", smiley)
		}
		bc.WriteString(txt)
		return nil
	}
	if class, _ := GetAttr(n, "class"); class == "scale_image" {
		if onclick, _ := GetAttr(n, "onclick"); onclick != "lightbox.init(this, $(this).width());" {
			return fmt.Errorf("img class is scale_image but no onclick")
		}
		if n.FirstChild == nil {
			if width, err := GetAttr(n, "width"); err == nil {
				bc.WriteString("[img=")
				bc.WriteString(width)
				bc.WriteString("]")
				bc.WriteString(alt)
				bc.WriteString("[/img]")
				return nil
			}
			bc.WriteString("[img=")
			bc.WriteString(alt)
			bc.WriteString("]")
			return nil
		}
	}
	if n.FirstChild == nil {
		bc.WriteString("[img=")
		bc.WriteString(src)
		bc.WriteString("]")
		return nil
	}
	bc.Node(n, "img")
	return nil
}

func Text(n *html.Node) string {
	if n == nil {
		return ""
	}
	if n.Type != html.TextNode {
		return ""
	}
	return n.Data
}

func (bc *BBCode) Blockquote(n *html.Node) error {
	if PartOfHidden(n) {
		strong := n.PrevSibling.PrevSibling.PrevSibling
		if tag := Text(strong.FirstChild); tag != "Hidden text" {
			return bc.NodeVal(n, "hide", tag)
		}
		return bc.Node(n, "hide")
	}
	if PartOfLinkedQuote(n) {
		// we ignore link info since they won't work across sites
		// a := n.PrevSibling
		// onclick, _ := GetAttr(a, "onclick")
		// link := LinkedQuoteRE.FindStringSubmatch(onclick)[1]
		strong := n.PrevSibling.FirstChild
		// tag := Text(strong.FirstChild) + "|" + link
		tag := Text(strong.FirstChild)
		return bc.NodeVal(n, "quote", tag)
	}
	if PartOfAttributedQuote(n) {
		strong := n.PrevSibling.PrevSibling
		tag := Text(strong.FirstChild)
		return bc.NodeVal(n, "quote", tag)
	}
	return bc.Node(n, "quote")
}

func (bc *BBCode) Hr(n *html.Node) error {
	return nil
}

func ParseStyle(style string) (sk, sv string, err error) {
	ss := strings.Split(style, ":")
	if len(ss) != 2 {
		err = fmt.Errorf("can't parse style %s", style)
		return sk, sv, err
	}
	sk, sv = strings.TrimSpace(ss[0]), strings.TrimSpace(ss[1])
	return sk, sv, err
}

func (bc *BBCode) SpanStyle(n *html.Node, v string) error {
	pad := false
	for _, style := range strings.Split(v, ";") {
		if style == "" {
			continue // skip empty
		}
		sk, sv, err := ParseStyle(style)
		if err != nil {
			return err
		}
		switch sk {
		case "font-style":
			if string(sv) != "italic" {
				return fmt.Errorf("unknown font-style %s", sv)
			}
			return bc.Node(n, "i")
		case "text-decoration":
			if string(sv) != "underline" {
				return fmt.Errorf("unknown text-decoration %s", sv)
			}
			return bc.Node(n, "underline")
		case "color":
			return bc.NodeVal(n, "color", sv)
		case "display":
			if sv == "inline-block" {
				pad = true
			} else {
				return fmt.Errorf("unknown display %s", sv)
			}
		case "padding":
			if !pad {
				return fmt.Errorf("unexpected padding %s", style)
			}
			re := regexp.MustCompile("([0-9]+)px")
			m := re.FindAllStringSubmatch(sv, -1)
			var p []string
			for i := range m {
				p = append(p, m[i][1])
			}
			return bc.NodeVal(n, "pad", strings.Join(p, "|"))
		default:
			return fmt.Errorf(`unknown span style "%s"`, style)
		}
	}
	return nil
}

func (bc *BBCode) Span(n *html.Node) error {
	for _, a := range n.Attr {
		switch a.Key {
		case "class":
			if len(a.Val) < 5 || a.Val[:4] != `size` {
				return fmt.Errorf("unknown span class %s", a.Val)
			}
			return bc.NodeVal(n, "size", a.Val[4:])
		case "style":
			return bc.SpanStyle(n, a.Val)
		}
	}
	return fmt.Errorf("unknown span")
}

func (bc *BBCode) DivStyle(n *html.Node, v string) error {
	for _, style := range strings.Split(v, ";") {
		sk, sv, err := ParseStyle(style)
		if err != nil {
			return err
		}
		if sk == "text-align" {
			switch sv {
			case "center", "left", "right":
				return bc.NodeVal(n, "align", sv)
			default:
				return fmt.Errorf(`unknown text align "%s"`, sv)
			}
		}
	}
	return fmt.Errorf("unknown div style %s", v)
}

func (bc *BBCode) Div(n *html.Node) error {
	style, err := GetAttr(n, "style")
	if err != nil {
		return err
	}
	return bc.DivStyle(n, style)
}

func (bc *BBCode) Strong(n *html.Node) error {
	class, _ := GetAttr(n, "class")
	switch class {
	case "important_text":
		return bc.Node(n, "important")
	case "quote":
		return bc.Node(n, "quote")
	default:
		if PartOfHidden(n) || PartOfAttributedQuote(n) {
			return nil
		}
		return bc.Node(n, "b")
	}
}

func AssertElement(n *html.Node, a atom.Atom) error {
	if n == nil {
		return fmt.Errorf("expected element node, got nil")
	}
	if n.Type != html.ElementNode {
		return fmt.Errorf("expected element node, got %d", n.Type)

	}
	if n.DataAtom != a {
		return fmt.Errorf("expected %s element, got %s", a, n.DataAtom)
	}
	return nil
}

func AssertText(n *html.Node, t string) error {
	if n == nil {
		return fmt.Errorf("expected text node, got nil")
	}
	if n.Type != html.TextNode {
		return fmt.Errorf("expected text node, got %d", n.Type)

	}
	if n.Data != t {
		return fmt.Errorf("expected text %s, got %s", t, n.Data)
	}
	return nil
}

func Parent(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	return n.Parent
}

func FirstChild(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	return n.FirstChild
}

func Prev(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	return n.PrevSibling
}

func Next(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	return n.NextSibling
}

func PartOfHidden(n *html.Node) bool {
	if n == nil {
		return false
	}
	var strong, colon, a, blockquote *html.Node
	switch n.Type {
	case html.ElementNode:
		switch n.DataAtom {
		case atom.Strong:
			strong = n
		case atom.A:
			strong = Prev(Prev(n))
		case atom.Blockquote:
			strong = Prev(Prev(Prev(n)))
		default:
			return false

		}
	case html.TextNode:
		strong = Prev(n)
	default:
		return false
	}
	colon = Next(strong)
	a = Next(colon)
	blockquote = Next(a)
	if strong == nil || colon == nil || a == nil || blockquote == nil {
		return false
	}
	if err := AssertElement(strong, atom.Strong); err != nil {
		return false
	}
	if colon.Data != ": " {
		return false
	}
	if href, _ := GetAttr(a, "href"); href != "javascript:void(0);" {
		return false
	}
	if onclick, _ := GetAttr(a, "onclick"); onclick !=
		"BBCode.spoiler(this);" {
		return false
	}
	if err := AssertElement(blockquote, atom.Blockquote); err != nil {
		return false
	}
	if class, _ := GetAttr(blockquote, "class"); class != "hidden spoiler" {
		return false
	}
	return true
}
func PartOfAttributedQuote(n *html.Node) bool {
	if n == nil {
		return false
	}
	var strong, wrote, blockquote *html.Node
	switch n.Type {
	case html.ElementNode:
		switch n.DataAtom {
		case atom.Strong:
			strong = n
		case atom.Blockquote:
			strong = Prev(Prev(n))
		default:
			return false

		}
	case html.TextNode:
		strong = Prev(n)
	default:
		return false
	}
	wrote = Next(strong)
	blockquote = Next(wrote)
	if strong == nil || wrote == nil || blockquote == nil {
		return false
	}
	if err := AssertElement(strong, atom.Strong); err != nil {
		return false
	}
	if Text(strong.FirstChild) == "" {
		return false
	}
	if wrote.Data != " wrote: " {
		return false
	}
	if err := AssertElement(blockquote, atom.Blockquote); err != nil {
		return false
	}
	return true
}

var LinkedQuoteRE = regexp.MustCompile(
	`^QuoteJump\(event, '([^']+)'\); return false;$`)

func PartOfLinkedQuote(n *html.Node) bool {
	if n == nil {
		return false
	}
	var a *html.Node
	switch n.Type {
	case html.ElementNode:
		switch n.DataAtom {
		case atom.A:
			a = n
		case atom.Strong:
			a = Parent(n)
		case atom.Blockquote:
			a = Prev(n)
		default:
			return false

		}
	case html.TextNode:
		a = Parent(n)
	default:
		return false
	}
	strong := FirstChild(a)
	wrote := Next(strong)
	blockquote := Next(a)
	if a == nil || strong == nil || wrote == nil || blockquote == nil {
		return false
	}
	if href, _ := GetAttr(a, "href"); href != "#" {
		return false
	}
	if onclick, _ := GetAttr(a, "onclick"); !LinkedQuoteRE.MatchString(onclick) {
		return false
	}
	if err := AssertElement(strong, atom.Strong); err != nil {
		return false
	}
	if Text(strong.FirstChild) == "" {
		return false
	}
	if wrote.Data != " wrote: " {
		return false
	}
	if err := AssertElement(blockquote, atom.Blockquote); err != nil {
		return false
	}
	return true
}

func (bc *BBCode) A(n *html.Node) error {
	href, err := GetAttr(n, "href")
	if err != nil {
		return err
	}
	switch true {
	case PartOfHidden(n) || PartOfLinkedQuote(n) || PartOfLinkedQuote(n):
		return nil
	case strings.HasPrefix(href, "artist.php?artistname="):
		a := strings.TrimPrefix(href, "artist.php?artistname=")
		if t := Text(n.FirstChild); a != t {
			return fmt.Errorf(
				"artist tag doesn't match text, %s != %s",
				a, t)
		}
		bc.Node(n, "artist")
		return nil
	case strings.HasPrefix(href, "/user.php?action=search&search="):
		return bc.NodeData(n, "user")
	case strings.HasPrefix(href, "/forums.php?action=viewthread&threadid="):
		// strip forum links, they won't work cross site
		return bc.convertChildren(n)
	case strings.HasPrefix(href, "/requests.php?action=view&id="):
		// strip request links, they won't work cross site
		return bc.convertChildren(n)
	case strings.HasPrefix(href, "/collages.php?id="):
		// strip collage links, they won't work cross site
		return bc.convertChildren(n)
	case strings.HasPrefix(href, "/torrents.php?id="):
		// strip torrent links, they won't work cross site
		return bc.convertChildren(n)
	case strings.HasPrefix(href, "/artist.php?id="):
		// strip artist links, they won't work cross site
		return bc.convertChildren(n)
	case strings.Contains(href, "/torrents.php?recordlabel="):
		a := href[strings.Index(href, "/torrents.php?recordlabel=")+
			len("/torrents.php?recordlabel="):]
		if t := Text(n.FirstChild); a != t {
			return fmt.Errorf(
				"recordlabel tag doesn't match text, %s != %s",
				a, t)
		}
		bc.convertChildren(n)
		return nil
	case strings.Contains(href, "/torrents.php?taglist="):
		a := href[strings.Index(href, "/torrents.php?taglist=")+
			len("/torrents.php?taglist="):]
		if t := Text(n.FirstChild); a != t {
			return fmt.Errorf(
				"taglist tag doesn't match text, %s != %s",
				a, t)
		}
		bc.convertChildren(n)
		return nil
	default:
		if AssertText(n.FirstChild, href) == nil {
			// href = anchor text
			// urls are autolinked
			// no tags required
			bc.WriteString(href)
			return nil
		}
		bc.WriteString(`[url=`)
		bc.WriteString(href)
		bc.WriteString(`]`)
		if err := bc.convertChildren(n); err != nil {
			return err
		}
		bc.WriteString("[/url]")
	}
	return nil
}

func (bc *BBCode) element(n *html.Node) error {
	switch n.DataAtom {
	case atom.Html, atom.Head, atom.Body:
		return bc.convertChildren(n)
	case atom.A:
		return bc.A(n)
	case atom.Blockquote:
		return bc.Blockquote(n)
	case atom.Br, atom.P:
		if err := bc.convertChildren(n); err != nil {
			return err
		}
		bc.WriteString("\n")
	case atom.Div:
		return bc.Div(n)
	case atom.Hr:
		return bc.Hr(n)
	case atom.Img:
		return bc.Img(n)
	case atom.Li:
		bc.WriteString(bc.lists[len(bc.lists)-1])
		return bc.convertChildren(n)
	case atom.Ol:
		bc.lists = append(bc.lists, "[#]")
		if err := bc.convertChildren(n); err != nil {
			return err
		}
		bc.lists = bc.lists[:len(bc.lists)-1]
	case atom.Pre:
		return bc.Node(n, "pre")
	case atom.Span:
		return bc.Span(n)
	case atom.Strong:
		return bc.Strong(n)
	case atom.Ul:
		bc.lists = append(bc.lists, "[*]")
		if err := bc.convertChildren(n); err != nil {
			return err
		}
		bc.lists = bc.lists[:len(bc.lists)-1]
	default:
		return fmt.Errorf("unknown element %s", n.Data)
	}
	return nil
}

func (bc *BBCode) convert(n *html.Node) error {
	switch n.Type {
	case html.ErrorNode:
		return fmt.Errorf("error node %s", n.Data)
	case html.TextNode:
		if PartOfHidden(n) || PartOfAttributedQuote(n) || PartOfLinkedQuote(n) {
			return nil
		}
		bc.WriteString(strings.ReplaceAll(n.Data, "\n", ""))
		return nil
	case html.DocumentNode:
		return bc.convertChildren(n)
	case html.ElementNode:
		return bc.element(n)
	case html.CommentNode, html.DoctypeNode:
		return nil // ignore
	default:
		return fmt.Errorf("unknown node type %d", n.Type)
	}
}

func (bc *BBCode) convertChildren(n *html.Node) error {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if err := bc.convert(c); err != nil {
			return err
		}
	}
	return nil
}

func Convert(h string) (bb string, err error) {
	bc := BBCode{}
	doc, err := html.Parse(strings.NewReader(h))
	if err != nil {
		return "", err
	}
	if err = bc.convert(doc); err != nil {
		return "", err
	}
	return bc.String(), nil
}
