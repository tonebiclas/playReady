package playReady

import (
   "crypto/ecdh"
   "crypto/sha256"
)

func public_sha256_digest(signing_key *ecdh.PrivateKey) [32]byte {
   // M = 4 ‖ X ‖ Y
   // secg.org/sec1-v2.pdf
   data := signing_key.PublicKey().Bytes()[1:]
   return sha256.Sum256(data)
}
