package vue

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/monopolly/file"
)

func New() (a *Vue) {
	return new(Vue)
}

type Vue struct {
	sync.Mutex
	dir    string
	files  []string
	js     []byte
	styles []byte
}

func Dir(dir, tojs, tocss string) (a *Vue) {
	a = new(Vue)
	a.Lock()
	a.dir = dir
	a.files = nil
	a.js = nil
	a.styles = nil
	var b bytes.Buffer
	var s bytes.Buffer
	file.Delete(tojs)
	file.Directory(dir, func(f os.FileInfo) {
		if f.IsDir() || filepath.Ext(f.Name()) != ".vue" {
			return
		}
		a.files = append(a.files, f.Name())
		c, err := a.compileFile(filepath.Join(dir, f.Name()))
		if err != nil {
			fmt.Println(err)
			return
		}
		b.Write(c.export())
		b.WriteString("\n")
		s.Write(c.styles)
	})
	a.js = b.Bytes()
	a.styles = s.Bytes()
	file.Save(tojs, a.js)
	file.Save(tocss, a.styles)
	a.Unlock()
	return
}

func (a *Vue) compileFile(file string) (c *component, err error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	return a.compile(b, file), nil
}

func (a *Vue) compile(b []byte, filename ...string) (c *component) {
	c = new(component)
	c.cssmap = make(map[string]string)
	c.body = b
	if len(filename) > 0 {
		c.filename = filename[0]
	}

	var doc bytes.Buffer
	doc.Write(b)
	c.doc, _ = goquery.NewDocumentFromReader(&doc)

	c.parse()
	return
}

func (a *Vue) CompileJS() []byte {
	return a.js
}

func (a *Vue) CompileCSS() []byte {
	return a.styles
}
