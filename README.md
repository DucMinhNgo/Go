# Run
go run hello.go

1. Định nghĩa về biến trong ngôn ngữ lập trình

2. Cú pháp biến
var i int
i = 42

Viết gọn
var i int = 10

Gọn hơn
i := 10

3. Global and block scope
Khai báo biến theo khối
var (
	n int = 10
	m int = 20
	h string = "abc"
)

4. Shadow

5. Declared and not used
Báo lỗi khi khai báo mà không sử dụng

6. Export global scope
Tên biết viết hoa thì bên ngoài có thể truy xuất được

7. Naming convention
Nguyên tắc: camel

8. convert type
var j float32 = float32(i) // chuyển số nguyên sang số thực
strconv.Itoa(2) // chuyến số sang chuỗi
