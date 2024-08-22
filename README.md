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

import "github.com/pauloevpr/aussiedollar/internal/gomarkup"

var fragment = gomarkup.CreateComponent("")
var a = gomarkup.CreateComponent("a")
var abbr = gomarkup.CreateComponent("abbr")
var acronym = gomarkup.CreateComponent("acronym")
var address = gomarkup.CreateComponent("address")
var applet = gomarkup.CreateComponent("applet")
var area = gomarkup.CreateVoidComponent("area")
var article = gomarkup.CreateComponent("article")
var aside = gomarkup.CreateComponent("aside")
var audio = gomarkup.CreateComponent("audio")
var b = gomarkup.CreateComponent("b")
var base = gomarkup.CreateVoidComponent("base")
var basefont = gomarkup.CreateComponent("basefont")
var bdi = gomarkup.CreateComponent("bdi")
var bdo = gomarkup.CreateComponent("bdo")
var big = gomarkup.CreateComponent("big")
var blockquote = gomarkup.CreateComponent("blockquote")
var body = gomarkup.CreateComponent("body")
var br = gomarkup.CreateVoidComponent("br")
var button = gomarkup.CreateComponent("button")
var canvas = gomarkup.CreateComponent("canvas")
var caption = gomarkup.CreateComponent("caption")
var center = gomarkup.CreateComponent("center")
var cite = gomarkup.CreateComponent("cite")
var code = gomarkup.CreateComponent("code")
var col = gomarkup.CreateVoidComponent("col")
var colgroup = gomarkup.CreateComponent("colgroup")
var data = gomarkup.CreateComponent("data")
var datalist = gomarkup.CreateComponent("datalist")
var dd = gomarkup.CreateComponent("dd")
var del = gomarkup.CreateComponent("del")
var details = gomarkup.CreateComponent("details")
var dfn = gomarkup.CreateComponent("dfn")
var dialog = gomarkup.CreateComponent("dialog")
var dir = gomarkup.CreateComponent("dir")
var div = gomarkup.CreateComponent("div")
var dl = gomarkup.CreateComponent("dl")
var dt = gomarkup.CreateComponent("dt")
var doctype = gomarkup.CreateComponent("!DOCTYPE")
var em = gomarkup.CreateComponent("em")
var embed = gomarkup.CreateVoidComponent("embed")
var fieldset = gomarkup.CreateComponent("fieldset")
var figcaption = gomarkup.CreateComponent("figcaption")
var figure = gomarkup.CreateComponent("figure")
var font = gomarkup.CreateComponent("font")
var footer = gomarkup.CreateComponent("footer")
var form = gomarkup.CreateComponent("form")
var frame = gomarkup.CreateComponent("frame")
var frameset = gomarkup.CreateComponent("frameset")
var h1 = gomarkup.CreateComponent("h1")
var h2 = gomarkup.CreateComponent("h2")
var h3 = gomarkup.CreateComponent("h3")
var h4 = gomarkup.CreateComponent("h4")
var h5 = gomarkup.CreateComponent("h5")
var h6 = gomarkup.CreateComponent("h6")
var head = gomarkup.CreateComponent("head")
var header = gomarkup.CreateComponent("header")
var hgroup = gomarkup.CreateComponent("hgroup")
var hr = gomarkup.CreateVoidComponent("hr")
var html = gomarkup.CreateComponent("html")
var i = gomarkup.CreateComponent("i")
var iframe = gomarkup.CreateComponent("iframe")
var img = gomarkup.CreateVoidComponent("img")
var input = gomarkup.CreateVoidComponent("input")
var ins = gomarkup.CreateComponent("ins")
var kbd = gomarkup.CreateComponent("kbd")
var label = gomarkup.CreateComponent("label")
var legend = gomarkup.CreateComponent("legend")
var li = gomarkup.CreateComponent("li")
var link = gomarkup.CreateVoidComponent("link")
var main = gomarkup.CreateComponent("main")
var map_ = gomarkup.CreateComponent("map")
var mark = gomarkup.CreateComponent("mark")
var menu = gomarkup.CreateComponent("menu")
var meta = gomarkup.CreateVoidComponent("meta")
var meter = gomarkup.CreateComponent("meter")
var nav = gomarkup.CreateComponent("nav")
var noframes = gomarkup.CreateComponent("noframes")
var noscript = gomarkup.CreateComponent("noscript")
var object = gomarkup.CreateComponent("object")
var ol = gomarkup.CreateComponent("ol")
var optgroup = gomarkup.CreateComponent("optgroup")
var option = gomarkup.CreateComponent("option")
var output = gomarkup.CreateComponent("output")
var p = gomarkup.CreateComponent("p")
var param = gomarkup.CreateVoidComponent("param")
var picture = gomarkup.CreateComponent("picture")
var pre = gomarkup.CreateComponent("pre")
var progress = gomarkup.CreateComponent("progress")
var q = gomarkup.CreateComponent("q")
var rp = gomarkup.CreateComponent("rp")
var rt = gomarkup.CreateComponent("rt")
var ruby = gomarkup.CreateComponent("ruby")
var s = gomarkup.CreateComponent("s")
var samp = gomarkup.CreateComponent("samp")
var script = gomarkup.CreateComponent("script")
var search = gomarkup.CreateComponent("search")
var section = gomarkup.CreateComponent("section")
var select_ = gomarkup.CreateComponent("select")
var small = gomarkup.CreateComponent("small")
var source = gomarkup.CreateVoidComponent("source")
var span = gomarkup.CreateComponent("span")
var strike = gomarkup.CreateComponent("strike")
var strong = gomarkup.CreateComponent("strong")
var style = gomarkup.CreateComponent("style")
var sub = gomarkup.CreateComponent("sub")
var summary = gomarkup.CreateComponent("summary")
var sup = gomarkup.CreateComponent("sup")
var svg = gomarkup.CreateComponent("svg")
var table = gomarkup.CreateComponent("table")
var tbody = gomarkup.CreateComponent("tbody")
var td = gomarkup.CreateComponent("td")
var template = gomarkup.CreateComponent("template")
var textarea = gomarkup.CreateComponent("textarea")
var tfoot = gomarkup.CreateComponent("tfoot")
var th = gomarkup.CreateComponent("th")
var thead = gomarkup.CreateComponent("thead")
var time_ = gomarkup.CreateComponent("time")
var title = gomarkup.CreateComponent("title")
var tr = gomarkup.CreateComponent("tr")
var track = gomarkup.CreateVoidComponent("track")
var tt = gomarkup.CreateComponent("tt")
var u = gomarkup.CreateComponent("u")
var ul = gomarkup.CreateComponent("ul")
var var_ = gomarkup.CreateComponent("var")
var video = gomarkup.CreateComponent("video")
var wbr = gomarkup.CreateVoidComponent("wbr")
var path = gomarkup.CreateComponent("path")

func foreach[T any](items []T) func(itemMaker func(item *T, index int) gomarkup.Component) gomarkup.Children {
	return gomarkup.ForEach(items)
}

func is(condition bool) string {
	return gomarkup.Is(condition)
}

func when(condition bool) func(maker func() gomarkup.Component) gomarkup.Component {
	return gomarkup.When(condition)
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
