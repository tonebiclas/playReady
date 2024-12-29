package playReady

import (
   "encoding/binary"
   "strconv"
)

type prkb struct {
   qwFormatIdentifier [4]byte
   dwFormatVersion uint32
   eAlign uint32
}

func (p *prkb) decode(data []byte) {
   copy(p.qwFormatIdentifier[:], data)
   data = data[4:]
   p.dwFormatVersion = binary.BigEndian.Uint32(data)
   data = data[4:]
   p.eAlign = binary.BigEndian.Uint32(data)
}

func (p *prkb) String() string {
   b := []byte("qwFormatIdentifier = ")
   b = strconv.AppendQuote(b, string(p.qwFormatIdentifier[:]))
   b = append(b, "\ndwFormatVersion = "...)
   b = strconv.AppendUint(b, uint64(p.dwFormatVersion), 10)
   b = append(b, "\neAlign = "...)
   b = strconv.AppendUint(b, uint64(p.eAlign), 10)
   return string(b)
}
