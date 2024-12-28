package playReady

import (
   "bytes"
   "fmt"
   "os"
   "testing"
)

func TestPrkb(t *testing.T) {
   data, err := os.ReadFile("ignore/playready-4.0.hds")
   if err != nil {
      t.Fatal(err)
   }
   i := bytes.Index(data, []byte("PRKB"))
   var play prkb
   play.decode(data[i:])
   fmt.Println(&play)
}
