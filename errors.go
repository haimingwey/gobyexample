package main

import (
	"errors"
	"fmt"
)

type myError struct {
	code    int
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("code:%d, message:%s", e.code, e.message)
}

func f1(n int) (int, error) {
	if n < 0 {
		return 0, &myError{-1, "parameter cannot be negative"}
	}
	return n + 1, nil
}

func f2(n int) (int, error) {
	if n > 5 {
		return 0, errors.New("parameter cannot be bigger then 5")
	}
	return n + 1, nil
}

func main() {

	fmt.Println(f1(1))
	fmt.Println(f1(-1))

	fmt.Println(f2(4))
	fmt.Println(f2(10))

	for _, i := range []int{-1, 2} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	_, e := f1(-3)
	// TODO 类型断言??
	if ae, ok := e.(*myError); ok {
		fmt.Println(ae.code)
		fmt.Println(ae.message)
	}
}
