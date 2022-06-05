package main

import (
    "encoding/json"
    "fmt"
    "tour-design-partten/creational/AbstractFactory"
    "tour-design-partten/creational/Singleton"
)

func main() {
    s1 := Singleton.GetInstance()
    s2 := Singleton.GetInstance()
    s1.AddOne()
    s2.AddOne()
    fmt.Printf("%v \n", s1)
    fmt.Printf("%v \n", s2)

    factory := &AbstractFactory.Factory{}
    shoe := factory.Create(AbstractFactory.BRAND_ADIDAS, "3 dong ke", 43, "ABC BAC")
    shoe.Show()
    shoe.Go("", 0)
    fmt.Println("--------------------------------------")
    e, _ := json.Marshal(shoe)
    fmt.Println(string(e))
    fmt.Println("--------------------------------------")
}
