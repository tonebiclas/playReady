package playReady

import (
   "os"
   "testing"
)

func TestCdm(t *testing.T) {
   doc := get_license_challenge()
   doc.Indent(1)
   doc.WriteTo(os.Stdout)
}
