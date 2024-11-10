package main

import (
   "crypto/ecdh"
   "encoding/hex"
   "fmt"
)

const public_key = "20d2f07cd7f74a60594d4a3da3da31266b6631fca3fce4b8125c9bd7036a19cfcf16ce392cc564cfc385da9c0920e7b752d51814e9239bd825c4b934121bf9b1"

func main() {
   so_file, err := hex.DecodeString("6eb9d6170372353bc402ace4980b3077bdb75ac273f0ef7dcea36f1af38b3693")
   if err != nil {
      panic(err)
   }
   private_key, err := ecdh.P256().NewPrivateKey(so_file)
   if err != nil {
      panic(err)
   }
   public := private_key.PublicKey().Bytes()[1:]
   fmt.Println(hex.EncodeToString(public) == public_key)
}
