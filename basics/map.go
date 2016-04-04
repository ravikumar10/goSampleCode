package main

import "fmt"

func main() {
    languages:= make(map[string]string)
    languages["cs"] = "C #"
    languages["js"] = "JavaScript"
    languages["rb"] = "Ruby"
    languages["go"] = "Golang"

    fmt.Println(languages["go"])	
    if lang, ok := languages["go"]; ok {    
    fmt.Println(lang, ok)
   }
/*
 languages := map[string]string{
    "cs": "C #",
    "js": "JavaScript",
    "rb": "Ruby",
    "go": "Golang",
}   
*/

}