package playReady

import (
   "crypto/ecdh"
   "encoding/binary"
   "errors"
   "log/slog"
)

// DrmBCertBasicInfo = Struct(
//    "cert_id" / Bytes(16),
//    "security_level" / Int32ub,
//    "flags" / Int32ub,
//    "cert_type" / Int32ub,
//    "public_key_digest" / Bytes(32),
//    "expiration_date" / Int32ub,
//    "client_id" / Bytes(16)
// )
type cert_basic_info struct {
   CertId [16]byte
   SecurityLevel uint32
   Flags uint32
   CertType uint32
   PublicKeyDigest [32]byte
   ExpirationDate uint32
   ClientId [16]byte
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

func (c *cert_basic_info) attribute() (*attribute, error) {
   var (
      attr attribute
      err error
   )
   attr.Attribute, err = binary.Append(nil, binary.BigEndian, c)
   if err != nil {
      return nil, err
   }
   attr.header.Flags = 1
   attr.header.Length = uint32(binary.Size(attr.header) + len(attr.Attribute))
   attr.header.Tag = 1
   return &attr, nil
}

// DrmBCertDeviceInfo = Struct(
//    "max_license" / Int32ub,
//    "max_header" / Int32ub,
//    "max_chain_depth" / Int32ub
// )
type cert_device_info struct {
   MaxLicense uint32
   MaxHeader uint32
   MaxChainDepth uint32
}

func NewLeafCert(
   cert_id, client_id [16]byte,
   security_level uint32,
   signing_key *ecdh.PrivateKey,
) error {
   basic_info := cert_basic_info{
      CertType: 2,
      CertId: cert_id,
      SecurityLevel: security_level,
      ExpirationDate: 0xFF_FF_FF_FF,
      ClientId: client_id,
      PublicKeyDigest: public_sha256_digest(signing_key),
   }
   basic_info_attribute, err := basic_info.attribute()
   if err != nil {
      return err
   }
   _ = basic_info_attribute
   device_info := cert_device_info{
      MaxLicense: 10240,
      MaxHeader: 15360,
      MaxChainDepth: 2,
   }
   _ = device_info
   return nil
   //device_info_attribute = Container(
   //   flags=1,
   //   tag=4,
   //   length=len(_BCertStructs.DrmBCertDeviceInfo.build(device_info)) + 8,
   //   attribute=device_info
   //)
   //feature = Container(
   //   feature_count=3,
   //   features=ListContainer([
   //       4,  # SECURE_CLOCK
   //       9,  # REVOCATION_LIST_FEATURE
   //       13  # SUPPORTS_PR3_FEATURES
   //   ])
   //)
   //feature_attribute = Container(
   //   flags=1,
   //   tag=5,
   //   length=len(_BCertStructs.DrmBCertFeatureInfo.build(feature)) + 8,
   //   attribute=feature
   //)
   //cert_key_sign = Container(
   //   type=1,
   //   length=512,  # bits
   //   flags=0,
   //   key=signing_key.public_bytes(),
   //   usages_count=1,
   //   usages=ListContainer([
   //       1
   //   ])
   //)
   //cert_key_encrypt = Container(
   //   type=1,
   //   length=512,  # bits
   //   flags=0,
   //   key=encryption_key.public_bytes(),
   //   usages_count=1,
   //   usages=ListContainer([
   //       2
   //   ])
   //)
   //key_info = Container(
   //   key_count=2,
   //   cert_keys=ListContainer([
   //       cert_key_sign,
   //       cert_key_encrypt
   //   ])
   //)
   //key_info_attribute = Container(
   //   flags=1,
   //   tag=6,
   //   length=len(_BCertStructs.DrmBCertKeyInfo.build(key_info)) + 8,
   //   attribute=key_info
   //)
   //manufacturer_info = parent.get_certificate(0).get_attribute(7)
   //new_bcert_container = Container(
   //   signature=b"CERT",
   //   version=1,
   //   total_length=0,  # filled at a later time
   //   certificate_length=0,  # filled at a later time
   //   attributes=ListContainer([
   //       basic_info_attribute,
   //       device_info_attribute,
   //       feature_attribute,
   //       key_info_attribute,
   //       manufacturer_info,
   //   ])
   //)
   //sign_payload = _BCertStructs.BCert.build(new_bcert_container)
   //hash_obj = SHA256.new(sign_payload)
   //signer = DSS.new(group_key.key, 'fips-186-3')
   //signature = signer.sign(hash_obj)
   //signature_info = Container(
   //   signature_type=1,
   //   signature_size=64,
   //   signature=signature,
   //   signature_key_size=512,  # bits
   //   signature_key=group_key.public_bytes()
   //)
   //signature_info_attribute = Container(
   //   flags=1,
   //   tag=8,
   //   length=len(_BCertStructs.DrmBCertSignatureInfo.build(signature_info)) + 8,
   //   attribute=signature_info
   //)
   //payload = _BCertStructs.BCert.build(new_bcert_container)
   //new_bcert_container.certificate_length = len(payload)
   //new_bcert_container.total_length = len(payload) + 144  # signature length
   //new_bcert_container.attributes.append(signature_info_attribute)
   //return cls(new_bcert_container)
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

func (b *bcert_chain) read(data []byte) error {
   n, err := binary.Decode(data, binary.BigEndian, &b.header)
   if err != nil {
      return err
   }
   data = data[n:]
   slog.Debug("bcert_chain", "header", b.header)
   for range b.header.CertificateCount {
      var cert bcert
      n, err := cert.decode(data)
      if err != nil {
         return err
      }
      data = data[n:]
      b.Certificates = append(b.Certificates, cert)
   }
   return nil
}

func (b *bcert) info() (*cert_basic_info, error) {
   for _, attr := range b.Attributes {
      if attr.header.Tag == 1 {
         info := &cert_basic_info{}
         _, err := binary.Decode(attr.Attribute, binary.BigEndian, info)
         if err != nil {
            return nil, err
         }
         return info, nil
      }
   }
   return nil, errors.New("bcert.info")
}

func (a *attribute) decode(data []byte) (int, error) {
   n, err := binary.Decode(data, binary.BigEndian, &a.header)
   if err != nil {
      return 0, err
   }
   slog.Debug("attribute", "header", a.header)
   a.Attribute = data[n:a.header.Length]
   return n + len(a.Attribute), nil
}

func (b *bcert) decode(data []byte) (int, error) {
   ns, err := binary.Decode(data, binary.BigEndian, &b.header)
   if err != nil {
      return 0, err
   }
   slog.Debug("bcert", "header", b.header)
   for ns < int(b.header.TotalLength) {
      var attr attribute
      n, err := attr.decode(data[ns:])
      if err != nil {
         return 0, err
      }
      ns += n
      b.Attributes = append(b.Attributes, attr)
   }
   return ns, nil
}

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

func (s signature) String() string {
   return string(s[:])
}

type signature [4]byte
