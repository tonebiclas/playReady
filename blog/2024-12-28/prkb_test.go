package playReady

import (
   "bytes"
   "encoding/binary"
   "fmt"
   "os"
   "testing"
)

func TestPrkb(t *testing.T) {
   data, err := os.ReadFile("ignore/playready-4.0.hds")
   if err != nil {
      t.Fatal(err)
   }
   i := bytes.Index(data, []byte("CHAI"))
   data = data[i:]
   i = bytes.Index(data, bytes.Repeat([]byte{0}, 19))
   data = data[:i]
   fmt.Printf("%q\n", data)
   size := uint32(len(data))
   fmt.Println(size)
   fmt.Printf("%x\n", binary.BigEndian.AppendUint32(nil, size))
   fmt.Printf("%q\n", binary.BigEndian.AppendUint32(nil, size))
}

/*
from what I can tell the size of the g1 is 1856 - which would be
```
00 00 07 40
```
or
```
\x00\x00\a@
```
*/
