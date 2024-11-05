package main

import (
   "flag"
   "fmt"
   "os"
   "os/exec"
)

func main() {
   err := os.Setenv("PATH", "C:/Program Files/Android/Android Studio/jbr/bin")
   compile := flag.Bool("c", false, "compile")
   flag.Parse()
   if *compile {
      cmd := exec.Command(
         "javac", "-d", ".", "*.java", "helpers/*.java", "mod/mspr/*.java",
      )
      cmd.Stderr = os.Stderr
      cmd.Stdout = os.Stdout
      fmt.Println(cmd.Args)
      err := cmd.Run()
      if err != nil {
         panic(err)
      }
   }
   os.Remove("secrets/genchain")
   if err != nil {
      panic(err)
   }
   data, err := exec.Command("java", "agsecres.tool.Hello").CombinedOutput()
   if err != nil {
      panic(err)
   }
   err = os.WriteFile("playReady.txt", data, os.ModePerm)
   if err != nil {
      panic(err)
   }
}
