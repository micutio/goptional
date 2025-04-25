package goptional_test

import (
	"testing"

	option "github.com/micutio/goptional"
)

func ExampleMaybe() {

	//Let's declare a function which returns a Maybe[int].
	myFn := func() option.Maybe[int] {
		return option.NewJust(1)
	}

	//We can call the function and may get either a Just or a Nothing.
	result := myFn()

	//We can check if the Maybe is a Just or a Nothing.
	switch result := result.(type) {
	case option.Just[int]:
		println(result.Get())
	case option.Nothing[int]:
		println("Nothing")
	}
}

func TestCreateEmptyOptional(t *testing.T) {
	emptyOpt := option.NewNothing[string]()

	if emptyOpt.IsPresent() {
		t.Errorf("optional %v should be empty!", emptyOpt)
	}
}

func TestCreatePresentOptional(t *testing.T) {
	presentOpt := option.NewJust("just an option")

	if !presentOpt.IsPresent() {
		t.Errorf("optional %v should be present!", presentOpt)
	}

	value := presentOpt.Get()
	if value != "just an option" {
		t.Errorf("optional %v should contain 'just an option', but contains '%s' instead", presentOpt, value)
	}
}
