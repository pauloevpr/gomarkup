package internal

import (
	"go/format"
	"strings"

	"golang.org/x/net/html"
)

func FromHTML(content string) (string, error) {
	roots, err := html.ParseFragment(strings.NewReader(content), &html.Node{
		Data: "gomarkup",
		Type: html.ElementNode,
	})
	if err != nil {
		return "", err
	}

	b := &strings.Builder{}
	for i := range roots {
		process(roots[i], b)
	}
	code, err := format.Source([]byte(b.String()))
	return strings.TrimSpace(string(code)), nil
}

func process(node *html.Node, b *strings.Builder) {

	if node.Type == html.ElementNode {
		multiple := len(node.Attr) > 1
		b.WriteString(node.Data + "(")
		if multiple {
			b.WriteString("\n")
		}
		for _, attr := range node.Attr {
			b.WriteString(`"`)
			if attr.Namespace != "" {
				b.WriteString(attr.Namespace + ":")
			}
			b.WriteString(attr.Key + `=` + attr.Val + `"`)
			if multiple {
				b.WriteString(",\n")
			}
		}
		b.WriteString(")")
	}

	if node.Type == html.TextNode && strings.TrimSpace(node.Data) != "" {
		multiline := strings.Index(node.Data, "\n") != -1
		if multiline {
			b.WriteString("`" + node.Data + "`")
		} else {
			b.WriteString(`"` + node.Data + `"`)
		}
	}

	children := []*html.Node{}
	next := node.FirstChild
	for next != nil {
		if !(next.Type == html.TextNode && strings.TrimSpace(next.Data) == "") {
			children = append(children, next)
		}
		next = next.NextSibling
	}

	if len(children) > 0 {
		singleTextNode := len(children) == 1 && children[0].Type == html.TextNode
		if singleTextNode {
			b.WriteString(`("` + strings.TrimSpace(children[0].Data) + `")`)
		} else {
			b.WriteString("(\n")
			for _, child := range children {
				process(child, b)
				b.WriteString(",\n")
			}
			b.WriteString(")")
		}
	}
}
