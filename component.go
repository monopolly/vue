package vue

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/monopolly/css/parser"
)

type component struct {
	suffix   string
	doc      *goquery.Document
	filename string
	name     string
	body     []byte
	template []byte
	styles   []byte
	scripts  []byte
	imports  [][]byte
	cssmap   map[string]string
}

func (a *component) parse() {

	/* html */
	t1 := bytes.Index(a.body, []byte("<template>"))
	if t1 == -1 {
		return
	}
	t2 := bytes.Index(a.body[t1+1:], []byte("</template>"))
	if t2 == -1 {
		return
	}
	a.template = a.body[t1+len("<template>") : t2+1]

	//fmt.Println(string(a.template))
	a.template = escape(a.template)
	a.template = removeHTMLComments(a.template)
	a.template = optimizeHtml(a.template)

	/* scripts */
	a.doc.Find("script").Each(func(i int, s *goquery.Selection) {
		a.parseScripts([]byte(s.Text()))
		a.scripts = removeSingleJSComments(a.scripts)
		a.scripts = removeJSComments(a.scripts)
		a.scripts = optimizeJS(a.scripts)
		a.scripts = optimizeJS(a.scripts)
	})

	/* styles */
	a.doc.Find("style").Each(func(i int, s *goquery.Selection) {

		//a.parseStyles([]byte(s.Text()))
		a.styles = []byte(s.Text())
		a.styles = optimizeCSS(a.styles)
		a.styles = optimizeCSS(a.styles)
		a.styles = a.cssByLines(a.styles)
	})

	a.htmlCSSreplace()
	a.styles = bytes.ReplaceAll(a.styles, []byte("\n"), []byte(""))
	return
}

func (a *component) parseStyles(b []byte) {
	a.styles = b
	list := make(map[string]bool)
	css, err := parser.Parse(string(b))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range css.Rules {
		for _, x := range s.Selectors {
			if list[x] {
				continue
			}
			//fmt.Println(a.name, x)
			list[x] = true
		}
	}
	return
}

func (a *component) cssByLines(b []byte) []byte {
	a.cssmap = map[string]string{}
	if a.suffix == "" {
		a.suffix = fmt.Sprintf("-%d", rand.Intn(999999))
	}

	b = bytes.ReplaceAll(b, []byte("}"), []byte("}\n"))
	b = bytes.ReplaceAll(b, []byte("}@media"), []byte("}\n@media"))

	for _, x := range bytes.Split(b, []byte("\n")) {

		i := bytes.Index(x, []byte("{"))
		if i == -1 {
			continue
		}
		// h1.title, h2.title, .nice, .me:hover, .me::after
		for _, x := range bytes.Split(x[:i], []byte(",")) {
			//h1.title .title.nice .nice .me:hover .me::after

			x = bytes.TrimSpace(x)

			//.me::after
			index := bytes.Index(x, []byte("::"))
			if index > -1 {
				class := x[:index]
				a.cssmap[string(class)] = string(class) + a.suffix
				continue
			}

			//.me:hover
			index = bytes.Index(x, []byte(":"))
			if index > -1 {
				class := x[:index]
				a.cssmap[string(class)] = string(class) + a.suffix
				continue
			}
			//h1.title
			if bytes.Index(x, []byte(".")) > 0 {
				classes := bytes.Split(x, []byte("."))
				//h1 title
				if len(classes) < 2 {
					continue
				}
				classes = classes[0:]
				//title
				for _, class := range classes {
					a.cssmap["."+string(class)] = "." + string(class) + a.suffix
				}
				continue
			}

			//.title.nice
			if bytes.Index(x, []byte(".")) > -1 {
				for _, class := range bytes.Split(x, []byte(".")) {
					class = bytes.TrimSpace(class)
					a.cssmap["."+string(class)] = "." + string(class) + a.suffix
				}
				continue
			}

		}
	}

	for original, replace := range a.cssmap {
		b = bytes.ReplaceAll(b, []byte(original+" "), []byte(replace+" "))
		b = bytes.ReplaceAll(b, []byte(original+"{"), []byte(replace+"{"))
		b = bytes.ReplaceAll(b, []byte(original+":"), []byte(replace+":"))
		b = bytes.ReplaceAll(b, []byte(original+"::"), []byte(replace+"::"))
	}
	return b
}

func (a *component) htmlCSSreplace() {
	a.template = replaceClasses(a.template, func(oldclass string) (newclass string) {
		if a.cssmap["."+oldclass] != "" {
			newclass = strings.TrimPrefix(a.cssmap["."+oldclass], ".")
			return
		}
		return
	})
	return
}

func (a *component) parseScripts(b []byte) {
	first := bytes.Index(b, []byte("{"))
	if first == -1 {
		return
	}
	first++
	last := bytes.LastIndex(b, []byte("}"))
	if last == -1 {
		return
	}

	a.scripts = b[first:last]
	a.name = noext(a.filename)

}

func (a *component) export() []byte {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf(`Vue.component('%s', {`, a.name))
	b.WriteString("template: `")
	b.Write(a.template)
	b.WriteString("`,")
	b.Write(a.scripts)
	b.WriteString(`})`)
	return optimizeJS(b.Bytes())
}
