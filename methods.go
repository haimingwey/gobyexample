package main

import "fmt"

type rectangle struct {
	height, width int
}

func area(rec rectangle) int {
	return rec.height * rec.width
}

func perimeter(rec *rectangle) int {
	return 2*rec.width + 2*rec.height
}

// 类型接收器方法
func (rec rectangle) area1() int {
	return rec.width * rec.height
}

func (rec *rectangle) perimeter1() int {
	return 2*rec.height + 2*rec.width
}

func main() {
	rec := rectangle{height: 2, width: 3}
	fmt.Println("area:", area(rec))
	fmt.Println("perimeter:", perimeter(&rec))
	fmt.Println("area1:", rec.area1())
	fmt.Println("perimeter1:", rec.perimeter1())
}
