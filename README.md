# gomarkup

A simple and yet expressive HTML templating library for go. No special syntax. No build step. No editor plugins. Just straight up go.

# Installation

## Alias File

To simplify writing HTML elements using `gomarkup`, you can create an alias file in your project. This file will allow you to use lowercase functions for HTML tags, making your code cleaner and more readable.

<details>
    <summary>
Here’s the full example alias file:
    </summary>


```go
package my_package

import "github.com/pauloevpr/gomarkup/pkg/mk"

var fragment = mk.CreateComponent("")
var a = mk.CreateComponent("a")
var abbr = mk.CreateComponent("abbr")
var acronym = mk.CreateComponent("acronym")
var address = mk.CreateComponent("address")
var applet = mk.CreateComponent("applet")
var area = mk.CreateVoidComponent("area")
var article = mk.CreateComponent("article")
var aside = mk.CreateComponent("aside")
var audio = mk.CreateComponent("audio")
var b = mk.CreateComponent("b")
var base = mk.CreateVoidComponent("base")
var basefont = mk.CreateComponent("basefont")
var bdi = mk.CreateComponent("bdi")
var bdo = mk.CreateComponent("bdo")
var big = mk.CreateComponent("big")
var blockquote = mk.CreateComponent("blockquote")
var body = mk.CreateComponent("body")
var br = mk.CreateVoidComponent("br")
var button = mk.CreateComponent("button")
var canvas = mk.CreateComponent("canvas")
var caption = mk.CreateComponent("caption")
var center = mk.CreateComponent("center")
var cite = mk.CreateComponent("cite")
var code = mk.CreateComponent("code")
var col = mk.CreateVoidComponent("col")
var colgroup = mk.CreateComponent("colgroup")
var data = mk.CreateComponent("data")
var datalist = mk.CreateComponent("datalist")
var dd = mk.CreateComponent("dd")
var del = mk.CreateComponent("del")
var details = mk.CreateComponent("details")
var dfn = mk.CreateComponent("dfn")
var dialog = mk.CreateComponent("dialog")
var dir = mk.CreateComponent("dir")
var div = mk.CreateComponent("div")
var dl = mk.CreateComponent("dl")
var dt = mk.CreateComponent("dt")
var doctype = mk.CreateComponent("!DOCTYPE")
var em = mk.CreateComponent("em")
var embed = mk.CreateVoidComponent("embed")
var fieldset = mk.CreateComponent("fieldset")
var figcaption = mk.CreateComponent("figcaption")
var figure = mk.CreateComponent("figure")
var font = mk.CreateComponent("font")
var footer = mk.CreateComponent("footer")
var form = mk.CreateComponent("form")
var frame = mk.CreateComponent("frame")
var frameset = mk.CreateComponent("frameset")
var h1 = mk.CreateComponent("h1")
var h2 = mk.CreateComponent("h2")
var h3 = mk.CreateComponent("h3")
var h4 = mk.CreateComponent("h4")
var h5 = mk.CreateComponent("h5")
var h6 = mk.CreateComponent("h6")
var head = mk.CreateComponent("head")
var header = mk.CreateComponent("header")
var hgroup = mk.CreateComponent("hgroup")
var hr = mk.CreateVoidComponent("hr")
var html = mk.CreateComponent("html")
var i = mk.CreateComponent("i")
var iframe = mk.CreateComponent("iframe")
var img = mk.CreateVoidComponent("img")
var input = mk.CreateVoidComponent("input")
var ins = mk.CreateComponent("ins")
var kbd = mk.CreateComponent("kbd")
var label = mk.CreateComponent("label")
var legend = mk.CreateComponent("legend")
var li = mk.CreateComponent("li")
var link = mk.CreateVoidComponent("link")
var main = mk.CreateComponent("main")
var map_ = mk.CreateComponent("map")
var mark = mk.CreateComponent("mark")
var menu = mk.CreateComponent("menu")
var meta = mk.CreateVoidComponent("meta")
var meter = mk.CreateComponent("meter")
var nav = mk.CreateComponent("nav")
var noframes = mk.CreateComponent("noframes")
var noscript = mk.CreateComponent("noscript")
var object = mk.CreateComponent("object")
var ol = mk.CreateComponent("ol")
var optgroup = mk.CreateComponent("optgroup")
var option = mk.CreateComponent("option")
var output = mk.CreateComponent("output")
var p = mk.CreateComponent("p")
var param = mk.CreateVoidComponent("param")
var picture = mk.CreateComponent("picture")
var pre = mk.CreateComponent("pre")
var progress = mk.CreateComponent("progress")
var q = mk.CreateComponent("q")
var rp = mk.CreateComponent("rp")
var rt = mk.CreateComponent("rt")
var ruby = mk.CreateComponent("ruby")
var s = mk.CreateComponent("s")
var samp = mk.CreateComponent("samp")
var script = mk.CreateComponent("script")
var search = mk.CreateComponent("search")
var section = mk.CreateComponent("section")
var select_ = mk.CreateComponent("select")
var small = mk.CreateComponent("small")
var source = mk.CreateVoidComponent("source")
var span = mk.CreateComponent("span")
var strike = mk.CreateComponent("strike")
var strong = mk.CreateComponent("strong")
var style = mk.CreateComponent("style")
var sub = mk.CreateComponent("sub")
var summary = mk.CreateComponent("summary")
var sup = mk.CreateComponent("sup")
var svg = mk.CreateComponent("svg")
var table = mk.CreateComponent("table")
var tbody = mk.CreateComponent("tbody")
var td = mk.CreateComponent("td")
var template = mk.CreateComponent("template")
var textarea = mk.CreateComponent("textarea")
var tfoot = mk.CreateComponent("tfoot")
var th = mk.CreateComponent("th")
var thead = mk.CreateComponent("thead")
var time_ = mk.CreateComponent("time")
var title = mk.CreateComponent("title")
var tr = mk.CreateComponent("tr")
var track = mk.CreateVoidComponent("track")
var tt = mk.CreateComponent("tt")
var u = mk.CreateComponent("u")
var ul = mk.CreateComponent("ul")
var var_ = mk.CreateComponent("var")
var video = mk.CreateComponent("video")
var wbr = mk.CreateVoidComponent("wbr")
var path = mk.CreateComponent("path")

func foreach[T any](items []T) func(itemMaker func(item *T, index int) mk.Component) mk.Children {
	return mk.ForEach(items)
}

func is(condition bool) string {
	return mk.Is(condition)
}

func when(condition bool) func(maker func() mk.Component) mk.Component {
	return mk.When(condition)
}
```
</details>

