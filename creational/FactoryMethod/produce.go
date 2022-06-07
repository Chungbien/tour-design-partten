package FactoryMethod

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "time"
)

func getSeason(i int) string {
    switch i {
    case 0:
        return "spring"
    case 1:
        return "summer"
    case 2:
        return "autumn"
    default:
        return "winter"

    }
}

func getCountry(i int) string {
    if i%2 == 0 {
        return "china"
    }
    return "rest"
}

func ProduceClothes() {
    for i := 0; i < 10; i++ {

        rand.Seed(time.Now().UnixNano())
        num := rand.Intn(4)

        s := getSeason(num)
        factory := &EuroFactory{season: s}

        product := factory.Create()
        fmt.Println(product)
    }
    fmt.Printf("------------------------------------------------")
    for i := 0; i < 10; i++ {

        rand.Seed(time.Now().UnixNano())
        num := rand.Intn(2)

        c := getCountry(num)
        factory := &AsianFactory{country: c}

        product := factory.Create()
        fmt.Println(product)
    }
}

func ProductSingle() {
    factory := &Factory{}
    shoe := factory.Create(BRAND_ADIDAS, "3 dong ke", 43, "ABC BAC")
    shoe.Show()
    shoe.Go("", 0)
    fmt.Println("--------------------------------------")
    e, _ := json.Marshal(shoe)
    fmt.Println(string(e))
    fmt.Println("--------------------------------------")
}
