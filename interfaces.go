package main

import (
	"fmt"
)

type car interface {
	brand() string
	wheels() string
	work() string
}

type taxi struct {
	name        string
	numOfWheels int
}

type truck struct {
	name        string
	numOfWheels int
}

func (t taxi) brand() string {
	return "taxi brand"
}

func (t taxi) work() string {
	return t.name + "is working"
}

func (t taxi) wheels() string {
	return fmt.Sprintf("%s has %d wheels", t.name, t.numOfWheels)
}

func (t truck) brand() string {
	return "truck brand"
}

func (t truck) work() string {
	return t.name + "is working"
}

func (t truck) wheels() string {
	return fmt.Sprintf("%s has %d wheels", t.name, t.numOfWheels)
}

func carInfo(c car) {
	fmt.Println(c.brand())
	fmt.Println(c.work())
	fmt.Println(c.wheels())
}

func main() {
	ta := taxi{name: "taxi", numOfWheels: 4}
	tr := truck{name: "truck", numOfWheels: 6}
	carInfo(ta)
	carInfo(tr)
}
