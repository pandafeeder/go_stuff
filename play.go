//playground and memo for go
package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"sync"
)

//unused global variable will not cause compiler error
var unUsed = 100

func main() {
	runtime.GOMAXPROCS(1)
	play1()
	play2()
	play3()
	play4()
	play5()
}

func play1() {
	println("###### PLAY1 ######:")
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
	println("###### PLAY2 ######:")
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

//below are for play3
type Stringer interface {
	String() string
}
type Printer interface {
	Stringer
	Print()
}
type User struct {
	id   int
	name string
}

func (u *User) String() string {
	return fmt.Sprintf("user %d, %s", u.id, u.name)
}
func (u *User) Print() {
	fmt.Println(u.String())
}

//empty interfact{} is like root object in an oo language
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

type Shape interface {
	area() float64
}
type Rect struct {
	height, width float64
}
type Circle struct {
	r float64
}

func (r *Rect) area() float64 {
	return r.height * r.width
}
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func totalArea(shapes ...Shape) (area float64) {
	for _, shape := range shapes {
		area += shape.area()
	}
	return
}

//make type func() directly implement interface directly
type Tester interface {
	Do()
}
type funcDo func()

func (f funcDo) Do() {
	f()
}

func play3() {
	println("###### PLAY3 ######:")
	var t Printer = &User{1, "Tom"}
	t.Print()
	println(typeof(t))
	a := "hello"
	println(typeof(a))
	b := 32
	println(typeof(b))

	var o interface{} = &User{1, "Tom"}
	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println(i)
	}
	u1 := o.(*User)
	fmt.Println(u1)

	rect := Rect{10, 10}
	c := Circle{10}
	println(totalArea(&rect, &c))

	u := User{1, "Tom"}
	var i interface{} = u
	u.id = 2
	u.name = "Jack"
	fmt.Printf("%v\n", u)
	fmt.Printf("%v\n", i.(User))

	var vi, pi interface{} = u, &u
	//vi.(User).name = "Jack" //Error: cannot assiagn to vi.(User).name
	println()
	pi.(*User).name = "Jim"
	fmt.Printf("%v\n", vi.(User))
	fmt.Printf("%v\n", pi.(*User))
	fmt.Printf("%v\n", u)

	var t2 Tester = funcDo(func() { println("HELLO") })
	t2.Do()
}

//above are for play3

//below are for goroutine
func sum(id int) {
	var x int64
	for i := 0; i < math.MaxUint32; i++ {
		x += int64(i)
	}
	println(id, x)
}
func play4() {
	println("###### PLAY4 ######:")
	wg := new(sync.WaitGroup)
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer println("A.defer")

		func() {
			defer println("B.defer")
			//terminate execution with defer exected
			runtime.Goexit()
			println("B") //not executed
		}()

		println("A") //not executed
	}()
	wg.Wait()

	data := make(chan int)
	exit := make(chan bool)

	go func() {
		for i := range data {
			println(i)
		}
		println("receive over")
		exit <- true
	}()

	data <- 1
	data <- 2
	data <- 3
	close(data)
	println("send over")
	println(<-exit)

	data2 := make(chan int, 3)
	exit2 := make(chan bool)

	data2 <- 1
	data2 <- 2
	data2 <- 3
	go func() {
		for {
			if d, ok := <-data2; ok {
				println(d)
			} else {
				exit2 <- true
				break
			}
		}
	}()
	data2 <- 4
	data2 <- 5
	close(data2)
	println(<-exit2)

	c := make(chan int)
	var sendOnly chan<- int = c
	var recvOnly <-chan int = c

	println()
	go func() {
		for i := range recvOnly {
			println(i)
		}
	}()

	for i := 0; i < 3; i++ {
		sendOnly <- i
	}
	close(c)

	c1, c2 := make(chan int, 3), make(chan int)

	go func() {
		v, ok, s := 0, false, ""
		for {
			select {
			case v, ok = <-c1:
				s = "c1"
			case v, ok = <-c2:
				s = "c2"
			}

			if ok {
				fmt.Println(s, v)
			} else {
				break
			}
		}
	}()

	for i := 0; i < 5; i++ {
		select {
		case c1 <- i:
		case c2 <- i:
		}
	}

	close(c1)
	close(c2)

	wg2 := sync.WaitGroup{}
	wg2.Add(3)
	sem := make(chan int, 1)

	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg2.Done()
			sem <- 1
			for x := 0; x < 3; x++ {
				fmt.Println(id, x)
			}
			<-sem
		}(i)
	}

	wg2.Wait()

}

//above are for goroutine

//reflect

func play5() {
	laoli := Teacher{
		Person{"laoli", 33},
		"Teaching",
	}

	//type assertion: seems not parctical in statement code
	//more sutiable for function parameter type check
	var i interface{} = laoli
	s, ok := i.(Teacher)
	if ok {
		fmt.Println("Type assertion:")
		fmt.Println(s)
	}

	//reflect: more common in statement
	//1. concrete type
	t := reflect.TypeOf(laoli)
	if t.String() == "main.Teacher" {
		fmt.Println("Type checking by String value")
	}
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
	//2. pointer
	p := reflect.TypeOf(&laoli)
	if p.Kind() == reflect.Ptr {
		p = p.Elem()
	}
	for i, n := 0, p.NumField(); i < n; i++ {
		f := p.Field(i)
		fmt.Println(f.Name, f.Type)
	}

	//different accessbility to methods between Type and *Type
	methods := func(t reflect.Type) {
		for i, n := 0, t.NumMethod(); i < n; i++ {
			m := t.Method(i)
			fmt.Println(m.Name)
		}
	}
	var xiaohong Teacher
	fmt.Println("--- value interface ---")
	methods(reflect.TypeOf(xiaohong))
	fmt.Println("--- pointer interface ---")
	methods(reflect.TypeOf(&xiaohong))
}

type Person struct {
	Name string
	age  int
}
type Teacher struct {
	Person
	Vocation string
}

func (*Person) PersonMethod() {
	println("Person method")
}
func (Teacher) TeacherMethod() {
	println("Teacher method")
}
