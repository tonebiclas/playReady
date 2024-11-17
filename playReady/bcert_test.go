package playReady

import (
   "log/slog"
   "os"
   "testing"
)

func TestBcert(t *testing.T) {
   buf, err := os.ReadFile("../secrets/g1")
   if err != nil {
      t.Fatal(err)
   }
   var bcert bcert_chain
   slog.SetLogLoggerLevel(slog.LevelDebug)
   err = bcert.read(buf)
   if err != nil {
      t.Fatal(err)
   }
}
