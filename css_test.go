package vue

//testing

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/aymerick/douceur/parser"
	"github.com/stretchr/testify/assert"

	"github.com/monopolly/file"
)

func TestAccount_Marshal(ggggg *testing.T) {
	function, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(function).Name()
	fn = fn[strings.LastIndex(fn, ".Test")+5:]
	fn = strings.Join(strings.Split(fn, "_"), ": ")
	fmt.Printf("\033[1;32m%s\033[0m\n", fn)

	a := assert.New(ggggg)
	_ = a

	var b bytes.Buffer
	b.Write(file.OpenE("test.vue"))
	doc, err := goquery.NewDocumentFromReader(&b)
	a.Nil(err)

	/* styles */
	id := fmt.Sprint(time.Now().Unix())
	var css bytes.Buffer

	doc.Find("style").Each(func(i int, s *goquery.Selection) {
		p, _ := parser.Parse(s.Text())
		for _, x := range p.Rules {
			switch x.Name != "" {
			case true:
				switch x.Name {
				case "@keyframes":
					css.WriteString(fmt.Sprintf("%s %s-%s ", x.Name, x.Prelude, id))
					css.WriteString("{")
					for _, x := range x.Rules {
						css.WriteString(x.String())
					}
					css.WriteString("}")
				case "@media":
					css.WriteString(fmt.Sprintf("%s %s", x.Name, x.Prelude))
					css.WriteString("{")
					for _, x := range x.Rules {
						parseSimpleSelectors(x.String(), "")
						css.WriteString(x.String())
					}
					css.WriteString("}")
				default:

				}
			default:

			}

			fmt.Println("Name", x.Name)
			//fmt.Println("Kind", x.Kind.String())
			fmt.Println("Prelude", x.Prelude)
			fmt.Println("Rules", x.Rules)
			for _, ss := range x.Selectors {
				hover := strings.Index(ss, ":")
				if hover > -1 {
					fmt.Println("selector", ss[0:hover])
				}

				mt := strings.Index(ss, " ")
				if mt > -1 {
					fmt.Println("selector", ss[0:mt])
				}

			}

			fmt.Println()
			//fmt.Println(fmt.Sprint("%s %s ", x.Name, x.Prelude,) )
			fmt.Println()

		}

	})

	file.Save("test.css", css.Bytes())

}

/* Rules [body {
   background-color: lightblue;
 }] */
func parseSimpleSelectors(s string, id string) []byte {
	var b bytes.Buffer
	p, _ := parser.Parse(s)
	for _, x := range p.Rules {

		if !strings.HasPrefix(x.Prelude, ".") {
			if strings.Index(x.Prelude, ":") > -1 || strings.Index(x.Prelude, ".") > -1 {
				fmt.Println("has norm selector")
			}
		}

		b.WriteString(fmt.Sprintf("%s-%s", x.Prelude, id))
		for _, c := range x.Declarations {
			b.WriteString(c.String())
		}
		fmt.Println("Simple selector")
		//fmt.Println("Name", x.Name)
		fmt.Println("Prelude", x.Prelude)
		fmt.Println("Rules", x.Rules)
		fmt.Println("CSS", x.Declarations)
		fmt.Println()
	}

	return b.Bytes()
}

func parseSimpleOneSelector(s string, id string) []byte {

	return nil
}
