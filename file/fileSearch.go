package main
import (
"fmt"
"time"
"os"
"path/filepath"
"strings"
)
func main() {
  fmt.Print("\nEnter the File Name :  ")
  var input string
  fmt.Scanln(&input)
  fmt.Println("\nHave Patience ! searching your file ......  \n")
  fmt.Println("\n====\tMatched File Found at below path's\t====  \n")
  start := time.Now()
  filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    if strings.Contains(path,input)==true{
        fmt.Println("\n"+path)
    }

    return nil
  })
  elapsed := time.Since(start)
  fmt.Println("\n==============================================================\n")
  fmt.Println("\nTotal Time Taken for the search in the Entire Directory Structure is : %s", elapsed)
  go forever()
     select {} // block forever
 }

 func forever() {
     for {
         //fmt.Printf("%v+\n", time.Now())
         time.Sleep(time.Second)
     }
 }
