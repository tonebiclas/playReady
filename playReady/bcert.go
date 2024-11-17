package playReady

import "encoding/binary"

// BCert = Struct(
//    "signature" / Const(b"CERT"),
//    "version" / Int32ub,
//    "total_length" / Int32ub,
//    "certificate_length" / Int32ub,
//    "attributes" / GreedyRange(Attribute)
// )
type Bcert struct {
   Signature signature
   Version uint32
   TotalLength uint32
   CertificateLength uint32
}

// BCertChain = Struct(
//    "signature" / Const(b"CHAI"),
//    "version" / Int32ub,
//    "total_length" / Int32ub,
//    "flags" / Int32ub,
//    "certificate_count" / Int32ub,
//    "certificates" / GreedyRange(BCert)
// )
type bcert_chain struct {
   Signature signature
   Version uint32
   TotalLength uint32
   Flags uint32
   CertificateCount uint32
}

func (s signature) String() string {
   return string(s[:])
}

type signature [4]byte

func (b *bcert_chain) decode(data []byte) (int, error) {
   return binary.Decode(data, binary.BigEndian, b)
}

/*
Attribute = Struct(
  "flags" / Int16ub,
  "tag" / Int16ub,
  "length" / Int32ub,
  "attribute" / Switch(
      lambda this_: this_.tag,
      {
          1: DrmBCertBasicInfo,
          2: DrmBCertDomainInfo,
          3: DrmBCertPCInfo,
          4: DrmBCertDeviceInfo,
          5: DrmBCertFeatureInfo,
          6: DrmBCertKeyInfo,
          7: DrmBCertManufacturerInfo,
          8: DrmBCertSignatureInfo,
          9: DrmBCertSilverlightInfo,
          10: DrmBCertMeteringInfo,
          11: DrmBCertExtDataSignKeyInfo,
          12: BCertExtDataContainer,
          13: DrmBCertExtDataSignature,
          14: Bytes(this.length - 8),
          15: DrmBCertServerInfo,
          16: DrmBcertSecurityVersion,
          17: DrmBcertSecurityVersion
      },
      default=Bytes(this.length - 8)
  )
)
*/
