#1. Problem
 + Ensure that a class has just a single instance
    + Chắc chắn rằng mỗi class sẽ chỉ có 1 instance trong toàn bộ app
 + Provide a global access point to that instance
    + Cung cấp 1 global access point để sử dụng
 + VD: 
    1. VD: kiểu như DB, hay file 
    2. giống context trong JS hoặc store trong redux
#2.  Nhược điểm
     + Debug khó khăn vì chỗ nào cũng access vào point đó 
     + Trong môi trường đa luồng thì config khá mất công vì đa luồng sẽ tạo ra nhiều instance (go routine) 
     + Chỉ đúng khi chạy tuần tự thôi
#3. Các cách implement 
1. Eager initialization
```go
      var instance *singleton = &singleton{
         count: 1,
      }ync
      func GetInstance() *singleton {
    	return instance
      }
```
> Đây là cách dễ nhất nhưng nó có một nhược điểm là mặc dù instance đã được khởi tạo nhưng có thể sẽ không dùng tới. vì vậy chúng ta có cách thứ 2.

2. Lazy initialization
```go
    var instance *singleton
    
    func GetInstance() Singleton {
        if instance != nil {
            return instance
        }
        instance = &singleton{
            1,
        }
        return instance
    }
```
> 1. Cách này khắc phục đc nhược điểm của cách 1 
> 2. Only use with single thread
> 3. Nếu dùng với g routine thì việc parttern này cần custom thêm vì các go routine chạy đồng thời.
> Dẫn đến cùng 1 thời điểm có thể 2 go routine check đc cả instance đều null
>   3.1 Thêm time.Sleep(1s)  GetInstance() và dùng go func để gọi. Kết quả point của các instance sẽ khác nhau. => sai 

## Resolve singleton with go routine 
1. Muốn chạy với go routine bắt buộc chỉ có 1 hàm GetInstance đc chạy cùng 1 thời 
điểm
    > Đọc sync.Once.Do(func(){}) 
2. Hoặc khởi tạo trước đó rồi
```go
    func init() {
         instance = &singleton{
             1,
         }
    }

    func GetInstance() *singleton {
        return instance
    }
```
> Cách này sẽ khởi tạo ngay khi import (Thực ra là cách 1)
> Nhưng 1 instance đc khởi tạo ngay khi import 
