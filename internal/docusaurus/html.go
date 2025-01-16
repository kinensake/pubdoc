package docusaurus

import (
	"io"
	"regexp"
	"slices"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

var voidTags = []string{
	"area", "base", "br", "col", "embed", "hr", "img",
	"input", "link", "meta", "param", "source", "track", "wbr",
}

func sanitizeHTML(s io.Reader) string {
	p := bluemonday.NewPolicy()

	p.AllowStandardURLs()

	p.AllowElements("article", "aside")
	p.AllowElements("figure")
	p.AllowElements("section")
	p.AllowElements("summary")
	p.AllowElements("h1", "h2", "h3", "h4", "h5", "h6")
	p.AllowElements("hgroup")
	p.AllowElements("br", "div", "hr", "p", "span", "wbr")
	p.AllowElements("b", "i", "pre", "small", "strike", "tt", "u")
	p.AllowElements("rp", "rt", "ruby")
	p.AllowElements("abbr", "acronym", "cite", "code", "dfn", "em",
		"figcaption", "mark", "s", "samp", "strong", "sub", "sup", "var")

	p.AllowAttrs("cite").OnElements("blockquote")
	p.AllowAttrs("href").OnElements("a")

	p.AllowLists()
	p.AllowTables()
	p.AllowImages()

	result := p.SanitizeReader(s).String()
	result = removeInvalidTags(result)

	return result
}

func removeInvalidTags(s string) string {
	r := regexp.MustCompile(`<\w+[^>]*\/>`)
	tags := r.FindAllString(s, -1)

	for _, v := range tags {
		end := strings.IndexAny(v, " /")

		if end > 0 && !slices.Contains(voidTags, v[1:end]) {
			s = strings.ReplaceAll(s, v, "")
		}
	}

	return s
}
