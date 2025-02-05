package playReady

import "encoding/binary"

// v2 = Struct(
//    "signature" / magic,
//    "version" / Int8ub,
//    "group_certificate_length" / Int32ub,
//    "group_certificate" / Bytes(this.group_certificate_length),
//    "encryption_key" / Bytes(96),
//    "signing_key" / Bytes(96),
// )
type play_ready struct {
   Signature [3]byte
   Version uint8
   GroupCertificateLength uint32
   GroupCertificate []byte
   EncryptionKey [96]byte
   SigningKey [96]byte
}

func (p *play_ready) encode() []byte {
   data := p.Signature[:]
   data = append(data, p.Version)
   data = binary.BigEndian.AppendUint32(data, p.GroupCertificateLength)
   data = append(data, p.GroupCertificate...)
   data = append(data, p.EncryptionKey[:]...)
   return append(data, p.SigningKey[:]...)
}

func (p *play_ready) decode(data []byte) {
   n := copy(p.Signature[:], data)
   data = data[n:]
   p.Version, data = data[0], data[1:]
   p.GroupCertificateLength, data = binary.BigEndian.Uint32(data), data[4:]
   p.GroupCertificate = data[:p.GroupCertificateLength]
   data = data[p.GroupCertificateLength:]
   n = copy(p.EncryptionKey[:], data)
   data = data[n:]
   copy(p.SigningKey[:], data)
}
