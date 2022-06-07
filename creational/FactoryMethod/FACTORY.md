# Factory Method

![markdown](https://refactoring.guru/images/patterns/diagrams/factory-method/structure.png)
**Thuật ngữ**

1. Creator: Interface Factory định nghĩa method
2. Concrete Creator: SubClass Factory triển khai các method dựa vào logic của riêng class đó
3. Concrete: instance đc tạo ra từ ConcreteCreator

> **_VD_**:
>
> - Trung tâm sản xuất quần áo hàng giả cần sản xuất lượng >quần áo giả của Nike và Adidas.
> - Nhưng phụ thuộc vào từng khu vực: Euro hoặc Asian mà >sản xuất lượng trang phục của mỗi hãng khác nhau sao cho >đỡ phí

## 1. What

```
Define an interface for creating an object, but let subclasses decide which class to instantiate. The Factory method lets a class defer instantiation it uses to subclasses.
```

## 2. Why

- Factory method là 1 pattern của Creational Pattern
- Che giấu việc khởi tạo instance
- Tăng khả năng tính mở rộng

- Tác dụng sẽ rất dễ nhận thấy trong code core, code common.
- Các instance đều có những method đó. Nhưng bạn ko muốn check cụ thể từng type hoặc có quá nhiều type dẫn đến code sẽ bị dài.
- Thay vào đó bạn tạo ra 1 cái **Factory** chuyên để tạo ra các instance.
- Vì các instance này có các method có input, output và chức năng tương đương nhau. (Cái khác nhau là khác ở logic extend bên trong)
- Giờ chỉ cần gọi là xong

> Xét VD:
> ở VN có hàng chục cái ngân hàng. Nhưng logic chuyển tiền của mỗi
> ngân hàng khác nhau. Nhưng input và output của việc chuyển tiền
> giống nhau.

```go

    bf := &BankFactory{}
    bank := bf.CreateBank() 
    if acc, exist := bank.GetAccout(id); exist {
        if acc.HasEnoughMoney() {
            acc.Transfer(amountMoney)
        }
    }
    return

```
Trong VD trên thì ta chẳng cần biết là bank gì. Logic check hạn mức hay tài khoản ra sao. Ta chỉ cần biết method của bank đó áp dụng đc là đc. Common cực kì nhanh. Extend việc thêm ngân hàng cũng cực kì dễ dàng. Không cần lặp code, ko sửa đổi gì code common.

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
2. factory-method.go
3. produce.go
```

Triển khai code trong VD ở đầu bài:

Euro sản xuất quần áo phụ thuộc vào mua => season

Asian người TQ thích NIke còn lại thích Adidas => country

Mỗi factory sẽ extend từ interface ReligonFactory và có các trả kết quả khác nhau. Không nhất thiết theo quy luật nào cả.
