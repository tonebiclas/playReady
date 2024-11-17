package playReady

import (
   "fmt"
   "os"
   "testing"
)

func TestBcert(t *testing.T) {
   data, err := os.ReadFile("../secrets/g1")
   if err != nil {
      t.Fatal(err)
   }
   var bcert bcert_chain
   _, err = bcert.decode(data)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", bcert)
}
