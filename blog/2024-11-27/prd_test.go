package playReady

import (
   "bytes"
   "os"
   "testing"
)

func TestRoundTrip(t *testing.T) {
   in, err := os.ReadFile("pr2k.prd")
   if err != nil {
      t.Fatal(err)
   }
   var play play_ready
   play.decode(in)
   out := play.encode()
   if !bytes.Equal(out, in) {
      t.Fatal("!bytes.Equal")
   }
}
