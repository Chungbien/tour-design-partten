package AbstractFactory

import "fmt"

type Nike struct {
    Logo        string
    Size        int
    Description string
}

func (n *Nike) Go(l string, s int) {
    fmt.Printf("Your size of %v is %v", l, s)
}
func (n *Nike) Show() {
    fmt.Printf("NIKE - Your Nike is so beautiful.")
}
