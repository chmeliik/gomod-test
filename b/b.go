package c

import (
	"fmt"
	"github.com/chmeliik/gomod-test/c/package-b"
)

func Main() {
	fmt.Println("Hello from \"c\" version 0.0.0")
	package_b.Foo()
}
