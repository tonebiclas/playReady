package playReady

import (
   "encoding/binary"
   "errors"
   "log/slog"
)

func new_leaf_cert() {
   sign_payload = _BCertStructs.BCert.build(new_bcert_container)
   hash_obj = SHA256.new(sign_payload)
   signer = DSS.new(group_key.key, 'fips-186-3')
   signature = signer.sign(hash_obj)
   signature_info = Container(
      signature_type=1,
      signature_size=64,
      signature=signature,
      signature_key_size=512,  # bits
      signature_key=group_key.public_bytes()
   )
   signature_info_attribute = Container(
      flags=1,
      tag=8,
      length=len(_BCertStructs.DrmBCertSignatureInfo.build(signature_info)) + 8,
      attribute=signature_info
   )
   new_bcert_container = Container(
      signature=b"CERT",
      version=1,
      total_length=0,  # filled at a later time
      certificate_length=0,  # filled at a later time
      attributes=ListContainer([
          basic_info_attribute,
          device_info_attribute,
          feature_attribute,
          key_info_attribute,
          manufacturer_info,
      ])
   )
   payload = _BCertStructs.BCert.build(new_bcert_container)
   new_bcert_container.certificate_length = len(payload)
   new_bcert_container.total_length = len(payload) + 144  # signature length
   new_bcert_container.attributes.append(signature_info_attribute)
   return cls(new_bcert_container)
}

// DrmBCertBasicInfo = Struct(
//    "cert_id" / Bytes(16),
//    "security_level" / Int32ub,
//    "flags" / Int32ub,
//    "cert_type" / Int32ub,
//    "public_key_digest" / Bytes(32),
//    "expiration_date" / Int32ub,
//    "client_id" / Bytes(16)
// )
type basic_info struct {
   CertId [16]byte
   SecurityLevel uint32
   Flags uint32
   CertType uint32
   PublicKeyDigest [32]byte
   ExpirationDate uint32
   ClientId [16]byte
}

func (b *bcert) info() (*basic_info, error) {
   for _, attr := range b.Attributes {
      if attr.header.Tag == 1 {
         info := &basic_info{}
         _, err := binary.Decode(attr.Attribute, binary.BigEndian, info)
         if err != nil {
            return nil, err
         }
         return info, nil
      }
   }
   return nil, errors.New("bcert.info")
}

// Attribute = Struct(
//   "flags" / Int16ub,
//   "tag" / Int16ub,
//   "length" / Int32ub,
//   "attribute" / Switch(
//       lambda this_: this_.tag,
//       {
//           1: DrmBCertBasicInfo,
//           2: DrmBCertDomainInfo,
//           3: DrmBCertPCInfo,
//           4: DrmBCertDeviceInfo,
//           5: DrmBCertFeatureInfo,
//           6: DrmBCertKeyInfo,
//           7: DrmBCertManufacturerInfo,
//           8: DrmBCertSignatureInfo,
//           9: DrmBCertSilverlightInfo,
//           10: DrmBCertMeteringInfo,
//           11: DrmBCertExtDataSignKeyInfo,
//           12: BCertExtDataContainer,
//           13: DrmBCertExtDataSignature,
//           14: Bytes(this.length - 8),
//           15: DrmBCertServerInfo,
//           16: DrmBcertSecurityVersion,
//           17: DrmBcertSecurityVersion
//       },
//       default=Bytes(this.length - 8)
//   )
// )
type attribute struct {
   header struct {
      Flags uint16
      Tag uint16
      Length uint32
   }
   Attribute []byte
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
   header struct {
      Signature signature
      Version uint32
      TotalLength uint32
      Flags uint32
      CertificateCount uint32
   }
   Certificates []bcert
}

func (b *bcert_chain) read(buf []byte) error {
   n, err := binary.Decode(buf, binary.BigEndian, &b.header)
   if err != nil {
      return err
   }
   buf = buf[n:]
   slog.Debug("bcert_chain", "header", b.header)
   for range b.header.CertificateCount {
      var cert bcert
      n, err := cert.decode(buf)
      if err != nil {
         return err
      }
      buf = buf[n:]
      b.Certificates = append(b.Certificates, cert)
   }
   return nil
}

func (a *attribute) decode(buf []byte) (int, error) {
   n, err := binary.Decode(buf, binary.BigEndian, &a.header)
   if err != nil {
      return 0, err
   }
   slog.Debug("attribute", "header", a.header)
   a.Attribute = buf[n:a.header.Length]
   return n + len(a.Attribute), nil
}

func (b *bcert) decode(buf []byte) (int, error) {
   ns, err := binary.Decode(buf, binary.BigEndian, &b.header)
   if err != nil {
      return 0, err
   }
   slog.Debug("bcert", "header", b.header)
   for ns < int(b.header.TotalLength) {
      var attr attribute
      n, err := attr.decode(buf[ns:])
      if err != nil {
         return 0, err
      }
      ns += n
      b.Attributes = append(b.Attributes, attr)
   }
   return ns, nil
}

func (s signature) String() string {
   return string(s[:])
}

type signature [4]byte

// BCert = Struct(
//    "signature" / Const(b"CERT"),
//    "version" / Int32ub,
//    "total_length" / Int32ub,
//    "certificate_length" / Int32ub,
//    "attributes" / GreedyRange(Attribute)
// )
type bcert struct {
   header struct {
      Signature signature
      Version uint32
      TotalLength uint32
      CertificateLength uint32
   }
   Attributes []attribute
}
