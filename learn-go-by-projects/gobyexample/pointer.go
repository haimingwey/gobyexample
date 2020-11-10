package main

func setI(i int) {
	i = 1
}

func setIByPointer(i *int) {
	*i = 5
}

func main() {

	i := 10
	println("i=", i)
	setI(i)
	println("after setI, i=", i)
	setIByPointer(&i)
	println("after setIByPointer, i=", i)

}
