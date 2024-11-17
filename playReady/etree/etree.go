package etree

import "github.com/beevik/etree"

type Document struct {
   *etree.Document
}

func NewDocument() *Document {
   var d Document
   d.Document = etree.NewDocument()
   return &d
}

func (d Document) Element() *Element {
   return &Element{&d.Document.Element}
}

type Element struct {
   *etree.Element
}

func (e *Element) Create(tag string, f func(*Element)) {
   f(&Element{
      e.Element.CreateElement(tag),
   })
}
