package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"
)

func basic_values() {
	fmt.Println("Hello " + "World")

	fmt.Println("1+1 =", 1+1, "2+2", 2+2)
	fmt.Println("3.0/7.0=", 3.0/7.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func basic_variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

func basic_const(s string) {
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	//const l int = 10
	var l int = 10
	//l += 1
	fmt.Println(l)
}

func basic_for_loop() {
	i := 10
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 5; n < 12; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func basic_if_else() {
	if 2 > 1 {
		fmt.Println("2 > 1")
	}

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println("i is even")
		} else if i == 3 {
			fmt.Println("i is 3")
		} else {
			fmt.Println("i is odd")
		}
	}
}

func basic_switch() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(l interface{}) {
		switch t := l.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func basic_array() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// c := [3][3]int This wont work, cant dynamically set multi-dimensional array
	var c [3][3]int
	fmt.Println(c)
	c[0][1] = 10
	fmt.Println(c)
}

func basic_slices() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	g := []int{1}
	// g := []int{1, "a"} Cant have multiple types
	// g := []int Cant dynamically create a slice without a value
	fmt.Println(reflect.TypeOf(g))
	fmt.Println(reflect.TypeOf(twoD))
	fmt.Println(reflect.ValueOf(twoD).Kind())
	fmt.Println(reflect.ValueOf(g).Kind())
}

func basic_maps() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k1"]
	fmt.Println("prs:", prs)

	x, _ := m["k1"]
	fmt.Println("x", x)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	map_type := map[string]map[string]int{"a": {"b": 1}}
	fmt.Println("map_type:", map_type)
	fmt.Println("map_type['a']['b']:", map_type["a"]["b"])
	a, b := map_type["a"]
	fmt.Println(a)
	fmt.Println(b)

}

func basic_range() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func basic_function_add(a int, b int) int {
	return a + b
}

func basic_function_add_add(a, b, c int, d string) int {

	fmt.Println(d)
	return a + b + c
}

func basic_func_mult_return() (int, int) {
	return 1, 5
}

func basic_func_mult_return_type() (int, string) {
	return 1, "Hi"
}

func basic_variadic(nums ...int) int {
	sum := 0
	fmt.Print(nums, " ") // no new line with Print
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	return sum
}

