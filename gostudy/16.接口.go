package main

import "fmt"

// Animal 定义一个animal的接口，它有唱，跳，rap的方法
type Animal interface {
	sing()
	jump()
	rap()
}

// Chicken 需要全部实现这些接口
type Chicken struct {
	Name string
}

func (c Chicken) sing() {
	fmt.Println("chicken 唱")
}

func (c Chicken) jump() {
	fmt.Println("chicken 跳")
}
func (c Chicken) rap() {
	fmt.Println("chicken rap")
}

func sing(obj Animal) {
	obj.sing()
}

func main() {
	var animal Animal
	animal = Chicken{"ik"}
	animal.sing()
	animal.jump()
	animal.rap()
	sing(animal)
}
