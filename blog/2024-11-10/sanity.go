package main

import (
   "crypto/ecdh"
   "encoding/hex"
   "fmt"
)

const private = "99cc43dbb3391b8ef295f05bd8dc3f315ccc6835f1ca68a3905c954e590922c0"

func main() {
   so_file, err := hex.DecodeString(private)
   if err != nil {
      panic(err)
   }
   private_key, err := ecdh.P256().NewPrivateKey(so_file)
   if err != nil {
      panic(err)
   }
   public := private_key.PublicKey().Bytes()
   fmt.Println("private", private)
   fmt.Println("public", len(public), hex.EncodeToString(public))
}
