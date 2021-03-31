package gsqli

import (
	"fmt"
	"testing"
)

func Test_SqlCheck(t *testing.T) {
	err := SQLInject("asdf asd ; -1' and 1=1 union/* foo */select load_file('/etc/passwd')--")
	if err == nil {
		t.Fail()
		return
	}
	fmt.Println(err.Error())
}

func Test_SqlTest(t *testing.T) {
	_, fp1 := SQLTest("dddddd asd ; -1' and 1=1 union/* foo */select ss('/estc/passwd')--33")
	fmt.Println(fp1)
	_, fp2 := SQLTest("ancd 222")
	fmt.Println(fp2)
}


func BenchmarkSQLInject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SQLInject("asdf asd ; -1' and 1=1 union/* foo */select load_file('/etc/passwd')--")
	}
}

func BenchmarkSQLInjectParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			SQLInject("asdf asd ; -1' and 1=1 union/* foo */select load_file('/etc/passwd')--")
		}
	})
}
