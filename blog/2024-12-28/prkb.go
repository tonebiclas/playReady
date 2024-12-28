package main

import (
   "bytes"
   "encoding/binary"
   "fmt"
   "os"
)

type prkb struct {
   a [4]byte
   b uint32
   c []byte
}

func (p *prkb) decode(data []byte) {
   copy(p.a[:], data)
   data = data[4:]
   p.b = binary.BigEndian.Uint32(data)
   data = data[4:]
   p.c = data[:p.b]
}

func main() {
   data, err := os.ReadFile("ignore/playready-4.0.hds")
   if err != nil {
      panic(err)
   }
   _, data, _ = bytes.Cut(data, []byte("PRKB"))
   var play prkb
   play.decode(data)
   fmt.Printf("%+v\n", play)
}
