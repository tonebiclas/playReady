package playReady

import (
   "os"
   "testing"
)

func TestLicense(t *testing.T) {
   resp, err := license()
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

func TestChallenge(t *testing.T) {
   doc := license_challenge()
   doc.Indent(1)
   doc.WriteTo(os.Stdout)
}
