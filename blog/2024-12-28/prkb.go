package playReady

import (
   "encoding/binary"
   "strconv"
)

func (p *prkb) String() string {
   b := []byte("qwFormatIdentifier = ")
   b = strconv.AppendQuote(b, string(p.qwFormatIdentifier[:]))
   b = append(b, "\ndwFormatVersion = "...)
   b = strconv.AppendUint(b, uint64(p.dwFormatVersion), 10)
   b = append(b, "\nc = "...)
   b = strconv.AppendQuote(b, string(p.c))
   return string(b)
}

type prkb struct {
   qwFormatIdentifier [4]byte
   dwFormatVersion uint32
   c []byte
}

func (p *prkb) decode(data []byte) {
   copy(p.qwFormatIdentifier[:], data)
   data = data[4:]
   p.dwFormatVersion = binary.BigEndian.Uint32(data)
   data = data[4:]
   p.c = data[:p.dwFormatVersion]
}
