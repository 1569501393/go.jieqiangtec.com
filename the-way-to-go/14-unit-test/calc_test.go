package calc

import "testing"

/*Administrator@SKY-20191125ANA MINGW64 /d/www/qldev/go/go.jieqiangtec.com (test)
go test 时，要带上测试调用程序
$ go test -v the-way-to-go/14-unit-test/calc_test.go  the-way-to-go/14-unit-test/calc.go
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
PASS
ok      command-line-arguments  0.292s
*/

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Error("add(1,2) should be equal to 3")
	}
}

func TestAdd1(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: add test cases.
		{"a", args{0, 0}, 0},
		{"b", args{1, 1}, 2},
		// {"c", args{2, 2}, 0},
		{"c", args{2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