Include this alias file in your project to start using lowercase HTML tags.
# Basics
## Elements and Attributes

In gomarkup, HTML elements are written using simple function calls. Here is a simple component that renders a form. 

```go
func MyForm() gomarkup.Component {
    return form(
        "method=post",
        "action=/users",
    )(
        label("for=name")("Name"),
        input(
            "id=name",
            "name=name",
            "type=text",
            "class=rounded border shadow h-12",
            "placeholder=First Name",
            "required",
        ),
        button(
            "class=font-medium rounded-full bg-pink-600 text-white",
            "type=submit",
        ),
    )
}
```

Basic rules:

- A typical HTML element is written using the format `element(attributes)(children)`
- If you don't need attributes, use: `element()(children)`
- If you need attributes but no children, use: `element(attributes)`
- For no attributes and no children, use either: `element`, `element()`, `element()()`
- For void elements like `input`, use: `element(attributes)`

### Attributes

Attributes in gomarkup are written as simple strings. Condider the example below:

```go
func MyInput() gomarkup.Component {
    label := "My custom input"
    return input(
        "id=age",
        "name=age",
        "type=number",
        "class=rounded border h-12",
        "required",
        "aria-label="+label,
    )
}
```

Basic rules:

- A typical attribute uses the format `"name=value"` which would be equivalent to `name="value"` in HTML:
    - The `name` is exactly the name you would use when writing HTML
    - The `value` is exactly the string value you would use when writing HTML
- For attributes without values, use just the attribute name: `"required"`
- For dynamic values, use simple string concatenation: `"name="+value`
- For dynamic non-string values, do the `fmt` thing outside the markup and use simple string concatenation: `"name="+formattedValue`



### Conditional attributes
### Conditional styling
## List Rendering
## Conditional Rendering
## Components
## Props
## Parent/child
## Special Elements
### <script>
### <style>
