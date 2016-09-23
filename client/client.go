package main

import (
	"github.com/haibin/pkg"
)

func main() {
    DoSomethingVerbosely(&pkg.Foo{}, 3)
}

func DoSomethingVerbosely(foo *pkg.Foo, verbosity int) {
    // Could combine the next two lines,
    // with some loss of readability.
    prev := foo.Option(pkg.Verbosity(verbosity))
    defer foo.Option(prev)
    // ... do some stuff with foo under high verbosity.
}