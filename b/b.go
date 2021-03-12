package b

import (
	"fmt"
	"github.com/chmeliik/gomod-test/b/package-b"
)

func Main() {
	fmt.Println("Hello from \"b\" version 0.0.0")
	package_b.Foo()
}
