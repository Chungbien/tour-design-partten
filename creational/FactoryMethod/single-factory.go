package FactoryMethod

type IFactory interface {
    Create(brand string, l string, s int, o string) IProduct
}

type Factory struct{}

func (f *Factory) Create(brand string, l string, s int, o string) IProduct {
    switch brand {
    case BRAND_NIKE:
        return &Nike{
            Logo:        l,
            Size:        s,
            Description: o,
        }
    case BRAND_ADIDAS:
        return &Adidas{
            Logo: l,
            Size: s,
            Note: o,
        }
    default:
        return nil
    }
}

/**
  !!! NOTE
  + Tại sao Kiểu trả ra là IProduct (value) mà return lại là 1 con trỏ
  =>
      Vì func Create là method của 1 con trỏ *Factory chứ ko Factory
      Nike và Adidas kế thừa Go and Show cũng là con trỏ chứ ko phải là value
      => IFactory ~ 1 point of Factory or Nike or Adidas
  ==> Nếu muốn return ra value
      1. *Factory -> Factory ở line 19
      2. Bỏ hết * ở Go và Show ở Nike và Adidas l
*/
