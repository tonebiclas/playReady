package etree

import (
   "os"
   "testing"
)

func TestElement(t *testing.T) {
   doc := NewDocument()
   doc.Element().Create("People", func(e *Element) {
      e.Create("Person", func(e *Element) {
         e.CreateAttr("name", "Jon")
      })
      e.Create("Person", func(e *Element) {
         e.CreateAttr("name", "Sally")
      })
   })
   doc.Indent(1)
   doc.WriteTo(os.Stdout)
}
