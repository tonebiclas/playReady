package playReady

def get_license_challenge(self, session_id: bytes, content_header: str) -> str:
  session = self.__sessions.get(session_id)
  if not session:
      raise InvalidSession(f"Session identifier {session_id!r} is invalid.")

  session.signing_key = self.signing_key
  session.encryption_key = self.encryption_key

  la_content = self._build_digest_content(
      content_header=content_header,
      nonce=base64.b64encode(get_random_bytes(16)).decode(),
      wmrm_cipher=base64.b64encode(self._get_key_data(session)).decode(),
      cert_cipher=base64.b64encode(self._get_cipher_data(session)).decode()
  )

  la_hash_obj = SHA256.new()
  la_hash_obj.update(la_content.encode())
  la_hash = la_hash_obj.digest()

  signed_info = self._build_signed_info(base64.b64encode(la_hash).decode())
  signed_info_digest = SHA256.new(signed_info.encode())

  signer = DSS.new(session.signing_key.key, 'fips-186-3')
  signature = signer.sign(signed_info_digest)

  # haven't found a better way to do this. xmltodict.unparse doesn't work
  main_body = (
      '<?xml version="1.0" encoding="utf-8"?>'
      '<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">'
          '<soap:Body>'
              '<AcquireLicense xmlns="http://schemas.microsoft.com/DRM/2007/03/protocols">'
                  '<challenge>'
                      '<Challenge xmlns="http://schemas.microsoft.com/DRM/2007/03/protocols/messages">'
                          + la_content +
                          '<Signature xmlns="http://www.w3.org/2000/09/xmldsig#">'
                              + signed_info +
                              f'<SignatureValue>{base64.b64encode(signature).decode()}</SignatureValue>'
                              '<KeyInfo xmlns="http://www.w3.org/2000/09/xmldsig#">'
                                  '<KeyValue>'
                                      '<ECCKeyValue>'
                                          f'<PublicKey>{base64.b64encode(session.signing_key.public_bytes()).decode()}</PublicKey>'
                                      '</ECCKeyValue>'
                                  '</KeyValue>'
                              '</KeyInfo>'
                          '</Signature>'
                      '</Challenge>'
                  '</challenge>'
              '</AcquireLicense>'
          '</soap:Body>'
      '</soap:Envelope>'
  )

  return main_body

def _decrypt_ecc256_key(self, session: Session, encrypted_key: bytes) -> bytes:
  point1 = Point(
      x=int.from_bytes(encrypted_key[:32], 'big'),
      y=int.from_bytes(encrypted_key[32:64], 'big'),
      curve=self.curve
  )
  point2 = Point(
      x=int.from_bytes(encrypted_key[64:96], 'big'),
      y=int.from_bytes(encrypted_key[96:128], 'big'),
      curve=self.curve
  )

  decrypted = self.elgamal.decrypt((point1, point2), int(session.encryption_key.key.d))
  return self.elgamal.to_bytes(decrypted.x)[16:32]

@staticmethod
def _verify_ecc_key(session: Session, licence: XMRLicense) -> bool:
  ecc_keys = list(licence.get_object(42))
  if not ecc_keys:
      raise InvalidLicense("No ECC public key in license")

  return ecc_keys[0].key == session.encryption_key.public_bytes()

def parse_license(self, session_id: bytes, licence: str) -> None:
  session = self.__sessions.get(session_id)
  if not session:
      raise InvalidSession(f"Session identifier {session_id!r} is invalid.")

  try:
      root = ET.fromstring(licence)
      license_elements = root.findall(".//{http://schemas.microsoft.com/DRM/2007/03/protocols}License")
      for license_element in license_elements:
          parsed_licence = XMRLicense.loads(license_element.text)

          if not self._verify_ecc_key(session, parsed_licence):
              raise InvalidLicense("Public encryption key does not match")

          for key in parsed_licence.get_content_keys():
              if Key.CipherType(key.cipher_type) == Key.CipherType.ECC256:
                  session.keys.append(Key(
                      key_id=UUID(bytes_le=key.key_id),
                      key_type=key.key_type,
                      cipher_type=key.cipher_type,
                      key_length=key.key_length,
                      key=self._decrypt_ecc256_key(session, key.encrypted_key)
                  ))
  except Exception as e:
      raise Exception(f"Unable to parse license, {e}")

def get_keys(self, session_id: bytes) -> List[Key]:
  session = self.__sessions.get(session_id)
  if not session:
      raise InvalidSession(f"Session identifier {session_id!r} is invalid.")

  return session.keys
