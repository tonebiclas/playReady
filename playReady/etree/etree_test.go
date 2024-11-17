package etree

import (
   "os"
   "testing"
)

func TestElement(t *testing.T) {
   doc := NewDocument()
   doc.Element().Create("WRMHEADER", func(e *Element) {
      e.CreateAttr("xmlns", "http://schemas.microsoft.com/DRM/2007/03/PlayReadyHeader")
      e.CreateAttr("version", "4.0.0.0")
      e.Create("DATA", func(e *Element) {
         e.Create("PROTECTINFO", func(e *Element) {
            e.Create("KEYLEN", func(e *Element) {
               e.CreateText("16")
            })
            e.Create("ALGID", func(e *Element) {
               e.CreateText("AESCTR")
            })
         })
         e.Create("KID", func(e *Element) {
            e.CreateText("UZ4Ci2rVvUSRD9S1/ZD7og==")
         })
      })
   })
   doc.Indent(1)
   doc.WriteTo(os.Stdout)
}
