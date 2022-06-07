package FactoryMethod

// Creator
type ReligionFactory interface {
    Create(religion string, logo string, size int, o string) IProduct
}

// ConcreteCreator
type EuroFactory struct {
    season string
}

// ConcreteCreator
type AsianFactory struct {
    country string
}

func (f *EuroFactory) Create() IProduct {
    switch f.season {
    case "spring":
        return &Adidas{
            Logo: "Adidas",
            Size: 43,
            Note: "Spring",
        }
    case "summer":
        return &Adidas{
            Logo: "Adidas",
            Size: 44,
            Note: "Summer",
        }
    case "autumn":
        return &Adidas{
            Logo: "Adidas",
            Size: 45,
            Note: "Spring",
        }
    case "winter":
        return &Nike{
            Logo:        "Adidas",
            Size:        45,
            Description: "Winter solider",
        }

    }
    return &Adidas{
        Logo: "Adidas",
        Size: 45,
        Note: "Spring",
    }
}

func (f *AsianFactory) Create() IProduct {
    switch f.country {
    case "china":
        return &Nike{
            Logo:        "Niki",
            Size:        36,
            Description: "Chinnaaaaa",
        }
    default:
        return &Adidas{
            Logo: "ADIDAS",
            Size: 36,
            Note: "Rest of asian ",
        }
    }
}
