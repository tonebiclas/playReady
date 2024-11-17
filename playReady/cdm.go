package playReady

import "playReady/etree"

func get_license_challenge() *etree.Document {
   doc := etree.NewDocument()
   doc.Element().Create("soap:Envelope", func(e *etree.Element) {
      e.CreateAttr("xmlns:soap", "http://schemas.xmlsoap.org/soap/envelope/")
      e.Create("soap:Body", func(e *etree.Element) {
         e.Create("AcquireLicense", func(e *etree.Element) {
            e.CreateAttr("xmlns", "http://schemas.microsoft.com/DRM/2007/03/protocols")
            e.Create("challenge", func(e *etree.Element) {
               e.Create("Challenge", func(e *etree.Element) {
                  e.CreateAttr("xmlns", "http://schemas.microsoft.com/DRM/2007/03/protocols/messages")
                  e.Create("LA", func(e *etree.Element) {
                     e.CreateAttr("Id", "SignedData")
                     e.CreateAttr("xml:space", "preserve")
                     e.CreateAttr("xmlns", "http://schemas.microsoft.com/DRM/2007/03/protocols")
                     e.Create("Version", func(e *etree.Element) {
                        e.SetText("1")
                     })
                     e.Create("ContentHeader", func(e *etree.Element) {
                        e.Create("WRMHEADER", func(e *etree.Element) {
                           e.CreateAttr("xmlns", "http://schemas.microsoft.com/DRM/2007/03/PlayReadyHeader")
                        })
                     })
                  })
               })
            })
         })
      })
   })
   return doc
}
