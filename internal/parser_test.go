package internal

import (
	"strings"
	"testing"
)

func TestFromHTML(t *testing.T) {
	html := `
<div id="my-div"
		 class="container">
	<span>Text</span>
</div>
`
	actual, _ := FromHTML(html)
	expected := strings.TrimSpace(`
div(
	"id=my-div",
	"class=container",
)(
	span()("Text"),
)
`)
	if actual != expected {
		t.Errorf("Expected:\n%v\nFound:\n%v\n", expected, actual)
	}
}
