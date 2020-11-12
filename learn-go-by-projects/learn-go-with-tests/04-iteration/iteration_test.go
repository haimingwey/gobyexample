package iteration

import "testing"

func TestRepeat(t *testing.T) {

	repeated := Repeat("a")
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

	repeated = RepeatWithStringsFunction("b")
	expected = "bbbbb"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

/**
run: ==> go test -bench=.
output:
goos: windows
goarch: amd64
pkg: github.com/haimingy/learn-go-with-tests/04-iteration
BenchmarkRepeat-4        7452990               182 ns/op
PASS
ok      github.com/haimingy/learn-go-with-tests/04-iteration    1.952s
*/
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func BenchmarkRepeatWithStringsFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWithStringsFunction("b")
	}
}

func ExampleRepeat() {
	Repeat("a")
	//output: "aaaaa"
}
