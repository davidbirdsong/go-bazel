package main

import (
	"fmt"

	"bazel-go/ar/bi/tr/ary/depth/subpackage"
	"bazel-go/types"
)

func main() {
	fmt.Printf("types import %v\n", types.NewFoo(0))
	subpackage.CallMe()
}
