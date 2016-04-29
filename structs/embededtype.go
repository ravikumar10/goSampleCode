package main

import ("fmt")

type Person struct {
    Name string
}
func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}
type Phone struct {
    Person
    Model string
}
func main() {
  p := new(Phone)
  p.Name="Shiju"
  p.Talk() //a.Person.Talk() will aso work
}
