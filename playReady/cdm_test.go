package playReady

import (
   "encoding/xml"
   "os"
   "testing"
)

func TestEnvelope(t *testing.T) {
   enc := xml.NewEncoder(os.Stdout)
   enc.Indent("", " ")
   err := enc.Encode(Envelope{})
   if err != nil {
      t.Fatal(err)
   }
}
