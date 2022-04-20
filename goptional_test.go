package goptional_test

import (
	"github.com/harnyk/goptional"
)

func ExampleMaybe() {

	//Let's declare a function which returns a Maybe[int].
	myFn := func() goptional.Maybe[int] {
		return goptional.NewJust(1)
	}

	//We can call the function and may get either a Just or a Nothing.
	result := myFn()

	//We can check if the Maybe is a Just or a Nothing.
	switch result := result.(type) {
	case goptional.Just[int]:
		println(result.Unwrap())
	case goptional.Nothing[int]:
		println("Nothing")
	}
}
