package vue

import (
	"bytes"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

/* экранируем `` для вью */
func escape(b []byte) []byte {
	b = bytes.ReplaceAll(b, []byte("`"), []byte("\\`"))
	b = bytes.ReplaceAll(b, []byte("\\\\"), []byte("\\"))
	return b
}

func noext(f string) string {
	return strings.TrimSuffix(strings.TrimSpace(filepath.Base(f)), filepath.Ext(f))
}

func removeJSComments(html []byte) []byte {
	var lines [][]byte
	var comment bool

	for _, x := range bytes.Split(html, []byte("\n")) {
		x = bytes.TrimSpace(x)
		if x == nil {
			continue
		}
		if bytes.HasPrefix(x, []byte("//")) {
			continue
		}
		/* <!-- comment --> */
		if bytes.HasPrefix(x, []byte("/*")) && bytes.HasSuffix(x, []byte("*/")) {
			continue
		}

		/* <!-- comment start */
		if bytes.HasPrefix(x, []byte("/*")) {
			comment = true
			continue
		}

		/* comment end --> */
		if bytes.HasSuffix(x, []byte("*/")) {
			comment = false
			continue
		}

		if comment {
			continue
		}

		lines = append(lines, x)
	}

	return bytes.Join(lines, []byte("\n"))
}

func removeSingleJSComments(html []byte) []byte {
	var lines [][]byte
	var script bool

	for _, x := range bytes.Split(html, []byte("\n")) {
		x = bytes.TrimSpace(x)
		if x == nil {
			continue
		}
		/* <!-- comment --> */
		if bytes.HasPrefix(x, []byte("<script>")) && bytes.HasSuffix(x, []byte("</script>")) {
			in := bytes.LastIndex(x, []byte("//"))
			if in > -1 && x[in-1] != ':' {
				x = x[:in]
				lines = append(lines, x)
				continue
			}
		}

		/* <!-- comment start */
		if bytes.HasPrefix(x, []byte("<script>")) {
			script = true
		}

		/* comment end --> */
		if bytes.HasSuffix(x, []byte("</script>")) {
			script = false
			in := bytes.LastIndex(x, []byte("//"))
			if in > -1 && x[in-1] != ':' {
				x = x[:in]
			}
		}

		if script {
			in := bytes.LastIndex(x, []byte("//"))
			if in > -1 && x[in-1] != ':' {
				x = x[:in]
			}
		}

		lines = append(lines, x)
	}

	return bytes.Join(lines, []byte("\n"))
}

func removeHTMLComments(html []byte) []byte {
	var lines [][]byte
	var comment bool

	for _, x := range bytes.Split(html, []byte("\n")) {
		x = bytes.TrimSpace(x)
		if x == nil {
			continue
		}
		/* <!-- comment --> */
		if bytes.HasPrefix(x, []byte("<!--")) && bytes.HasSuffix(x, []byte("-->")) {
			continue
		}

		/* <!-- comment start */
		if bytes.HasPrefix(x, []byte("<!--")) {
			comment = true
			continue
		}

		/* comment end --> */
		if bytes.HasSuffix(x, []byte("-->")) {
			comment = false
			continue
		}

		if comment {
			continue
		}

		lines = append(lines, x)
	}

	return bytes.Join(lines, []byte("\n"))
}

func optimizeHtml(b []byte) []byte {
	b = bytes.ReplaceAll(b, []byte(">\n<"), []byte("><"))
	b = bytes.ReplaceAll(b, []byte(">\n"), []byte(">"))
	b = bytes.ReplaceAll(b, []byte("\n<"), []byte("<"))
	return b
}

func optimizeJS(b []byte) []byte {
	b = bytes.ReplaceAll(b, []byte("],\n["), []byte("],["))
	b = bytes.ReplaceAll(b, []byte("[\n["), []byte("[["))
	b = bytes.ReplaceAll(b, []byte("},\n{"), []byte("},{"))
	b = bytes.ReplaceAll(b, []byte("{\n"), []byte("{"))
	b = bytes.ReplaceAll(b, []byte("}\n},"), []byte("}},"))
	b = bytes.ReplaceAll(b, []byte("',\n"), []byte("', "))
	b = bytes.ReplaceAll(b, []byte("\n]},"), []byte("]},"))
	b = bytes.ReplaceAll(b, []byte("\n],"), []byte("],"))
	b = bytes.ReplaceAll(b, []byte(",\n},"), []byte("},"))
	b = bytes.ReplaceAll(b, []byte(",]"), []byte("]"))
	b = bytes.ReplaceAll(b, []byte(",}"), []byte("}"))
	b = bytes.ReplaceAll(b, []byte("\n,}}"), []byte("}}"))
	b = bytes.ReplaceAll(b, []byte("\n}}"), []byte("}}"))
	b = bytes.ReplaceAll(b, []byte("\n}"), []byte("}"))
	b = bytes.ReplaceAll(b, []byte(", }"), []byte("}"))
	b = bytes.ReplaceAll(b, []byte(",\n"), []byte(", "))
	b = bytes.ReplaceAll(b, []byte(")\n"), []byte(");"))
	b = bytes.ReplaceAll(b, []byte(");\n"), []byte(");"))
	b = bytes.ReplaceAll(b, []byte(":\t"), []byte(":"))
	b = bytes.ReplaceAll(b, []byte(": "), []byte(":"))
	b = bytes.ReplaceAll(b, []byte(", "), []byte(","))
	b = bytes.ReplaceAll(b, []byte("},}"), []byte("}}"))
	b = bytes.ReplaceAll(b, []byte(", {"), []byte(",{"))
	b = bytes.ReplaceAll(b, []byte("',\n"), []byte("',"))
	b = bytes.ReplaceAll(b, []byte("return\n"), []byte("return;"))
	b = bytes.ReplaceAll(b, []byte("[\n"), []byte("["))
	b = bytes.ReplaceAll(b, []byte(";}"), []byte("}"))
	b = bytes.ReplaceAll(b, []byte(") {"), []byte("){"))
	b = bytes.ReplaceAll(b, []byte(" ("), []byte("("))
	b = bytes.ReplaceAll(b, []byte(" ["), []byte("["))
	b = bytes.ReplaceAll(b, []byte(" {"), []byte("{"))
	b = bytes.ReplaceAll(b, []byte("} "), []byte("}"))
	b = bytes.ReplaceAll(b, []byte(" ="), []byte("="))
	b = bytes.ReplaceAll(b, []byte("= "), []byte("="))
	b = bytes.ReplaceAll(b, []byte(" ||"), []byte("||"))
	b = bytes.ReplaceAll(b, []byte("|| "), []byte("||"))
	b = bytes.ReplaceAll(b, []byte(" &&"), []byte("&&"))
	b = bytes.ReplaceAll(b, []byte("&& "), []byte("&&"))
	b = bytes.ReplaceAll(b, []byte(":\n"), []byte(":"))
	b = bytes.ReplaceAll(b, []byte("\n:"), []byte(":"))
	b = bytes.ReplaceAll(b, []byte(";\n"), []byte(";"))
	b = bytes.ReplaceAll(b, []byte(" *"), []byte("*"))
	b = bytes.ReplaceAll(b, []byte("* "), []byte("*"))
	b = bytes.ReplaceAll(b, []byte(" /"), []byte("/"))
	b = bytes.ReplaceAll(b, []byte("/ "), []byte("/"))

	return b
}

var space = regexp.MustCompile(`\s+`)

func optimizeCSS(b []byte) []byte {
	b = deleteTag(b, "/*", "*/")
	b = bytes.ReplaceAll(b, []byte("  "), []byte(" "))
	b = bytes.ReplaceAll(b, []byte("\n"), []byte(""))
	b = bytes.ReplaceAll(b, []byte("\n}"), []byte("}"))
	b = bytes.ReplaceAll(b, []byte("{\n"), []byte("{"))
	b = bytes.ReplaceAll(b, []byte(";\n"), []byte(";"))
	b = bytes.ReplaceAll(b, []byte(":  "), []byte(":"))
	b = bytes.ReplaceAll(b, []byte(": "), []byte(":"))
	b = bytes.ReplaceAll(b, []byte("{ "), []byte("{"))
	b = bytes.ReplaceAll(b, []byte(" {"), []byte("{"))
	b = bytes.ReplaceAll(b, []byte(";  "), []byte(";"))
	b = bytes.ReplaceAll(b, []byte("; "), []byte(";"))
	b = bytes.ReplaceAll(b, []byte("  ;"), []byte(";"))
	b = bytes.ReplaceAll(b, []byte(" ;"), []byte(";"))
	b = bytes.ReplaceAll(b, []byte(", "), []byte(","))
	b = []byte(space.ReplaceAllString(string(b), " "))
	b = bytes.ReplaceAll(b, []byte(";}"), []byte("}"))
	return b
}

func deleteTag(html []byte, tagStart, tagEnd string) []byte {
	const spaceRune = byte(32)
	lenStart := len(tagStart)
	lenEnd := len(tagEnd)
	var part, end []byte
	var begin bool
	var start, ends int
	for pos, x := range html {
		switch begin {
		case false:
			part = append(part, x)
			if len(part) == lenStart+1 {
				part = part[1:]
			}
			if string(part) == tagStart {
				start = pos - lenStart + 1
				begin = true
			}
			continue
		case true:
			end = append(end, x)
			if len(end) == lenEnd+1 {
				end = end[1:]
			}
			if string(end) == tagEnd {
				ends = pos
				begin = false
				for x := start; x < ends+1; x++ {
					html[x] = spaceRune
				}
				continue
			}
		}
	}
	return html
}

func replaceClasses(b []byte, h func(class string) (replace string)) []byte {

	t1 := []byte(`class="`)
	t2 := []byte(`"`)
	//fmt.Println(string(b))
	var begin bool
	var x1, x2, last int
	var res bytes.Buffer
	for {
		switch begin {
		case true:
			x2 = bytes.Index(b[last:], t2)

			if x2 > -1 {
				x2 = x2 + last
				//fmt.Println("x2", x2)
				//fmt.Println("class", string(b[x1:x2]))
				//stop
				classes := []string{}

				//vue format
				if bytes.Index(b[x1:x2], []byte("{")) > -1 ||
					bytes.Index(b[x1:x2], []byte("[")) > -1 ||
					bytes.Index(b[x1:x2], []byte("]")) > -1 ||
					bytes.Index(b[x1:x2], []byte("}")) > -1 ||
					bytes.Index(b[x1:x2], []byte("$")) > -1 ||
					bytes.Index(b[x1:x2], []byte(":")) > -1 {

					//{pay-button-disable:!enable, nice: true}
					if bytes.Index(b[x1:x2], []byte("{")) > -1 && bytes.Index(b[x1:x2], []byte("}")) > -1 {
						cav := bytes.ReplaceAll(b[x1:x2], []byte("{"), []byte(""))
						cav = bytes.ReplaceAll(cav, []byte("}"), []byte(""))
						//pay-button-disable:!enable, nice: true

						for _, classd := range bytes.Split(cav, []byte(",")) {
							classd = bytes.TrimSpace(classd)

							//pay-button-disable:!enable
							part := bytes.Split(classd, []byte(":"))
							if len(part) != 2 {
								continue
							}

							//'pay-button-disable'
							part[0] = bytes.TrimSpace(part[0])
							part[1] = bytes.TrimSpace(part[1])
							var orn bool

							if bytes.Index(part[0], []byte("'")) > -1 {
								orn = true
								part[0] = bytes.ReplaceAll(part[0], []byte("'"), []byte(""))
								//pay-button-disable
							}

							newclass := h(string(part[0]))
							if newclass == "" {
								newclass = string(part[0])
							}
							if orn {
								classes = append(classes, fmt.Sprintf(`'%s':%s`, newclass, string(part[1])))
							} else {
								classes = append(classes, fmt.Sprintf(`%s:%s`, newclass, string(part[1])))
							}
						}
						res.WriteString(fmt.Sprintf(`{%s}`, strings.Join(classes, ",")))
						classes = nil
					}

					//{pay-button-disable:!enable}

					/* fmt.Println("Vue Css vars")
					res.Write(b[x1:x2])
					*/
					begin = false
					last = x2
					continue
				}

				//clear css
				for _, x := range bytes.Fields(b[x1:x2]) {
					x = bytes.TrimSpace(x)
					newclass := h(string(x))
					if newclass == "" {
						classes = append(classes, string(x))
						continue
					}
					classes = append(classes, newclass)
				}
				res.WriteString(strings.Join(classes, " "))
				//res.WriteString(`"`)
				last = x2
				begin = false
				continue
			}
			res.Write(b[last:])
			break
		default:
			//fmt.Println("last", last)
			x1 = bytes.Index(b[last:], t1)
			if x1 > -1 {
				x1 = x1 + last + len(t1)
				res.Write(b[last:x1])
				last = x1
				begin = true
				continue
			}
			res.Write(b[last:])
			break
		}
		return res.Bytes()
	}

}

//достает все стили и удаляет их
func extractCSS(body []byte) (css, newbody []byte) {
	var b bytes.Buffer
	b.Write(body)
	doc, _ := goquery.NewDocumentFromReader(&b)
	var lines [][]byte
	doc.Find("style").Each(func(i int, s *goquery.Selection) {
		lines = append(lines, []byte(s.Text()))
	})

	doc.Find("*").Each(func(i int, s *goquery.Selection) {})

	return bytes.Join(lines, []byte("\n")), deleteTag(body, "<style>", "</style>")
}
