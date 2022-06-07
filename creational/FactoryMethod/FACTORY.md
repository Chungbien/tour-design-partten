# Factory Method

> VD:
>
> - Trung tâm sản xuất quần áo hàng giả cần sản xuất lượng >quần áo giả của Nike và Adidas.
> - Nhưng phụ thuộc vào từng khu vực: Euro hoặc Asian mà >sản xuất lượng trang phục của mỗi hãng khác nhau sao cho >đỡ phí
> - Giả sử thì mỗi nơi cần sản xuất 10 bộ
> - Eu sản xuất 5 nike và 5 Adidas
> - Asian thì 2 nike và 8 adidas

## 1. What

```
Define an interface for creating an object, but let subclasses decide which class to instantiate. The Factory method lets a class defer instantiation it uses to subclasses.
```

## 2. Why

- Factory method là 1 pattern của Creational Pattern
- Che giấu việc khởi tạo instance
- Tăng khả năng tính mở rộng

## 3. How

> Nhìn vào code

### 3.1 Single Factory Method

```go
1. product.go
2. single-factory.go
```

Thực ra thì single ko đc tính là 1 pattern vì nó khá dễ. Đơn giản ta tạo ra 1 factory và trong **factory** đó có **method** có thể trả ra nhiều đầu ra nhưng tựu chung lại thì sẽ đều là 1 subclass kế thừa từ **IProduct**

```go
Create(brand string, l string, s int, o string) IProduct
```

### 3.1 Single Factory Method

```go
1. product.go
2. single-factory.go
```