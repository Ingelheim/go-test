package template

import (
  "fmt"
)

type Template struct {
  Title string
  Body []byte
}

type Templateable interface {
	Render()
  SetTitle(title string)
}

type HomeTemplate struct {
  Template
}

func (t Template) Render() {
	fmt.Printf("RENDERING %s\n", t.Title)
}
