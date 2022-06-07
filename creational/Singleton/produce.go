package Singleton

import "fmt"

func PatternProduce() {
    s1 := GetInstance()
    s2 := GetInstance()
    s1.AddOne()
    s2.AddOne()
    fmt.Printf("%v \n", s1)
    fmt.Printf("%v \n", s2)
}
