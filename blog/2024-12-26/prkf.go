package main

import (
   "bytes"
   "fmt"
   "os"
)

func main() {
   data, err := os.ReadFile("PlayReadykeybox.bin")
   if err != nil {
      panic(err)
   }
}
