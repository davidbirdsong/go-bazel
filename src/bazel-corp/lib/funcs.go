package subpackage

import (
	"bazel-go/types"
	"fmt"
)

func CallMe() {
	fmt.Println(types.NewFoo(100))
}
