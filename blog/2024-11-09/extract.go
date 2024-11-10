package main

import (
   "crypto/ecdh"
   "encoding/hex"
   "fmt"
   "os"
)

//fail
//const public_key = "142ab9b1ef6dc50137a8b0c18f9d4d1dacc40db4551296201ebeae3fee825e50a6eaab6b08a3f44e1637e4b1a3fe6bcbb1d42ffdcdf4956cca5247eb7bf4bf53"
//const public_key = "d893cbc24c6823e8a243cfcdaac20c0018e4227547786ab0fd6a6890e30038d3567d7e0d933d85ba9c2ffffdfcd78b64f28d0fc9eea5b3c6aea713346edaebcf"
const public_key = "b85186ca93a95ce987647b8ab76be7e9d00ae261036088bbbd5f7403f059cf52759b8207a6033404994671f5137dd0fafc9eceb15402af64293c93e17f1fa645"

func main() {
   so_file, err := os.ReadFile("libwkf_support.so")
   if err != nil {
      panic(err)
   }
   var count int
   for i := range so_file {
      count++
      if count == 999 {
         fmt.Println(len(so_file)-i)
         count = 0
      }
      key := so_file[i:i+32]
      private_key, err := ecdh.P256().NewPrivateKey(key)
      if err != nil {
         fmt.Println(err)
         continue
      }
      public := private_key.PublicKey().Bytes()[1:]
      if hex.EncodeToString(public) == public_key {
         fmt.Printf("%x\n", key)
         break
      }
   }
}
