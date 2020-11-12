package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	got := Add(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}

}

/**
启动本地文档，启动后在 Third part部分可以看到我们的文档
godoc -http=localhost:6060
using module mode; GOMOD=...\gobyexample\learn-go-by-projects\learn-go-with-tests\go.mod

*/
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// 必须要写下面的一句话，不然文档中没有例子
	// output:6
}
