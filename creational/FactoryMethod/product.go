const (
    BRAND_NIKE   = "nike"
    BRAND_ADIDAS = "adidas"
)

type IProduct interface {
    Go(logo string, size int)
    Show()
}
