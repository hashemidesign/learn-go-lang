package main

import "fmt"

// interface type

type Animal interface {
	HowManyLegs() int
	Says() string
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func main() {
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	Riddle(&dog)

	var cat Cat
	cat.Name = "cat"
	cat.Sound = "meow"
	cat.NumberOfLegs = 4
	cat.HasTail = true

	Riddle(&cat)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf("This animal says %s and has %d legs. What animal is it?", a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
