package main
import (
	"fmt"
	"math/rand"
	"time"
)

type Gopher struct{
	value int
}
func (g *Gopher) UpdateValue(){
	rand.Seed(time.Now().Unix())
	rv:=rand.Intn(100)
	fmt.Println("Random value is",rv)
	g.value+=rv
}

func main() {
	g:=&Gopher { 100}
	g.UpdateValue()
	fmt.Println("Updated value is",g.value)
}
