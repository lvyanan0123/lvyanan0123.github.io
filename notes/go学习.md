# Go

## 接口
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
```golang
pakage main

type Decode interface {
	Decode() error
}

type Encode interface {
	Encode() error
}

func DecodeJson(d Decode) {
	d.Decode()
}

func EncodeJson(d Encode) {
	d.Encode()
}

type ID struct {
}

func (d ID) Decode() error {
	fmt.Println("xxxxxxxxxxx")
	return nil
}

func (d ID) Encode() error {
	fmt.Println("yyyyyyyy")

	return nil
}

func main() {
	id := ID{}
	DecodeJson(id)
	EncodeJson(id)
}


```