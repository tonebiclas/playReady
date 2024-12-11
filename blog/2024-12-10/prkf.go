package main

import (
   "bytes"
   "fmt"
   "os"
)

func zero(r rune) bool {
   return r == 0
}

func main() {
   a, err := os.ReadFile("PlayReadykeybox.bin")
   if err != nil {
      panic(err)
   }
   b, err := os.ReadFile("../../secrets/g1")
   if err != nil {
      panic(err)
   }
   begin := bytes.Index(a, b[:9])
   end := bytes.Index(a, b[len(b)-9:])
   for _, field := range bytes.FieldsFunc(a[:begin], zero) {
      fmt.Printf("%v %q\n", len(field), field)
   }
   fmt.Printf("\n%q\n\n", a[begin:end+9])
   for _, field := range bytes.FieldsFunc(a[end+9:], zero) {
      fmt.Printf("%v %q\n", len(field), field)
   }
}
