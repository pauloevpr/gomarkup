package mk

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

const (
	MARKER_FALSE = "__$f$__"
)

type HtmlElement struct {
	tag        string
	attributes []string
	children   []Component
}

type Safe string
type Component interface{}
type ComponentMaker func(attributes ...string) ChildrenMaker
type Children []Component
type ChildrenMaker func(children ...Component) Component

func CreateComponent(tag string) ComponentMaker {
	return func(attributes ...string) ChildrenMaker {
		el := HtmlElement{
			tag:        tag,
			attributes: attributes,
		}
		var captureChildren ChildrenMaker = func(children ...Component) Component {
			el.children = children
			return el
		}
		return captureChildren
	}
}

func CreateVoidComponent(tag string) func(attributes ...string) Component {
	return func(attributes ...string) Component {
		return HtmlElement{
			tag:        tag,
			attributes: attributes,
		}
	}
}

func ForEach[T any](items []T) func(itemMaker func(item *T, index int) Component) Children {
	return func(itemMaker func(item *T, index int) Component) Children {
		collection := Children{}
		for i := range items {
			el := itemMaker(&items[i], i)
			collection = append(collection, el)
		}
		return collection
	}
}

func Is(condition bool) string {
	if condition {
		return ""
	}
	return MARKER_FALSE
}

func When(condition bool) func(maker func() Component) Component {
	return func(maker func() Component) Component {
		if condition {
			return maker()
		}
		return HtmlElement{}
	}
}

func CreateParent() (parent func(...Component) ChildrenMaker, childrenPlaceholder ChildrenMaker) {
	var children Children
	parent = func(rootComponents ...Component) ChildrenMaker {
		return func(childComponents ...Component) Component {
			children = childComponents
			return HtmlElement{
				children: rootComponents,
			}
		}
	}
	childrenPlaceholder = func(_ ...Component) Component {
		if children != nil {
			return children
		}
		return HtmlElement{}
	}
	return parent, childrenPlaceholder
}

func Render(c Component, w io.Writer) error {
	switch node := c.(type) {
	case HtmlElement:
		if node.tag == "" {
			for i := range node.children {
				err := Render(node.children[i], w)
				if err != nil {
					return err
				}
			}
		}
		if node.tag != "" {
			_, err := w.Write([]byte("<" + node.tag))
			if err != nil {
				return err
			}
			attributes := parseAttributes(node.attributes)
			for key := range attributes {
				_, err = w.Write([]byte(
					" " + key + "=\"" + template.HTMLEscapeString(attributes[key]) + "\"",
				))
				if err != nil {
					return err
				}
			}
			_, err = w.Write([]byte(">"))
			if err != nil {
				return err
			}
			_, void := voids[node.tag]
			if !void {
				for i := range node.children {
					err = Render(node.children[i], w)
					if err != nil {
						return err
					}
				}
				_, err = w.Write([]byte(
					"</" + node.tag + ">",
				))
				if err != nil {
					return err
				}
			}
		}
	case Children:
		for i := range node {
			err := Render(node[i], w)
			if err != nil {
				return err
			}
		}
	case ChildrenMaker:
		err := Render(node(), w)
		if err != nil {
			return err
		}
	case ComponentMaker:
		err := Render(node()(), w)
		if err != nil {
			return err
		}
	case Safe:
		_, err := w.Write([]byte(node))
		if err != nil {
			return err
		}
	default:
		_, err := w.Write([]byte(
			template.HTMLEscapeString(
				fmt.Sprintf("%v", node),
			)),
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseAttributes(attributes []string) map[string]string {
	result := map[string]string{}
	for _, attr := range attributes {
		parts := strings.SplitN(attr, "=", 2)
		key := parts[0]
		value := ""
		if len(parts) > 1 {
			value = parts[1]
			_, markedFalse := strings.CutSuffix(value, MARKER_FALSE)
			if markedFalse {
				continue
			}
		} else {
			_, markedFalse := strings.CutSuffix(key, MARKER_FALSE)
			if markedFalse {
				continue
			}
		}
		current, exists := result[key]
		if exists {
			value = current + " " + value
		}
		result[key] = value
	}
	return result
}

var voids = map[string]bool{
	"!DOCTYPE": true,
	"area":     true,
	"base":     true,
	"br":       true,
	"col":      true,
	"embed":    true,
	"hr":       true,
	"img":      true,
	"input":    true,
	"link":     true,
	"meta":     true,
	"param":    true,
	"source":   true,
	"track":    true,
	"wbr":      true,
	"command":  true,
	"keygen":   true,
}
