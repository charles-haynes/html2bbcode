package html2bbcode

import (
	"fmt"
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

func GetAttr(as []html.Attribute, key string) string {
	for _, a := range as {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

type BBCode struct {
	strings.Builder
	lists []string // stack of nexted list types
}

func (bc BBCode) Node(n *html.Node, tag string) error {
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

func (bc BBCode) NodeData(n *html.Node, tag string) error {
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

func (bc BBCode) NodeVal(n *html.Node, tag, v string) error {
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

func (bc BBCode) NodeValData(n *html.Node, tag, v string) error {
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

func (bc *BBCode) A(n *html.Node) error {
	return nil
}

func (bc *BBCode) Img(n *html.Node) error {
	return nil
}

func (bc *BBCode) BlockQuote(n *html.Node) error {
	return nil
}

func (bc *BBCode) Hr(n *html.Node) error {
	return nil
}

func (bc *BBCode) Pre(n *html.Node) error {
	return nil
}

func (bc *BBCode) SpanStyle(n *html.Node, v string) error {
	for _, style := range strings.Split(v, ";") {
		ss := strings.Split(style, ":")
		if len(ss) != 2 {
			return fmt.Errorf("can't parse style %s", style)
		}
		sk, sv := ss[0], ss[1]
		switch sk {
		case "font-style":
			if string(sv) != "italic" {
				return fmt.Errorf("unknown font-style %s", sv)
			}
			return bc.Node(n, "italic")
		case "text-decoration":
			if string(sv) != "underline" {
				return fmt.Errorf("unknown text-decoration %s", sv)
			}
			return bc.Node(n, "underline")
		case "color: ":
			bc.NodeVal(n, "color", sv)
		case "display:inline-block":
			// TODO
		default:
			return fmt.Errorf("unknown span style %s", ss)
		}
	}
	return nil
}

func (bc *BBCode) Span(n *html.Node) error {
	for _, a := range n.Attr {
		switch a.Key {
		case "class":
			if len(a.Val) < 5 || string(a.Val[:4]) != "size" {
				return fmt.Errorf("unknown span class %s", a.Val)
			}
			return bc.NodeVal(n, "size", a.Val[4:])
		case "style":
			return bc.SpanStyle(n, a.Val)
		}
	}
	return nil
}

func (bc *BBCode) DivStyle(n *html.Node, v string) error {
	for _, style := range strings.Split(v, ";") {
		ss := strings.Split(style, ":")
		if len(ss) != 2 {
			return fmt.Errorf("can't parse style %s", style)
		}
		sk, sv := ss[0], ss[1]
		switch sk {
		case "text-align":
			if string(sv) == "center" {
				// TODO
				return fmt.Errorf("todo")
			}
			if string(sv) == "right" {
				// TODO
				return fmt.Errorf("todo")
			}
		default:
			return fmt.Errorf("unknown divstyle %s", ss)
		}
	}
	return nil
}

func (bc *BBCode) Div(n *html.Node) error {
	for _, a := range n.Attr {
		switch a.Key {
		case "style":
			return bc.DivStyle(n, a.Val)
		}
	}
	return nil
}

func (bc *BBCode) Strong(n *html.Node) error {
	for _, a := range n.Attr {
		switch a.Key {
		case "class":
			switch a.Val {
			case "important_text":
				return bc.Node(n, "important")
			case "quote":
				return bc.Node(n, "quote")
			}
		}
	}
	return bc.Node(n, "strong")
}

func (bc *BBCode) ParseA(n *html.Node) error {
	for _, a := range n.Attr {
		switch a.Key {
		case "href":
			switch true {
			case a.Val == n.Data:
				// href = anchor text
				// urls are autolinked
				// no tags required
				bc.WriteString(a.Val)
				return nil
			case a.Val == "javascript:void(0);":
				if GetAttr(n.Attr, "onclick") ==
					"BBCode.spoiler(this)" {
					// do nothing. Will get picked up
					// by the blockquote node
				}
			case strings.HasPrefix(a.Val, "artist.php?artistname="):
				return bc.NodeData(n, "artist")
			case strings.HasPrefix(a.Val, "user.php?action=search&search="):
				return bc.NodeData(n, "user")
			case strings.HasPrefix(a.Val, "forums.php?action=viewthread&threadid="):
				return fmt.Errorf("todo")
			case strings.HasPrefix(a.Val, "requests.php?action=view&id="):
				return fmt.Errorf("todo")
			case strings.HasPrefix(a.Val, "collages.php?id="):
				return fmt.Errorf("todo")
			case strings.HasPrefix(a.Val, "torrents.php?id="):
				return fmt.Errorf("todo")
			case strings.HasPrefix(a.Val, "torrents.php?recordlabel="):
				return fmt.Errorf("todo")
			case strings.HasPrefix(a.Val, "torrents.php?taglist="):
				return fmt.Errorf("todo")
			case strings.HasPrefix(a.Val, "rel=\"noreferrer\" target=\"_blank\" href=\"http...\""):
				return fmt.Errorf("todo")
			case strings.HasPrefix(a.Val, "artist.php?artistname="):
				return fmt.Errorf("todo")
			default:
				return bc.NodeValData(n, "url", a.Val)
			}
		case "img":
			switch true {
			case strings.HasPrefix(a.Val, `alt="..." src="..."`):
				return fmt.Errorf("todo")
			case strings.HasPrefix(a.Val, `img border="0" src="static/common/smileys/..." alt=""`):
				return fmt.Errorf("todo")
			case strings.HasPrefix(a.Val, `width="18" class="scale_image" onclick="lightbox.init(this, $(this).width());" alt="http..." src="http...`):
				return fmt.Errorf("todo")
			case strings.HasPrefix(a.Val, `class="scale_image" onclick="lightbox.init(this, $(this).width());" alt="http..." src="http..."`):
				return fmt.Errorf("todo")
			}
		}
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
		return bc.BlockQuote(n)
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
		return bc.Pre(n)
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
		bc.WriteString(n.Data)
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
