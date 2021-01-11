package main

type x struct{}

func (x) String() string {
	return "Haaaaaaa"
}

type Int int

func (Int) String() string {
	return "integerrrrrrrr"
}

func (Int) toString() string {
	return "xxx"
}

func main() {
	PrintString(Int(0))
}

func PrintString(s Stringer) {
	println(s.String())
}

type Stringer interface {
	String() string
}
