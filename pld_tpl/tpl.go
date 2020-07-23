package pld_tpl

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
	"html/template"
	"regexp"
)

var m *minify.M

var DefaultFuncMap = make(template.FuncMap)

func GetMinify() *minify.M {
	return m
}
func init() {
	m = minify.New()
	m.AddFunc("text/css", css.Minify)
	m.Add("text/html", &html.Minifier{
		KeepConditionalComments: true,
		KeepDefaultAttrVals:     true,
		KeepDocumentTags:        true,
		KeepEndTags:             true,
		KeepQuotes:              true,
		KeepWhitespace:          false,
	})
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)

	DefaultFuncMap["Key"] = Key
	DefaultFuncMap["Html"] = Unescape
	DefaultFuncMap["Json"] = Json
	DefaultFuncMap["Type"] = Type
	DefaultFuncMap["Add"] = Add
	DefaultFuncMap["Nl2br"] = Nl2br
}