func basic_variadic_type(a string, nums ...int) int {
	fmt.Println(a)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func basic_closure() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func basic_factorial(x int) int {
	fmt.Println("Start x: ", x)
	if x == 0 {
		fmt.Println("x == 0 ", x)
		return 1
	}
	y := x * basic_factorial(x-1)
	fmt.Println("y: ", y)
	return y
}

func basic_fib(past int, cur int, breaker int) int {
	for cur <= breaker {
		return basic_fib(cur, (past + cur), breaker)
	}

	return cur
}

func basic_next_fib(next int) int {
	some := basic_fib(0, 1, next)
	return some
}

func basic_pointer_zeroval(ival int) {
	fmt.Println("Ival Initial: ", ival)
	ival = 0
	fmt.Println("Ival Zero Set from Function: ", ival)
}

func basic_pointer_zeroptr(iptr *int) {
	fmt.Println("Ptr Initial Value: ", *iptr)
	fmt.Println("Ptr Address: ", iptr)
	*iptr = 0
	fmt.Println("Ptr Post Value: ", *iptr)
	fmt.Println("Ptr Address: ", &iptr)
}

type person struct {
	name string
	age  int
}

func new_person_ptr(person_name string) *person {

	p := person{name: person_name}
	p.age = 42
	return &p
}

func new_person(person_name string) person {

	p := person{name: person_name}
	p.age = 42
	return p
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return math.Pi * 2 * c.radius
}

func measure(g geometry) {
	fmt.Println("Shape: ", g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perim: ", g.perim())
}

func f_err_1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Cant work with 42.")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f_err_2(arg int) (int, error) {

	if arg == 42 {
		return -1, argError{arg, "cant work with 42"}
	}

	return arg + 3, nil
}

func main() {
	// basic_values()
	//basic_variables()
	//const s string = "constant"
	//basic_const(s)
	//basic_for_loop()
	//basic_if_else()
	//basic_switch()
	//basic_array()
	//basic_slices()
	//basic_maps()
	//basic_range()
	//fmt.Println(basic_function_add(1, 2))
	//basic_function_add_add(1, 2, 3, "hi")
	//a, b := basic_func_mult_return()
	//fmt.Println(a)
	//fmt.Println(b)

	//_, c := basic_func_mult_return()
	//fmt.Println(c)
	//d, e := basic_func_mult_return_type()
	//fmt.Println(d)
	//fmt.Println(e)

	//basic_variadic(1, 2, 3, 4)

	//c := make([]int, 0)
	//c = append(c, 1, 2, 3)
	//basic_variadic(c...)
	//basic_variadic_type("'String'", 3, 4, 5)

	////////////////////////////////
	// Closures
	////////////////////////////////
	//pos, neg := basic_closure(), basic_closure()

	//for i := 0; i < 10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(i*-1))
	//}

	//next_int := basic_closure()
	////The closure keeps track of the state for each initialized object
	//fmt.Println(next_int(1))
	//fmt.Println(next_int(2))
	//fmt.Println(next_int(3))

	//next_ints := basic_closure()
	//fmt.Println(next_ints(4))
	//fmt.Println(next_int(4))

	////////////////////////////////
	// Recursion
	////////////////////////////////
	//fmt.Println(basic_factorial(7))
	//fmt.Println(basic_fib(0, 1, 100))
	//fmt.Println(basic_next_fib(5))

	//var fib func(n int) int
	//fib = func(n int) int {
	//	if n < 2 {
	//		return n
	//	}
	//	return fib(n-1) + fib(n-2)

	//}
	//fmt.Println(fib(7))

	////////////////////////////////
	///// Pointers
	////////////////////////////////
	// ptr_val := 1
	// fmt.Println("Initial ptr_val: ", ptr_val)
	// basic_pointer_zeroval(ptr_val)
	// fmt.Println("Post zeroval ptr_val: ", ptr_val) // The function changes the value within its scope, not the variable within the arg

	// fmt.Println("Initial ptr_val: ", ptr_val)
	// fmt.Println("Initial Address: ", &ptr_val)
	// basic_pointer_zeroptr(&ptr_val)
	// fmt.Println("Post ptr_val: ", ptr_val)

	////////////////////////////////
	//// Struct
	////////////////////////////////

	// first_person := person{name: "Leslie Scott", age: 30}
	// fmt.Println("First Person: ", first_person)
	// fmt.Println("First Person name: ", first_person.name)

	// second_person := person{name: "Carla Benz"}
	// fmt.Println("First Person: ", second_person)
	// fmt.Println("First Person name: ", second_person.age)

	// third_person := person{"Abdul Rabbani", 23}
	// fmt.Println("Third Person: ", third_person)

	// fourth_person := &person{"Mark Reppert", 31}
	// fmt.Println("Fourth Person: ", fourth_person)
	// fmt.Println("Fourth Person's Age: ", fourth_person.age)

	// third_person.age = 42
	// fourth_person.age = 42
	// // Interesting how you can change the age through the pointer in the same way that you can through a direct reference...
	// fmt.Println("Change Third Person's age to 42 plz: ", third_person.age)
	// fmt.Println("Change Fourth Person's age to 42 plz: ", fourth_person.age)

	// temp_person := person{"Gerina Xhiherri", 23}
	// fmt.Println("New Person: ", new_person)

	// fifth_person := &temp_person
	// fmt.Println("Fifth Person: ", fifth_person)
	// fmt.Println("Fifth Person: ", fifth_person.age)
	// fifth_person.age = 42
	// fmt.Println("Change Fifth Person's age to 42 plz: ", fifth_person.age)

	// sixth_person := new_person("Tommy Parker")
	// fmt.Println("Sixth Person: ", sixth_person)

	// seventh_person := new_person_ptr("Trevor Brabrant")
	// fmt.Println("Seventh Person: ", seventh_person)

	////////////////////////////////
	////// methods
	////////////////////////////////
	/*
		Go automatically handles conversion between values and pointers for method calls.
		You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
	*/

	// shape_1 := rect{width: 10, height: 10}
	// fmt.Println("shape_1: ", shape_1)

	// fmt.Println("shape_1 area: ", shape_1.area())
	// fmt.Println("shape_1 perim: ", shape_1.perim())

	// shape_1_ptr := &shape_1
	// fmt.Println("shape_1_ptr: ", shape_1_ptr)
	// fmt.Println("shape_1_ptr area: ", shape_1_ptr.area())
	// fmt.Println("shape_1_ptr perim: ", shape_1_ptr.perim())

	////////////////////////////////
	///// Interface
	////////////////////////////////

	// circ := circle{radius: 5}
	// measure(circ)

	// rectangle := rect{width: 10, height: 10}
	// measure(&rectangle)
	// measure(rectangle) Wont work cause one of rectangle's methods requires a pointer input
	// You can pass a pointer when a function is expecting a value, but not a value when a func expects a pointer

	////////////////////////////////
	/// Errors
	////////////////////////////////

	for _, i := range []int{7, 42} {
		if r, e := f_err_1(i); e != nil {
			fmt.Println("f_err_1 failed: ", e)
		} else {
			fmt.Println("f_err_1 succeded: ", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f_err_2(i); e != nil {
			fmt.Println("f_err_2 failed: ", e)
		} else {
			fmt.Println("f_err_2 succeded: ", r)
		}
	}

	_, e := f_err_2(42)
	if ae, ok := e.(argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
