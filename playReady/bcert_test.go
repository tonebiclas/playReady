package playReady

import (
   "fmt"
   "log/slog"
   "os"
   "testing"
)

func TestBcert(t *testing.T) {
   data, err := os.ReadFile("../secrets/g1")
   if err != nil {
      t.Fatal(err)
   }
   var bcert bcert_chain
   slog.SetLogLoggerLevel(slog.LevelDebug)
   err = bcert.read(data)
   if err != nil {
      t.Fatal(err)
   }
   for _, cert := range bcert.Certificates {
      info, err := cert.info()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", info)
   }
}
