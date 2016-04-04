package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
}
