package main

import "fmt"

type Rectangle struct {
    length, width int
    name string
}

func main() {
    r1 := Rectangle{2, 1, "rect1"}
    fmt.Println("Rectangle r1 is: ", r1)

    r2 := Rectangle{width:3, name:"rect2", length:4}
    fmt.Println("Rectangle r2 is: ", r2)

    pr := new (Rectangle) //get pointer to an instance with new keyword
    (*pr).width = 6 //set value using . notation by dereferencing pointer.
    pr.length = 8 //set value using . notation - same as previous.  There is no -> operator like in c++. Go automatically converts
    pr.name = "ptr_to_rectangle"
    fmt.Println("Rectangle pr as address is: ", pr) //Go performs default printing of structs
    fmt.Println("Rectangle pr as value is: ", *pr) //address and value are differentiated with an & symbol
}
