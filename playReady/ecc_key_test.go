package playReady

import (
   "crypto/ecdh"
   "fmt"
   "testing"
)

func TestEcc(t *testing.T) {
   data := [32]byte{31: 1}
   key, err := ecdh.P256().NewPrivateKey(data[:])
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", public_sha256_digest(key))
}
