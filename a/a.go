package a

import (
	"fmt"
	"github.com/chmeliik/gomod-test/a/v2/package-a"
)

func Main() {
	fmt.Println("Hello from \"a\" version 0.0.0")
	package_a.Foo()
}
