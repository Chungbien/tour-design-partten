package FactoryMethod

import "fmt"

type Adidas struct {
    Logo string
    Size int
    Note string
}

func (a Adidas) Go(l string, s int) {
    fmt.Printf("Your size of %v is %v \n", l, s)
    fmt.Printf("Your size of %v is %v \n", a.Logo, a.Size)
}
func (a *Adidas) Show() {
    fmt.Printf("ADIDAS - Amazing - good job | %v", a.Note)
}
