package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	slideAndRange()
	stringAndRange()
	mapAndRange()
	structAndPointer()
	slice()
	closure()
	interfaceTest()
	stringsTest()

}

func slideAndRange() {
	var arrays []int
	arrays = append(arrays, 1)
	arrays = append(arrays, 2)
	arrays = append(arrays, 3)
	for i, num := range arrays {
		fmt.Println(i, " ", num)
	}
	fmt.Println("========")
}

// stringAndRange 在 go 中，字符为rune，rune底层实际上为"整形"，如果想要输出字符串，需要使用 string(c)
func stringAndRange() {
	str := "abcdefg"
	for i, c := range str {
		fmt.Println(i, " ", string(c))
	}
	fmt.Println("========")
}

func mapAndRange() {
	hashmap := make(map[string]int, 1)
	hashmap["a"] = 1
	hashmap["b"] = 2
	fmt.Println(len(hashmap))
	for k, v := range hashmap {
		fmt.Println(k, " ", v)
	}
	fmt.Println("========")
}

func structAndPointer() {
	vertex := struct {
		X int
		Y int
	}{1, 2}
	fmt.Println(vertex, " ", vertex.X, " ", vertex.Y)
	p := &vertex
	fmt.Println(p, " ", p.X, " ", vertex.Y)
	fmt.Println("========")
}

func slice() {
	a := []byte{'r', 'o', 'a', 'd'}
	b := a[:2]
	b[0] = 'l'
	fmt.Println(string(a))

	c := []byte{'i', 'n', 'g'}
	a = append(a, c...)
	fmt.Println(string(a))
	fmt.Println("========")
}

func closure() {
	f := createClosure()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
	fmt.Println("\n========")
}

func createClosure() func() int {
	a := 1
	b := 0
	return func() int {
		tmp := a
		a = b
		b = tmp + a
		return a
	}
}

func interfaceTest() {
	var ab Abser

	var myfloat MyFloat = -1
	vertext := Vertex{3, 4}

	ab = myfloat
	fmt.Println(ab.Abs())

	ab = &vertext
	fmt.Println(ab.Abs())
	fmt.Println("========")
}

type Abser interface {
	Abs() float64
}

type MyFloat float64
type Vertex struct {
	X int
	Y int
}

func (myFloat MyFloat) Abs() float64 {
	if myFloat < 0 {
		return float64(-myFloat)
	} else {
		return float64(myFloat)
	}
}
func (vertex *Vertex) Abs() float64 {
	return math.Sqrt(math.Pow(float64(vertex.X), 2) +
		math.Pow(float64(vertex.Y), 2))
}

func stringsTest() {
	fields := strings.Fields("hello world")
	for i, f := range fields {
		fmt.Println(i, " ", f)
	}
	fmt.Println("========")
}
