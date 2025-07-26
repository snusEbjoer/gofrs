# gofrs

gofrs is iterator utility package

## Installation

```shell
go get github.com/snusEbjoer/gofrs
```

## Example

```go
func main() {
	type person struct {
		name string
		age  int
	}

	arr := []person{
		{name: "zazix", age: 22},
		{name: "mamix", age: 22},
		{name: "papix", age: 20},
		{name: "perdix", age: 21},
		{name: "corporate", age: 21},
		{name: "hypnosix", age: 21},
	}

	groups := seq.GroupBy(seq.FromSlice(arr), func(p person) int { return p.age })

	for _, persons := range groups {
		fmt.Printf("sended to brown country %v\n", persons)
	}
} // => sended to brown country [{zazix 22} {mamix 22}]
  // sended to brown country [{papix 20}]
  // sended to brown country [{perdix 21} {corporate 21} {hypnosix 21}]
``` 

## 0% test coverage because high TESTosteron level