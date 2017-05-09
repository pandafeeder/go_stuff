//playground and memo for go
package main

import (
	"fmt"
)

//unused global variable will not cause compiler error
var unUsed = 100

func main() {
	play1()
	play2()
}

func play1() {
	//unused local const will not cause compiler error
	const (
		_ = iota
		one
		two
		three = "III"
		//once iota is interupted, need to recover it explicitly, previous enumeration still counts
		four = iota
		five
	)
	println("five is", five)

	s := "hello世界"
	//byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c,", s[i])
	}
	println()
	//rune
	for _, r := range s {
		fmt.Printf("%c,", r)
	}

	a := [3]int{1, 2, 3}
	//range will copy iterated object
	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a)
		}
		//prints out 0,1 | 1,2 | 2, 3
		println(i, v)
	}
	//slice is reference type, doesn't have above issue
	b := []int{1, 2, 3}
	for i, v := range b {
		if i == 0 {
			b[1], b[2] = 999, 999
			fmt.Println(b)
		}
		//prints out 0,1 | 1,999 | 2, 999
		println(i, v)
	}

L1:
	for x := 0; x < 3; x++ {
	L2:
		for y := 0; y < 5; y++ {
			if y > 2 {
				continue L2
			}
			if x > 1 {
				break L1
			}
			println(x, y)
		}
	}

	// return value of a map["key"] is a temporarily copied value
	// so, if the value is a sturct, there's no point modifying the returned struct
	myMap := map[int]struct {
		name string
		age  int
	}{
		1: {"Lisa", 28},
		2: {"Tom", 33},
	}
	//myMap[1].age = 29 error:cannot assign to struct field myMap[1].age in map
	//1: store back modifed value to map
	lisa := myMap[1]
	lisa.age = 29
	myMap[1] = lisa
	fmt.Println(myMap)
	//2: use pointer
	type user struct {
		name string
		age  int
	}
	myMap2 := map[int]*user{
		1: &user{"Lisa", 28},
		2: &user{"Tom", 33},
	}
	myMap2[1].age = 29
	fmt.Println(*myMap2[1])
}

//below are for play2
//think of struct like a object
type List struct {
	elements []int
}

//create new object using factory pattern
func newList() *List {
	return &List{elements: make([]int, 0, 10)}
}

//methods
func (l *List) length() int {
	return len(l.elements)
}
func (l *List) push(e int) {
	l.elements = append(l.elements, e)
}

func play2() {
	l := newList()
	println(l.length())
	i := 10
	l.push(i)
	l.push(i)
	l.push(i)
	fmt.Println(l.elements)

	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)
	//compiler will automatically convert between Type *Type for methods
	d.valueTest()
	d.pointerTest()

	//method value form, implicitly pass receiver
	mValue := d.pointerTest
	mValue()

	//method expression form, explicitly pass receiver
	mExpression := (*Data).pointerTest
	mExpression(&d)

	//Type cannot access *Type methods
	//mExpression2 := Data.pointerTest

	//while *Type can access Type methods
	mExpression3 := (*Data).valueTest
	mExpression3(&d)

	println()
	p.pointerTest()
	p.valueTest()
	fmt.Println(d)

	dog := Dog{}
	dog.makeSound()
	dog.sleep()
}

type Data struct {
	x int
}

func (d Data) valueTest() {
	//value type argument is copied, won't mutate original data
	fmt.Printf("Value: %p\n", &d)
	d.x = 100
}

func (d *Data) pointerTest() {
	fmt.Printf("Pointer: %p\n", d)
}

//mimic inheritance and method overrdie
type Animal struct {
}
type Dog struct {
	Animal
}

func (Animal) makeSound() {
	println("Animal making sound")
}
func (Animal) sleep() {
	println("I'm sleeping")
}
func (Dog) makeSound() {
	println("Wang Wang Wang")
}

//above are for play2
